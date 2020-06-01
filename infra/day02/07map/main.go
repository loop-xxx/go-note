package main

import "fmt"

func main() {

	fmt.Println("hello, loop")
	var studentTable map[string]uint
	if studentTable != nil {
		studentTable["loop"] = 0 //assignment to entry in nil map
	}

	/**
	go 语言map类型
	map类型变量的本质:
		struct map{
			TREE* data ;
		};
	推测:
		底层是红黑树
		使用make()指定capcity时:
		(1)分配capactiy块的内存, 用于维护一个NODE容器,容器只是用来储存, 并不描述结构.
		(2)使用容器中的NODE构建树, 用于描述NODE节点之间的关系

		studentTable["loop"] = 0 //该执行操作, 初始化NODE容器中的一个节点,并建立关系挂载到树
	*/
	studentTable = make(map[string]uint, 10) //设置map的capcity
	studentTable["loop"] = 0
	fmt.Println(len(studentTable), studentTable)

	/**
	map数据类型 = 赋值运算符, 在map变量进行赋值操作时使用浅copy
	*/
	var st = studentTable
	fmt.Println(len(st), st)
	st["loop"] = 0x10
	st["luna"] = 0x11
	fmt.Println(len(studentTable), studentTable)

	for kay, val := range st {
		fmt.Println(kay, val)
	}

	/**
	map 动态扩容
	*/
	var dynamicMap = make(map[int][4]int, 0x1)
	dynamicMap[0] = [4]int{0x1, 0x2, 0x3, 0x4}
	//fmt.Printf("%p, %p\n", dynamicMap, &dynamicMap[0])
	/**
	自动动态扩容, map中用于储存的NODE容器可能会被扩展另一块更大的内存地址
	&dynamicMap[0]两次打印的值可能会不同, 但go语言不允许打印map容器中储存的变量的地址
	*/
	dynamicMap[1] = [4]int{0xff, 0xff, 0xff, 0xff}
	//fmt.Printf("%p, %p\n", dynamicMap, &dynamicMap[0])

	fmt.Printf("%#v\n", dynamicMap)

	var variable map[int]string
	variable = map[int]string{0: "loop", 1: "luna"}
	fmt.Printf("%#v", variable)
	variable[2] = "xi"
	fmt.Printf("%#v", variable)

}
