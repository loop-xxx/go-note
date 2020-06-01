package main

import "fmt"

func main() {
	fmt.Println("hello, world")
	var number1 = 0777
	var number2 = uint16(0xff)
	number3 := int8(0xa) //int8('\n')
	var number4 uint64 = 0xffffffffffffffff

	fmt.Printf("%T,%T,%T,%T\n", number1, number2, number3, number4)
	fmt.Printf("%#v", number3)
}
