package main

import "fmt"

func main() {
	var ch1 = 'a'
	var ch2 = '你'

	fmt.Printf("%T,%T\n%#v, %#v\n", ch1, ch2, ch1, ch2)

	//go 语言中的字符类型就是rune类型, rune是int32的别名
	var ch1a rune = 'a'
	var ch2a rune = '你'
	fmt.Printf("%T,%T\n%#v,%#v\n", ch1a, ch2a, ch1a, ch2a)

	/*
		byte是uint8的别名, string类型在保存字符串时,为了节省空间没有直接使用rune类型, 而是使用byte类型,
		utf-8 ascii码字符 每个占1个字节, 使用1个byte保存
		utf-8 中文字符 每个占3个字节,使用3个byte保存
	*/
	var ch1b byte = 'a'
	//var ch2b byte = '你' //中文字符一个字节无法存储
	fmt.Printf("%T\n%#v\n", ch1b, ch1b)

	var str = string("hello, world 你好世界")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println("\n-----------------")

	/**
	使用for range遍历, 会将每个字符统一赋值保存在rune类型的变量character中, 即使是只需一个字节存储的ascii码字符
	*/
	for index, character := range str {
		fmt.Printf("%#v:%T(%c); ", index, character, character)
	}

	var prev = int(0)
	var total = int(0)
	for i := range str {
		if i-prev > 1 {
			total++
		}
		prev = i
	}

	if len(str)-prev > 1 {
		total++
	}

	fmt.Println(total)
}
