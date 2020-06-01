package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func get() {
	//if rep, err := http.Get("http://127.0.0.1:2333/second?name=loop"); err == nil {
	if rep, err := http.DefaultClient.Get("http://127.0.0.1:2333/second?name=李嘉图"); err == nil {
		//defer rep.Body.Close()
		if data, err := ioutil.ReadAll(rep.Body); err == nil {
			fmt.Println(string(data))
		}
	} else {
		fmt.Println(err)
	}
}
func post() {
	url := "http://127.0.0.1:2333/second"
	data := `{"name":"映像研"}`
	if resp, err := http.Post(url, "application/json", strings.NewReader(data)); err == nil {
		//defer resp.Body.Close()
		if data, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Println(string(data))
		}
	} else {
		fmt.Println(err)
	}
}

func coustmURL() {

	urlObj, err := url.ParseRequestURI("http://127.0.0.1:2333/second")
	if err != nil {
		return
	}

	// url.Values 底层为map
	values := make(url.Values, 0x1)
	values.Set("名字", "映像研")
	fmt.Println(values.Encode())
	urlObj.RawQuery = values.Encode()
	fmt.Println(urlObj.String())

	if resp, err := http.Get(urlObj.String()); err == nil {
		//defer resp.Body.Close()
		if data, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Println(string(data))
		}
	} else {
		fmt.Println(err)
	}

}

func main() {
	get()
	//coustmURL()
}
