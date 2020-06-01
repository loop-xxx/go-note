package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

//返回值命名, 简写返回语句
func func1(a int, b int) (ret int) {
	ret = a + b
	return
}

//返回多个值
func func2() (int, string) {
	return 1, "hello,loop"
}

//参数类型简写
func func3(x, y, z int, prevfix, suffix string) {

}

//可选参数
func func4(text string, arr ...int) {
	fmt.Println(text)
	fmt.Printf("%#v\n", arr)
}

func selectSort(arr []int) {
	for i := int(0); i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			arr[min] = arr[min] ^ arr[i]
			arr[i] = arr[min] ^ arr[i]
			arr[min] = arr[min] ^ arr[i]
		}
	}
}

func bubbleSort(arr []int) {
	for i := int(1); i < len(arr); i++ {
		for j := int(0); j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j] = arr[j] ^ arr[j+1]
				arr[j+1] = arr[j] ^ arr[j+1]
				arr[j] = arr[j] ^ arr[j+1]
			}
		}
	}
}

func nonameRetVar(a int, b int) int {
	var ret = a + b
	defer func() {
		if ret > 0 {
			panic("name return variable")
		}
	}()
	return ret
}

func nameRetVar(a int, b int) (ret int) {
	ret = a + b
	defer func() {
		if ret > 0 {
			panic("name return variable")
		}
	}()
	return
}

func main() {
	fmt.Println("hello, loop")

	sum(1, 1)
	_, str := func2()
	fmt.Println(str)
	func4("hi", 1, 2, 3, 4)
	func4("hello")

	var arr = [4]int{11, 24, 0x10, 0}
	bubbleSort(arr[:])
	fmt.Println(arr)

	selectSort(arr[:])
	fmt.Println(arr)

	/**
	无论是否对返回值命名, 都存在一个用作记录返回值的局部变量
	*/
	nonameRetVar(1, 1)
	//nameRetVar(1, 1)
}
