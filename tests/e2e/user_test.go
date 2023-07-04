package e2e

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/go-faker/faker/v4"

	"exchange/internal/domain/user"
)

func (suite *Suite) TestValidCreateUser() {
	u := user.NewUser(faker.Email())

	rec := httptest.NewRecorder()

	suite.createUser(u.Email, rec)

	suite.Equal(http.StatusOK, rec.Code)

	ok, err := suite.checkRowExistInFileDB(u.Email)
	suite.Require().NoError(err)
	suite.True(ok)
}

func (suite *Suite) TestInvalidEmail() {
	u := user.NewUser(faker.Word())

	rec := httptest.NewRecorder()

	suite.createUser(u.Email, rec)

	suite.Equal(http.StatusBadRequest, rec.Code)
}

func (suite *Suite) TestSameEmail() {
	u := user.NewUser(faker.Email())

	rec := httptest.NewRecorder()

	suite.createUser(u.Email, rec)
	suite.Equal(http.StatusOK, rec.Code)

	rec = httptest.NewRecorder()

	suite.createUser(u.Email, rec)
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
