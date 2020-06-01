package main

import "fmt"

/**
匿名函数
*/

var vfunc = func(a, b uint32) (ret uint32) {
	ret = a + b
	fmt.Println("func(a, b uint32) (ret uint32)")
	return ret
}

func main() {
	fmt.Println("hello, loop")
	fmt.Println(vfunc(0x2, 0x4))
	//匿名函数可以在函数内部定义
	var vf = func(number uint32) {
		fmt.Println(number)
	}
	vf(0x10)

	//匿名函数 立即调用
	func(msg string) {
		fmt.Println(msg)
	}("success")
}
