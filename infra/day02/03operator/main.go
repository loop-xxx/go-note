package main

import "fmt"

func main() {
	fmt.Println("hello, loop")

	var (
		number  = uint64(0x10)
		number1 = uint32(0x10)
		number2 = uint32(010)
	)

	fmt.Println(number + uint64(number1)) //go 语言不支持隐式类型转换
	fmt.Println(number1 - number2)
	fmt.Println(number1 * number2)
	fmt.Println(number1 / number2)
	fmt.Println(number1 % number2)

	if number == uint64(number1) {
		fmt.Println("number == number1")
	}

	if number != uint64(number1) {
		fmt.Println("number != number1")
	}

	if number1 >= number2 {
		fmt.Println("number1 >= number2")
	}

	if age := uint8(25); age >= 20 && age < 35 {
		fmt.Println("I belive magic")
	}

	if age := uint8(19); age < 20 || age >= 35 {
		fmt.Println("I don't kown")
	}

	if age := uint8(19); !(age == 25) {
		fmt.Println("Do you belive magic?")
	}

	var flag = uint64(0x1010)
	fmt.Println(flag | 0x10100000)
	fmt.Println(flag & 0xf)

	flag = 0x3
	fmt.Println(flag ^ 0x7)

	var n = int32(-4)
	fmt.Println(n << 2)
	fmt.Println(n >> 1)

	/*
		go 语言中的= += -=...不能再次作为=的左值或右值使用
	*/

	var temp1 = uint32(1)
	var temp2 = uint32(2)

	//temp1 = (temp2 += 2)
	temp2 += 2
	temp1 = temp2
	fmt.Println(temp1, temp2)

	//go语言 后置++,--只能作为单独的语句使用, 不能作为=的左值或右值, 相当于go语言中的variable += 1
	number++
	number--

	/**
	有符号数和无符号数有区别的几种情况:
	(1)(/)除法, (%)取余
	(2) (>>) 左移运算
	(3) 关系表达 正数和负数进行(>)(<)(>=)(<=)判断时
	(4) 类型上升转换时, 例如:
		int8类型的-1转换为uint32类型后为0xffffffff
		uint8类型的-1转换为uint32类型后为0xff
	*/

	var char1 = uint8(0xff)
	var char2 = int8(char1)
	var word uint32

	word = uint32(char1)
	fmt.Printf("%#x, %#x\n", char1, word)
	word = uint32(char2)
	fmt.Printf("%#x, %#x\n", char2, word)
}
