package main

import "fmt"

/**
第三方的接口
*/
func addCallback(callback func()) {
	callback()
}

/**
需要设置的回调函数
*/
func loopCallback1(a int, b int) {
	fmt.Println("loop callback 1")
	fmt.Println(a + b)
}
func loopCallback2(a int, b int) {
	fmt.Println("loop callback 2")
	fmt.Println(a + b)
}
func loopCallback3(a int, b int) {
	fmt.Println("loop callback 3")
	fmt.Println(a + b)
}

/**
闭包 :函数内部使用外部作用域的变量, 并且外部作用域不是全局作用域
疑问一:使用到外部的变量也是外部函数的局部变量, 虽然变量的生命周期会由gc机制管理, 但是函数(包括匿名函数)的生命周期是整个程序,
也就是一直到程序结束,函数都有可能被调用, 那么gc怎么去判定这个变量已经没有用了可以被回收
*/
func closePackage() func() {
	var number = int(0x10)
	var temp1 int
	var temp2 int
	fmt.Printf("%p, %#v\n", &number, number)
	fmt.Printf("&temp1:%p, &temp2:%p\n", &temp1, &temp2)
	return func() {
		fmt.Printf("%p, %#v\n", &number, number)
		number += 0x10
		fmt.Printf("%p, %#v\n", &number, number)
		fmt.Println(number) //使用外部作用域的nunber变量
	}
}

func adapter(raw func(int, int), arg1 int, arg2 int) func() {
	return func() {
		raw(arg1, arg2)
	}
}

func main() {
	fmt.Println("hello, world")

	var f = closePackage()
	f()
	f()
	fmt.Println("---------------------------")
	/**
	  函数类型的参数, 类型不匹配的问题
	*/
	/**
	解决方式一, 使用匿名函数包装回调函数, 对于多个不匹配的函数, 需要定义多个匿名函数
	修改传入的参数也必须再次定义一个匿名函数
	*/
	addCallback(func() {
		loopCallback1(1, 1)
	})
	addCallback(func() {
		loopCallback1(2, 2)
	})
	addCallback(func() {
		loopCallback2(2, 2)
	})
	addCallback(func() {
		loopCallback3(2, 2)
	})
	fmt.Println("---------------------------")
	/**
	解决方式二 闭包
	*/
	var a1 = adapter(loopCallback1, 1, 1)
	var a2 = adapter(loopCallback1, 2, 2)
	var a3 = adapter(loopCallback2, 2, 2)
	var a4 = adapter(loopCallback3, 2, 2)

	addCallback(a1)
	addCallback(a2)
	addCallback(a3)
	addCallback(a4)
	/**
	疑问二:　a1-a4每个变量都应该对应一个闭包, 但是这些变量类型的大小为0x8个字节,0x8字节为一个地址, 指向同一个闭包函数,
	调用闭包, 闭包函数如何索引到, 当前闭包应该使用的外部变量
	*/
	fmt.Printf("%T, %p, &:%p, %T, %p, &:%p, %T, %p, &:%p, %T, %p, &:%p\n",
		a1, a1, &a1, a2, a2, &a2, a3, a3, &a3, a4, a4, &a4)
}
