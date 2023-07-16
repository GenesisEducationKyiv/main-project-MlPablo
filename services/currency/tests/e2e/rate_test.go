package e2e

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"currency/internal/domain/rate"
)

func (suite *Suite) TestGetRate() {
	rec := httptest.NewRecorder()

	suite.getRate(rec)
	suite.Require().Equal(http.StatusOK, rec.Code)

	var respRate rate.Currency

	err := json.Unmarshal(rec.Body.Bytes(), &respRate)
	suite.NoError(err)
	suite.Require().Equal(http.StatusOK, rec.Code)
	suite.Require().NotZero(respRate.Value)
}

func (suite *Suite) getRate(w http.ResponseWriter) {
	req := httptest.NewRequest(http.MethodGet, "/api/rate", nil)
	suite.e.ServeHTTP(w, req)
}
