package main

import "fmt"

func calc(base uint64) (add func(uint64) uint64, sub func(uint64) uint64) {
	add = func(number uint64) uint64 {
		base += number
		return base
	}
	sub = func(number uint64) uint64 {
		base -= number
		return base
	}
	return
}

func main() {
	fmt.Println("hello, loop")
	add, sub := calc(0x11)

	fmt.Println(add(1), sub(2)) //18, 16
	fmt.Println(add(3), sub(4)) //19, 15
	fmt.Println(add(5), sub(6)) //20, 14
}
