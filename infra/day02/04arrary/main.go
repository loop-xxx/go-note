package main

import "fmt"

/**
go 语言中array类型对传统c语言array的概念进行了重新定义,不再是c语言中的数组名代表数组首元素地址, 数组名不再 ≈ 指针
数组名成为了真正的数组变量,代表了整个数组所有元素的数据, 同类型的数组变量之间可以互相赋值
go语言中的数组像是c语言中的
struct arr_LEN{
	TYPE data[LEN];
};
*/
func main() {
	fmt.Println("hello, loop")
	var arr1 [0x10]int
	var arr2 [0x11]int
	var barr1 [0x4]bool
	var barr2 [0x5]bool

	/**
	未初始化变量,go默认将其初始化为0
	*/
	fmt.Printf("%T,%T,%T,%T\n", arr1, arr2, barr1, barr2)
	fmt.Println(arr1, arr2, barr1, barr2)

	//一般初始化/赋值
	var arr3 = [4]int{1, 2, 3, 4}
	fmt.Printf("%T, %#v\n", arr3, arr3)
	//初始化/赋值 明确写出部分元素的值, 未明确写出的其他元素, 默认为0
	arr3 = [4]int{2, 3}
	fmt.Printf("%T, %#v\n", arr3, arr3)

	//初始化/赋值 明确写出指定索引元素的值, 未明确写出的其他元素, 默认为0
	var arr5 = [4]int{0: 5, 3: 5}
	fmt.Printf("%T, %#v\n", arr5, arr5)

	//初始化 不明确写出元素个数
	var arr6 = [...]int{1, 2, 3, 4}
	fmt.Printf("%T, %#v\n", arr6, arr6)

	for i := int(0); i < len(arr6); i++ {
		fmt.Printf("%d ", arr6[i]) //支持[]运算符, 其底层为内建函数
	}
	fmt.Println()

	for index, element := range arr6 {
		fmt.Println(index, element)
	}

	var arr7 [4][4]int
	arr7 = [4][4]int{[4]int{1, 2}, [4]int{3, 4}, [4]int{5, 6}, [4]int{7, 8}}
	fmt.Println(arr7)

	//len() 内建函数支持
	for i := int(0); i < len(arr7); i++ {
		for j := int(0); j < len(arr7[i]); j++ {
			fmt.Printf("%d ", arr7[i][j])
		}
		fmt.Println()
	}

	//for range 遍历, 底层为内建函数支持
	for _, v := range arr7 {
		for _, e := range v {
			fmt.Printf("%d ", e)
		}
		fmt.Println()
	}

	//数组变量支持 == != 关系运算符, 底层为内建函数支持
	var va1 = [2]uint32{1, 2}
	var va2 = va1
	va2[0] = 100
	fmt.Println(va1)
	fmt.Println(va2)
	if va2 != va1 {
		fmt.Println("hello, go")
	}

	var data = [...]int{1, 3, 5, 7, 8}
	var sum int //默认初始化为0
	for i := int(0); i < len(data); i++ {
		sum += data[i]
	}
	fmt.Println(sum)

	for i := int(0); i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
}
