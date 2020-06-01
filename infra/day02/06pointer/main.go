package main

import "fmt"

func print(ptext *string) {
	fmt.Println(*ptext)
	*ptext = "hello, loop"
}

func main() {
	var intptr *int

	if intptr == nil {
		fmt.Println(intptr)
	}
	intptr = new(int)
	*intptr = 10
	fmt.Println(intptr, *intptr)

	var number = int(0x10)
	intptr = &number
	*intptr = 0x11
	fmt.Println(intptr, number)

	var text = string("hello, world")
	print(&text)
	fmt.Println(text)
}
