package main

import "fmt"

var gnumber = uint32(0xff)
var gn = ninit()

func ninit() (ret uint64) {
	fmt.Println("nint")
	ret = 0xffff
	return
}

func init() {
	fmt.Println("init", gnumber)
}

func main() {
	fmt.Println("hello, loop")
}
