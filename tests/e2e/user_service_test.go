package e2e

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/bxcodec/faker/v3"

	"exchange/internal/domain"
)

func (suite *Suite) TestValidCreateUser() {
	user := domain.NewUser(faker.Email())

	rec := httptest.NewRecorder()

	suite.createUser(user.Email, rec)

	suite.Equal(http.StatusOK, rec.Code)

	ok, err := suite.checkRowExistInFileDB(user.Email)
	suite.Require().NoError(err)
	suite.True(ok)
}

func (suite *Suite) TestInvalidEmail() {
	user := domain.NewUser(faker.Word())

	rec := httptest.NewRecorder()

	suite.createUser(user.Email, rec)

	suite.Equal(http.StatusBadRequest, rec.Code)
}

func (suite *Suite) TestSameEmail() {
	user := domain.NewUser(faker.Email())

	rec := httptest.NewRecorder()

	suite.createUser(user.Email, rec)
	suite.Equal(http.StatusOK, rec.Code)

	rec = httptest.NewRecorder()

	suite.createUser(user.Email, rec)
	suite.Equal(http.StatusConflict, rec.Code)
}

func (suite *Suite) createUser(email string, w http.ResponseWriter) {
	req := httptest.NewRequest(http.MethodPost, "/api/subscribe", nil)

	form, err := url.ParseQuery(req.URL.RawQuery)
	suite.Require().NoError(err)

	form.Add("email", email)
	req.URL.RawQuery = form.Encode()

	suite.e.ServeHTTP(w, req)
}
