package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

func (f HandleFunc) ServerHttp(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func main() {
	hf := HandleFunc(HelloHandler)
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", bytes.NewBuffer([]byte("test")))
	hf.ServerHttp(resp, req)

	bts, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bts))
}

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world"))
}
