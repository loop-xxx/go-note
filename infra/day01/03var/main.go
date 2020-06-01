package main

import "fmt"

var gnumber int            //变量声明
var gname1 string = "loop" //变量声明并初始化(编译器推荐使用自动类型推导)
var gname2 = "ljc"         //变量声明并初始化,使用自动类型推导

var ( //多变量声明
	gchild1 int
	gchild2 int = 2 //变量声明并初始化(编译器推荐使用自动类型推导)
	gchild3     = 3 //变量声明并初始化,使用自动类型推导
)

func main() {
	var number1 int       //变量声明
	var number2 int16 = 2 //变量声明并初始化
	var number3 = 3       //变量声明并初始化,使用自动类型推导
	number4 := 4          //局部变量支持(变量声明并初始化,使用自动类型推导)的简写

	number1 = 1

	var ( //多变量声明
		child1 int       //未初始化变量,默认值为0
		child2 int32 = 2 //变量声明并初始化
		child3       = 3 //变量声明并初始化,使用自动类型推导
	)

	fmt.Println("hello, loop")
	fmt.Printf("%d, %d, %d ,%d\n", number1, number2, number3, number4)
	fmt.Printf("%d, %d, %d\n", child1, child2, child3)
}
