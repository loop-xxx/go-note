package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type firstHandle struct {
}

func (f firstHandle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if data, err := ioutil.ReadFile("./index.html"); err == nil {
		rw.Write(data)
	} else {
		rw.Write([]byte("<h1>File not found!</h1>"))
	}
}

func secondHandleFunc(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("-----------------------------")
	fmt.Println(r.URL)
	fmt.Println(r.URL.Path)
	fmt.Println(r.Method)
	values := r.URL.Query()
	fmt.Printf("name:%s\n", values.Get("name"))

	if data, err := ioutil.ReadAll(r.Body); err == nil {
		fmt.Printf("%#v\n", data)
	}

	if data, err := ioutil.ReadFile("./index2.html"); err == nil {
		rw.Write(data)
	} else {
		rw.Write([]byte("<h1>File not found!</h1>"))
	}
}
func main() {
	http.Handle("/first", firstHandle{})
	http.HandleFunc("/second", secondHandleFunc)
	if err := http.ListenAndServe("127.0.0.1:2333", nil); err == nil {
	} else {
		fmt.Println(err)
	}

}
