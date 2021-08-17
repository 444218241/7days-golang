package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Fatal("Failed", err)
	}
}

func TestConn(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:9999")
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)
	go http.Serve(ln, nil)

	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handleError(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "您好" {
		t.Fatal("expected hello world, but got", string(body))
	}
}
