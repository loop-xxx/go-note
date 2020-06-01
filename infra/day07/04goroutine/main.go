package main

import (
	"fmt"
	"time"
)

func f(i int) {
	fmt.Println(i)
}

func test01() {
	for i := 0; i < 0x100; i++ {
		go f(i)
	}

	//time.Sleep(time.Microsecond * 0x1000)
	time.Sleep(time.Second)
}
func test02() {
	for i := 0; i < 0x100; i++ {
		go func() { //使用闭包, 闭包索引的外部变量的值, 在闭包函数外被修改
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}

func test03() {
	for i := 0; i < 0x100; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(time.Second)
}
func main() {
	test03()
}
