package client

import (
	"net/http"
	"net/http/httputil"
)

type (
	Logger interface {
		Infof(format string, args ...any)
	}

	Option func(*http.Client)

	loggingRoundTripper struct {
		proxied http.RoundTripper
		log     Logger
	}
)

func New(opts ...Option) *http.Client {
	cli := http.DefaultClient

	for _, opt := range opts {
		opt(cli)
	}

	return cli
}

func WithLogger(logger Logger) Option {
	return func(cli *http.Client) {
		cli.Transport = loggingRoundTripper{http.DefaultTransport, logger}
	}
}

func (lrt loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	lrt.log.Infof("Request to:\n%s\n", r.URL.String())

	resp, err := lrt.proxied.RoundTrip(r)

	respBytes, _ := httputil.DumpResponse(resp, true)

	lrt.log.Infof("Response:\n%s\n", respBytes)

	return resp, err
}
