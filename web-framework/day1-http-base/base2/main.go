package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path : %s \n", req.URL.Path)
	case "/hello":
		for s, strings := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q \n", s, strings)
		}
	default:
		fmt.Fprintf(w, "Not Found %s \n", req.URL.Path)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	engine := &Engine{}
	log.Fatalln(http.ListenAndServe("127.0.0.1:9000", engine))
	/**
	此处就不走indexHandler()、helloHandler()了。
	将所有的HTTP请求转向了我们自己的处理逻辑
	在实现Engine之后，我们拦截了所有的HTTP请求，拥有了统一的控制入口。
	*/

}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "你好 %s \n", request.URL.Path)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, " 你好 %s \n", request.URL.Path)
}
