package main

import (
	"fmt"
	"net"
)

func main() {
	if listener, err := net.Listen("tcp", "192.168.43.212:2333"); err == nil {
		//listener, _ = net.Listen("tcp", "127.0.0.1:2333") //可以同时监听两个网卡的相同端口
		for {
			if conn, err := listener.Accept(); err == nil {
				fmt.Println("listen_ok")
				conn.Close()
			} else {
				fmt.Printf("accept faild err:%v", err)
			}
		}

	} else {
		fmt.Printf("listen failed err:%v", err)
	}
}
