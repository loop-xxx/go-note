package main

import (
	"fmt"
	"net"
)

func network() {
	if conn, err := net.Dial("tcp", "192.168.43.212:2333"); err == nil {
		fmt.Println("conn_ok")
		conn.Close()

	} else {
		fmt.Printf("Dial failed %v\n", err)
	}
}
func local() {
	if conn, err := net.Dial("tcp", "127.0.0.1:2333"); err == nil {
		fmt.Println("conn_ok")
		conn.Close()

	} else {
		fmt.Printf("Dial failed %v\n", err)
	}
}

func main() {
	network()
}
