package main

import (
	"fmt"
	"strings"
)

/**
go string类型
go语言中的string类型像是c语言中的
struct string{
	size_t len ;
	char const * data ; //指向[...]byte类型的数组
};
*/
func main() {
	fmt.Println("hello, loop")

	var name string = "loop"
	fmt.Println(name, len(name))
	var content string = `
曾经沧海难为水
	除却巫山不是云
	"
	\/
	`
	fmt.Println(content)

	/*
		(1)string类型 = 赋值运算符, 推测使用浅copy, string类型不提供修改索引的[...]byte类型数组的方法,并且go语言实现了gc垃圾回收机制,
		不需要关心内存释放的问题,所以单纯的赋值操作,没有必要进行深copy,使用浅copy可以提高性能,节省内存空间
	*/

	var title = string("boss")

	//拼接字符串
	fmt.Println(fmt.Sprintf("%s:%s", name, title))
	//(2)string类型支持+运算符, 其底层实现为内建函数
	fmt.Println(name + ":" + title)

	//获取子串在字符串中的位置,不存在返回-1
	fmt.Println(strings.Index(name, "oss"), strings.LastIndex(name, "o"))

	//按指定子串分割
	var signal = string("hello, loop\nhello, linux\nhello go")
	var strList = strings.Split(signal, "\n")
	fmt.Println(strList)

	//是否包含子串
	fmt.Println(strings.Contains(title, "o"), strings.Contains(title, "l"))

	//连接字符串
	fmt.Println(strings.Join(strList, "\t"))

	//是否含有前缀后缀
	var str = string("head_dog_")
	fmt.Println(strings.HasPrefix(str, "head"), strings.HasSuffix(str, "tail"))

	//(3)string类型支持+=运算符, 其底层实现为内建函数
	str += "tail"
	fmt.Println(str)

	//(4)string类型支持[]运算符,返回索引在字符串中对应byte, 其底层实现为内建函数
	fmt.Printf("%T, %#v\n", str[0], str[0])

	//(5)len()内建函数, 返回string字符串的长度(以byte为单位)
	fmt.Printf("%T, %#v\n", len(str), len(str))

	//(6)range str 其底层实现为内建函数
	for i, c := range str {
		fmt.Printf("[%d]=%c ", i, c)
	}

	var str1 = string("hello")
	var str2 = string("hello,loop")
	var str3 string
	str1 += ",loop"
	//(7)string类型支持== !=运算符, 其底层实现为内建函数
	fmt.Println(str1 == str2)
	//fmt.Println(str3 == nil)
	//string类型变量大小: 0x10
	fmt.Printf("%p, %p, %p\n", &str1, &str2, &str3)
}
