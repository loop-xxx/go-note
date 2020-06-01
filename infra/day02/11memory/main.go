package main

import "fmt"

/**
golang不允许程序员操作函数栈, 函数中定义的所有变量全部创建在gc机制维护的运行时堆中,哪怕是一个基本类型(如:uint8, bool)的变量, 其内存也分配在运行时堆中, 被gc机制管理释放
golang 比起"21世纪的c语言"这个比喻, 个人感觉"编译型的python"更贴切一点
*/
func clearFuncStack() {
	var (
		n1 = uint64(0)
		n2 = uint64(0)
		n3 = uint64(0)
		n4 = uint64(0)
	)
	fmt.Println(n1, n2, n3, n4)
}

func retBaseTypePointer() *uint32 {
	var number = uint32(0xffffffff)
	return &number
}

func retArrPointer() *[4]int {
	var arr = [4]int{0x11, 0x12, 0x13, 0x14}
	return &arr
}

func retSlicePointer() *[]int {
	var slice = []int{0x1, 0x2, 0x3, 0x4}
	return &slice
}

func retStringPointer() *string {
	var text = string("hello, world")
	return &text
}

func main() {
	fmt.Println("hello, loop")
	var temp uint64

	var uintptr *uint32
	uintptr = retBaseTypePointer()
	clearFuncStack()
	fmt.Printf("&temp:%p,&uniptr:%p, uniptr:%p, *uniptr:%#v\n", &temp, &uintptr, uintptr, *uintptr)

	var arrptr *[4]int
	arrptr = retArrPointer()
	clearFuncStack()
	fmt.Printf("&temp:%p,&arrptr:%p, arrptr:%p, *arrptr:%#v\n", &temp, &arrptr, arrptr, *arrptr)

	var sliceptr *[]int
	sliceptr = retSlicePointer()
	clearFuncStack()
	fmt.Printf("&temp:%p,&sliceptr:%p, sliceptr:%p, *slliceptr:%#v\n", &temp, &sliceptr, sliceptr, *sliceptr)

	var strptr *string
	strptr = retStringPointer()
	clearFuncStack()
	fmt.Printf("&temp:%p,&strptr:%p, strptr:%p, *strptr:%#v\n", &temp, &strptr, strptr, *strptr)
}
