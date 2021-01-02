package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	urlPath, _ := url.Parse("http://localhost:8083/server2")
	proxy := httputil.NewSingleHostReverseProxy(urlPath)

	http.HandleFunc("/server1", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL)
		proxy.ServeHTTP(writer, request)
		fmt.Println(request.URL)
	})

	http.ListenAndServe(":8081", nil)

	ch := make(chan int)
	<-ch
}
