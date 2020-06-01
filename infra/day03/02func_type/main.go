package main

import "fmt"

func f1(x, y int) bool {
	fmt.Println("hello, world")
	return true
}
func f2Sub(text string) int {
	return 0
}

func f2(f func(int, int) bool) func(string) int {
	return f2Sub
}

func main() {
	fmt.Println("hello loop")
	var vf1 = f1
	var vf2 = f1
	fmt.Printf("%T, %p\n", f1, f1) //&f1不允许
	fmt.Printf("%T, %p ,&:%p\n", vf1, vf1, &vf1)
	fmt.Printf("%T, %p ,&:%p\n", vf2, vf2, &vf2)

	var v1 = int(0x10)
	var v2 = int(0x10)
	fmt.Printf("%T, %v\n", int(0x10), int(0x10)) //&int(0x10)不允许
	fmt.Printf("%T, %v ,&:%p\n", v1, v1, &v1)
	fmt.Printf("%T, %v ,&:%p\n", v2, v2, &v2)

	fmt.Printf("%T, %p\n", f2, f2)
}
