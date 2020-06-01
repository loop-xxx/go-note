package main

import (
	"fmt"
	"strconv"
)

func main() {
	if n1, err := strconv.Atoi("1111"); err == nil {
		fmt.Printf("%T, %v\n", n1, n1)
	} else {
		fmt.Println(err)
	}
	{
		str := strconv.Itoa(0x1000)
		fmt.Printf("%s\n", str)
	}

	//解析16位时不支持前缀0x
	if n2, err := strconv.ParseInt("-400", 16, 64); err == nil {
		fmt.Printf("%T, %v\n", n2, n2)
	} else {
		fmt.Println(err)
	}
	//解析16位时支持前缀0x
	if n3, err := strconv.ParseUint("400", 16, 64); err == nil {
		fmt.Printf("%T, %v\n", n3, n3)
	} else {
		fmt.Println(err)
	}
	if n4, err := strconv.ParseUint("0755", 8, 64); err == nil {
		fmt.Printf("%T, %v\n", n4, n4)
	} else {
		fmt.Println(err)
	}
	if b1, err := strconv.ParseBool("true"); err == nil {
		fmt.Printf("%T, %v\n", b1, b1)
	} else {
		fmt.Println(err)
	}

}
