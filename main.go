package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

var servers = []string{
	"http://localhost:3000",
	"http://localhost:3001",
	"http://localhost:3002",
	"http://localhost:3003",
}

var nextServer = 0
var mutex sync.Mutex

func handler(rw http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	rpUrl, _ := url.Parse(servers[nextServer%len(servers)])
	fmt.Println(nextServer)
	nextServer = (nextServer + 1) % len(servers)
	mutex.Unlock()
	proxy := httputil.NewSingleHostReverseProxy(rpUrl)
	proxy.ServeHTTP(rw, req)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
