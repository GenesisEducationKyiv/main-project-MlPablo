package coingecko

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"currency/internal/domain/rate"
)

const (
	ids           = "ids"
	quoteCurrency = "vs_currencies"
)

func (api *CoingeckoAPI) GetCurrency(ctx context.Context, cur *rate.Rate) (float64, error) {
	resp, err := api.simplePrice(ctx, cur.BaseCurrency, cur.QuoteCurrency)
	if err != nil {
		if api.next != nil {
			return api.next.GetCurrency(ctx, cur)
		}

		return 0, err
	}

	return getValueFromResponse(resp)
}

func (api *CoingeckoAPI) simplePrice(
	ctx context.Context,
	base, quote string,
) ([]byte, error) {
	const latest = "simple/price"

	baseGeckoID, ok := api.mapper[strings.ToUpper(base)]
	if !ok {
		return nil, fmt.Errorf("coingecko api: no such coin: %w", rate.ErrBaseNotSupported)
	}

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
	q.Add(ids, baseGeckoID)
	q.Add(quoteCurrency, strings.ToUpper(quote))
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Join(
			fmt.Errorf(
				"cant reach coingecko. code:%v",
				resp.StatusCode,
			),
			rate.ErrThirdPartyRequest,
		)
	}

	return io.ReadAll(resp.Body)
}

// Rsponse format due to coingecko documentation
// So we need get generic currency without concrate type implementation
//
//	{
//		"bitcoin": {
//			"uah": 1111111,
//		},
//	}
func getValueFromResponse(m []byte) (float64, error) {
	resp := make(map[string]map[string]any)

	if err := json.Unmarshal(m, &resp); err != nil {
		return 0, err
	}

	for _, quote := range resp {
		for _, v := range quote {
			if value, ok := v.(float64); ok {
				return value, nil
			}
		}
	}

	return 0, errors.Join(
		fmt.Errorf("unable to get currency. resp: %s", m),
		rate.ErrThirdPartyRequest,
	)
}
