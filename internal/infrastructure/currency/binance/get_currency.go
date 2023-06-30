package binance

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"exchange/internal/domain/rate"
)

const symbol = "symbol"

type getPriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (api *BinanceAPI) GetCurrency(ctx context.Context, cur *rate.Rate) (float64, error) {
	byteResp, err := api.tickerPrice(ctx, cur.BaseCurrency, cur.QuoteCurrency)
	if err != nil {
		return 0, err
	}

	resp := new(getPriceResponse)

	err = json.Unmarshal(byteResp, resp)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(resp.Price, 64)
}

func (api *BinanceAPI) tickerPrice(
	ctx context.Context,
	base, quote string,
) ([]byte, error) {
	const url = "ticker/price"

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/%s", api.cfg.baseURL, url),
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add(symbol, fmt.Sprintf("%s%s", strings.ToUpper(base), strings.ToUpper(quote)))
	req.URL.RawQuery = q.Encode()

	resp, err := api.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Join(
			fmt.Errorf(
				"cant reach binance. code:%v",
				resp.StatusCode,
			),
			rate.ErrThirdPartyRequest,
		)
	}

	return io.ReadAll(resp.Body)
}
