package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Request struct{
	Msg string
}

type Response struct{
	Msg string
}

func main(){
	var res Response
	if dial, err := net.Dial("tcp", "127.0.0.1:2333"); err == nil{
		client := rpc.NewClient(dial)
		if err := client.Call("Dog.Speak", Request{"call1"}, &res); err != nil{
			panic(err)
		}

		fmt.Println(res)

		if err := client.Call("CatPointer.Run", Request{"call2"}, &res); err != nil{
			panic(err)
		}
		fmt.Println(res)

		_ = client.Close()
	}
}
