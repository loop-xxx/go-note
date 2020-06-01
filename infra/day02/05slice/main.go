package main

import "fmt"

func main() {
	/**
	go语言slice类型,slice的本质以c语言表示
	struct slice{
		size_t cap ;
		size_t len ;
		TYPE *data ;
	}
	*/
	//切片声明并默认初始化
	var slice1 []int
	/**
	切片不支持 == 关系运算符, 但是go语言提供一种语法来判断slice类型变量是否分配内存,
	slice1 == nil --> slice1.data == nil

	[]TYPE {} != nil
	make([]TYPE, 0, 0) != nil
	*/
	if slice1 == nil {
		fmt.Printf("%d, %d, %#v\n", len(slice1), cap(slice1), slice1) //len 内建函数支持
		fmt.Println("slice is not init")
	}

	/**
	创建var arr = [4]int{1,2,3,4}
	slice.len = 4
	slice.cap = 4
	slice.data = &arr[0]
	*/
	slice1 = []int{1, 2, 3, 4}

	/**
	  slice类型 = 赋值运算符的操作和string类型相同,使用浅copy的方式, 这也受益于gc机制的好处,
	  浅copy可以达到c语言数组传址的效果, 弥补了go [...]TYPE array类型不能传址问题
	*/
	var slice2 = slice1
	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Printf("%p, %p\n", &slice1[0], &slice2[0])

	/*
		由数组得到切片
		silice = arr[a:b] 翻译为c ==>\
		slice.len = b-a ;
		slice.cap = arr.len - a;
		slice.data = &arr.data[a] ;
		底层为内建函数实现
	*/
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice3 = arr[1:3]

	slice3[0] = 0x400
	fmt.Println(slice3, len(slice3), cap(slice3), arr)

	slice3 = arr[:5] // [0:5]
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice3 = arr[:] //[0:len(arr)]
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice3 = arr[5:] //[5:len(arr)]
	fmt.Println(slice3, len(slice3), cap(slice3))

	/**
	由切片得到切片
	slice2 = slice1[a:b] 转换为c代码==>
	slice2.len = b-a ;
	slice2.cap = slice1.cap - a ;
	slice2.data = &slice1.data[a] ;
	底层为内建函数实现
	*/
	var slice4 = slice3[2:]
	slice4[0] = 0x1000
	fmt.Println(slice4, len(slice4), cap(slice4), slice3)

	/**
	创建 var arr = [10]int{0, 0}
	slice.len = 2
	slice.cap = 10
	slice.data = &arr[0]
	*/
	var slice5 = make([]int, 2, 10)
	fmt.Println(slice5, len(slice5), cap(slice5))

	/**
	slice = 赋值运算符默认使用浅copy, 如果要进行深拷贝,
	使用copy函数, copy(dest, src) 拷贝个元素MIN(dest.len, src.len)的数据到dest.data指向的内存地址
	*/
	var slice6 = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	var slice7 = make([]int, 0x10, 0x10)
	copy(slice7, slice6)
	fmt.Printf("%p:%#v, %p:%#v\n", &slice6[0], slice6, &slice7[0], slice7)
	copy(slice7, arr[:])
	fmt.Printf("%p:%#v, %p:%#v\n", &arr[0], arr, &slice7[0], slice7)

	/**
	append追加元素, 追加元素n
	当n <= slice.cap - slice.len时, 返回的slice变量引用的原先的内存地址
	当n > slice.cap-slice.len时, 进行capcity扩展, 返回的slice变量引用扩展后新的内存地址
	*/
	var slice8 = make([]int, 2, 0x10)
	var slice9 = append(slice8, 0x10, 0x11)
	fmt.Printf("%p:%#v, %p:%#v\n", slice8, slice8, slice9, slice9)
	slice8 = append(slice9, arr[:]...)

	var slice10 = append(slice8, slice6...)
	fmt.Printf("%p:%#v, %p:%#v\n", slice10, slice10, slice8, slice8)

	/**
	使用切片分解和组装string(ps:只是将string的每个字符拷贝出来, 修改后重新生成一个新的string变量,
	不会修改到原始string变量副本引用数据)
	*/
	var str = string("loop")
	var runeSlice = []rune(str)
	runeSlice[2] = 'O'
	fmt.Println(str, string(runeSlice))

	//fmt.Printf("%p, %p, %p, %p, %p, %p, %p, %p, %p, %p", &slice1, &slice2, &slice3, &slice4, &slice5, &slice6, &slice7, &slice8, &slice9, &slice10)
}
