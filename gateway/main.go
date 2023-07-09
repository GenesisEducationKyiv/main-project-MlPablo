package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Define the backend servers
	currencyURL, _ := url.Parse("http://currency:8001")
	notifierURL, _ := url.Parse("http://notifier:8002")

	// Create the reverse proxy handler
	proxyC := httputil.NewSingleHostReverseProxy(currencyURL)
	proxyN := httputil.NewSingleHostReverseProxy(notifierURL)

	// Handle requests and route them to the appropriate backend server
	http.HandleFunc("/api/rate", func(w http.ResponseWriter, r *http.Request) {
		proxyC.ServeHTTP(w, r)
	})
	http.HandleFunc("/api/subscribe", func(w http.ResponseWriter, r *http.Request) {
		proxyN.ServeHTTP(w, r)
	})
	http.HandleFunc("/api/sendEmails", func(w http.ResponseWriter, r *http.Request) {
		proxyN.ServeHTTP(w, r)
	})

	// Start the reverse proxy server on port 80
	log.Fatal(http.ListenAndServe(":8080", nil))
}
