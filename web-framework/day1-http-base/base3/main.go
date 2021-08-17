package main

import (
	"7days-golang-web-framework/day1-http-base/base3/gee"
	"fmt"
	"net/http"
)

func main() {

	r := gee.New()
	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run("127.0.0.1:9999")
}
