package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	rpUrl, _ := url.Parse("http://localhost:3000")
	proxy := httputil.NewSingleHostReverseProxy(rpUrl)
	proxy.ServeHTTP(rw, req)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
