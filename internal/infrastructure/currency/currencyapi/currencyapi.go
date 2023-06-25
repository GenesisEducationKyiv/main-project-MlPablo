package currencyapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"exchange/internal/domain/rate_domain"
)

const (
	apiKey       = "apikey"
	baseCurrency = "base_currency"
	currencies   = "currencies"
)

type CurrencyAPI struct {
	cfg *Config
}

// This is the implementation of logic that can get currency.
// So service doesn't need to know about how we do this, and we can implement any currency api and interfaces we want
// I'm not sure about putting this into infrastructure folder.
func NewCurrencyAPI(cfg *Config) *CurrencyAPI {
	return &CurrencyAPI{
		cfg: cfg,
	}
}

func (api *CurrencyAPI) GetCurrency(ctx context.Context, cur *rate_domain.Rate) (float64, error) {
	resp, err := api.makeLatestCurrencyRequest(ctx, cur.BaseCurrency, cur.QuoteCurrency)
	if err != nil {
		return 0, err
	}

	return getValueFromResponse(resp, cur.QuoteCurrency)
}

func (api *CurrencyAPI) makeLatestCurrencyRequest(
	ctx context.Context,
	base, quote string,
) ([]byte, error) {
	const latest = "latest"

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/%s", api.cfg.baseURL, latest),
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add(apiKey, api.cfg.apiKey)
	q.Add(baseCurrency, strings.ToUpper(base))
	q.Add(currencies, strings.ToUpper(quote))
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Join(
			fmt.Errorf(
				"cant reach currencyapi. code:%v",
				resp.StatusCode,
			),
			rate_domain.ErrThirdPartyRequest,
		)
	}

	return io.ReadAll(resp.Body)
}

// Rsponse format due to currencyapi documentation
// So we need get generic currency without concrate type implementation
//
//	{
//	    "meta": {
//	        "last_updated_at": "2022-01-01T23:59:59Z"
//	    },
//	    "data": {
//	        "AED": {
//	            "code": "AED",
//	            "value": 3.67306
//	        },
//	        "AFN": {
//	            "code": "AFN",
//	            "value": 91.80254
//	        },
//	        "...": "150+ more currencies"
//	    }
//	}
func getValueFromResponse(m []byte, curr string) (float64, error) {
	const (
		dataField  = "data"
		valueField = "value"
	)

	parserErr := errors.Join(
		fmt.Errorf("unable to get field: %s, from: %s", curr, m),
		rate_domain.ErrThirdPartyRequest,
	)

	resp := make(map[string]interface{})

	if err := json.Unmarshal(m, &resp); err != nil {
		return 0, err
	}

	data, ok := resp[dataField].(map[string]interface{})
	if !ok {
		return 0, parserErr
	}

	info, ok := data[curr].(map[string]interface{})
	if !ok {
		return 0, parserErr
	}

	val, ok := info[valueField].(float64)
	if !ok {
		return 0, parserErr
	}

	return val, nil
}
