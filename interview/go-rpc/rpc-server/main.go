package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)
type Dog struct {

}
type Request struct{
	Msg string
}

type Response struct{
	Msg string
}

func (d Dog)Speak(input Request, output *Response)(err error){
	fmt.Println("wang!")
	output.Msg = fmt.Sprintf("ok:%s", input.Msg)
	return
}

type Cat struct{}
func (cp *Cat)Run(input Request, output *Response)(err error){
	fmt.Println("run")
	output.Msg = fmt.Sprintf("ok:%s", input.Msg)
	return
}

func main() {
	var d Dog
	var c Cat
	if listen, err := net.Listen("tcp", "127.0.0.1:2333"); err == nil{
		s := rpc.NewServer()
		if err := s.Register(d); err != nil{
			log.Println(err)
		}
		if err := s.RegisterName("CatPointer", &c); err != nil{
			log.Println(err)
		}
		s.Accept(listen)

	}else{
			panic(err)
	}

}
