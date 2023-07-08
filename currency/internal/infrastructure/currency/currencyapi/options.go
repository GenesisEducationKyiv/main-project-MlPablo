package currencyapi

import (
	"net/http"
	"net/http/httputil"
)

type (
	Logger interface {
		Infof(format string, args ...any)
	}

	Option func(*CurrencyAPI)

	loggingRoundTripper struct {
		proxied http.RoundTripper
		log     Logger
	}
)

func WithLogger(logger Logger) Option {
	return func(b *CurrencyAPI) {
		b.cli.Transport = loggingRoundTripper{http.DefaultTransport, logger}
	}
}

func (lrt loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := lrt.proxied.RoundTrip(r)

	respBytes, _ := httputil.DumpResponse(resp, true)

	lrt.log.Infof("CurrencyAPI-Response:\n%s\n", respBytes)

	return resp, err
}
