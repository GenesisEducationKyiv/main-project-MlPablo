package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const HeaderTimeout = 3

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

	server := &http.Server{
		Addr:              ":" + os.Getenv("HTTP_SERVER_ADDR"),
		ReadHeaderTimeout: HeaderTimeout * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		logrus.Fatal(err)
	}
}

func loadProxy() *Proxy {
	currencyAddress := net.JoinHostPort(os.Getenv("CURRENCY_ADDRESS"), os.Getenv("CURRENCY_PORT"))
	notiferAddress := net.JoinHostPort(os.Getenv("NOTIFIER_ADDRESS"), os.Getenv("NOTIFIER_PORT"))
	currencyURL, _ := url.Parse(fmt.Sprintf("http://%s", currencyAddress))
	notifierURL, _ := url.Parse(fmt.Sprintf("http://%s", notiferAddress))

	return &Proxy{
		currency: httputil.NewSingleHostReverseProxy(currencyURL),
		notifier: httputil.NewSingleHostReverseProxy(notifierURL),
	}
}

func registerNotifierHandlers(p *Proxy) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.notifier.ServeHTTP(w, r)
	})

	http.Handle("/api/subscribe", logHttp(handler))
	http.Handle("/api/sendEmails", logHttp(handler))
}

func registerCurrencyHandlers(p *Proxy) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.currency.ServeHTTP(w, r)
	})
	http.Handle("/api/rate", logHttp(handler))
}

func logHttp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := logrus.WithFields(logrus.Fields{
			"Method": r.Method,
			"URI":    r.RequestURI,
		})

		l.Info()

		next.ServeHTTP(w, r)
	})
}