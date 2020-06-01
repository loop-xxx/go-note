package main

import "fmt"

func main() {
	var flag bool
	var uchar uint8 = 0xff
	/* go语言总bool类型和int类型不能互相转换
	flag = bool(number)
	number = uint8(flag)
	*/
	var short int16
	short = int16(uchar) //go 语言不存在隐式类型转换,所有的类型转换必须显示写出

	fmt.Println(flag, uchar, short)
	fmt.Println("hello, loop")
	/*
		%T 打印数据的类型
		%v 万能占位符, 打印变量的值
		%#v 以go 语言的语法格式显示变量的值
		%+v 打印结构体变量时, 显示字段名
	*/
	fmt.Printf("%T\n", short+int16(uchar)) //go 语言不存在隐式类型转换

	var llu uint64 = 0xffffffffffffffff
	fmt.Printf("%T, %v\n", llu, llu)
	fmt.Printf("%T, %#v\n", llu, llu)

	//var n1, n2 uint8 = 1, 2  //只能指定一个变量类型
	var n1, n2 = uint8(1), uint16(2) //指定两个变量类型
	fmt.Printf("%T,%T, %#v, %#v\n", n1, n2, n1, n2)

	for i := uint8(0); i < uint8(0xff); i++ {
		fmt.Printf("%T, %#v\n", i, i)
	}
}
