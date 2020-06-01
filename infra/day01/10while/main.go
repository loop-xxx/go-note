package main

import (
	"fmt"
)

func main() {
	//fmt.Println("hello, loop")

	var age = int(19)
	if age > 30 {
		fmt.Println("welcome, while")
	} else if age > 20 {
		fmt.Println("welcom, loop")
	} else {
		fmt.Println("please go back")
	}

	if age := int(25); age > 30 {
		fmt.Println("hello, while")
	} else if name := string("loop"); age > 20 {
		fmt.Println("hello, ", name)
	} else {
		fmt.Println("please go back")
	}

	for i, name := int(0), string("hello"); i < 10; i++ {
		fmt.Println(name)
	}

	for i := int(0); i < 10; i++ {
		fmt.Println(i)
	}

	var i = int(0)
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("----------------------")
	var ii = int(0)
	for ii < 10 {
		fmt.Println(ii)
		ii++
	}

	var str = string("hello,loop")
	for i, c := range str { //(6)range str 其底层实现为内建函数
		fmt.Printf("[%d]=%c ", i, c)
	}
	fmt.Println("\n----------------------------")

	//break
	for i := int(0); i < 0x10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("%T(%d) ", i, i)
	}
	fmt.Println()

	for i := uint8(0); i < 0xf; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("%T(%d) ", i, i)
	}
	//死循环
	for {
	}

}
