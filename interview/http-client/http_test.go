package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)


func TestHttp(t *testing.T){
	clients := [0x10]http.Client{}
	//所有的http.client对象共享使用http包底层的链接池
	for i := int(0); i < len(clients); i++{
		if res, err := clients[i].Get("https://baidu.com"); err == nil{
			fmt.Println(ioutil.ReadAll(res.Body)) //res.Body 被读取完毕, 此次请求已经结束,将连接重新放回到连接池,以备复用.
			time.Sleep(time.Second*4)
			//如果主动调用res.Body.Close()方法, 不会将链接放回连接池, 会将连接直接关闭.
			//res.Body.Close()
		}
	}
}

func TestHttp2(t *testing.T) {
	tr := http.Transport{DisableKeepAlives: true}
	client := http.Client{Transport: &tr} //设置为短连接后, http服务端和客户端都不会保持连接

	for i := int(0); i < 0x10; i++{
		if res, err := client.Get("https://baidu.com"); err == nil{
			fmt.Println(res.Body)
			time.Sleep(time.Minute)
			fmt.Println(ioutil.ReadAll(res.Body)) //res.Body 被读取完毕, 此次请求已经结束, 因为设置了短链接, 链接不会被放回链接池, 而是直接被关闭.
		}
	}
}