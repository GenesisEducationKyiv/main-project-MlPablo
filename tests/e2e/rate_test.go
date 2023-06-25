package e2e

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func (suite *Suite) TestGetRate() {
	rec := httptest.NewRecorder()

	suite.getRate(rec)
	suite.Require().Equal(http.StatusOK, rec.Code)

	var respRate float64

	err := json.Unmarshal(rec.Body.Bytes(), &respRate)
	suite.NoError(err)
	suite.Require().Equal(http.StatusOK, rec.Code)
	suite.Require().NotZero(respRate)
}

func (suite *Suite) getRate(w http.ResponseWriter) {
	req := httptest.NewRequest(http.MethodGet, "/api/rate", nil)
	suite.e.ServeHTTP(w, req)
}
