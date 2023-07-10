package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Proxy struct {
	currency *httputil.ReverseProxy
	notifier *httputil.ReverseProxy
}

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}

	proxy := loadProxy()

	registerCurrencyHandlers(proxy)
	registerNotifierHandlers(proxy)

	logrus.Fatal(http.ListenAndServe(":"+os.Getenv("HTTP_SERVER_ADDR"), nil))
}

func loadProxy() *Proxy {
	currencyURL, _ := url.Parse(
		fmt.Sprintf("http://%s:%s", os.Getenv("CURRENCY_ADDRESS"), os.Getenv("CURRENCY_PORT")),
	)
	notifierURL, _ := url.Parse(
		fmt.Sprintf("http://%s:%s", os.Getenv("NOTIFIER_ADDRESS"), os.Getenv("NOTIFIER_PORT")),
	)

	return &Proxy{
		currency: httputil.NewSingleHostReverseProxy(currencyURL),
		notifier: httputil.NewSingleHostReverseProxy(notifierURL),
	}
}

func registerNotifierHandlers(p *Proxy) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.notifier.ServeHTTP(w, r)
	})

	http.Handle("/api/subscribe", exampleMiddleware(handler))
	http.Handle("/api/sendEmails", exampleMiddleware(handler))
}

func registerCurrencyHandlers(p *Proxy) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.currency.ServeHTTP(w, r)
	})
	http.Handle("/api/rate", exampleMiddleware(handler))
}

func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := logrus.WithFields(logrus.Fields{
			"Method": r.Method,
			"URI":    r.RequestURI,
		})

		l.Info()

		next.ServeHTTP(w, r)
	})
}
