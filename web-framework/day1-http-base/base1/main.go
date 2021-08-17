package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatalln(http.ListenAndServe("127.0.0.1:9000", nil))
	/**
	第一个参数是地址，:9999表示在 9999 端口监听。
	第二个参数则代表处理所有的HTTP请求的实例，nil 代表使用标准库中的实例处理。
	第二个参数，则是我们基于net/http标准库实现Web框架的入口。
	*/
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.Path = %s \n", request.URL.Path)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	for s, strings := range request.Header {
		fmt.Fprintf(writer, "Header[%q] = %q \n", s, strings)
	}
}
