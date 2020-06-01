package main

import "fmt"

func main() {
	var stat = int(4)
	switch stat {
	case 1:
		fmt.Println("yi")
	case 2:
		fmt.Println("er")
	case 3:
		fmt.Println("san")
	case 4:
		fmt.Println("si")
	default:
		fmt.Println("unkonw")
	}

	switch stat := int(2); stat {
	case 1:
		fmt.Println("yi")
	case 2:
		fmt.Println("er")
	case 3:
		fmt.Println("san")
	case 4:
		fmt.Println("si")
	default:
		fmt.Println("unkonw")
	}

	switch stat := int(3); stat {
	case 1, 2, 3:
		fmt.Println("small")
	case 4:
		fmt.Println("big")
	default:
		fmt.Println("unkonw")
	}

	/**
	使用fallthrough关键字, 模仿c语言代码中不break的case选项
	*/
	switch stat := int(2); stat {
	case 1:
		fmt.Println("yi")
	case 2:
		fmt.Println("er")
		fallthrough
	case 3:
		fmt.Println("san")
	case 4:
		fmt.Println("si")
	default:
		fmt.Println("unkonw")
	}

	/**
	if else 的另一种写法
	*/
	age := int(11)
	switch {
	case age > 35:
		fmt.Println("中年")
	case age > 25:
		fmt.Println("青年")
	case age > 18:
		fmt.Println("少年")
	case age > 0:
		fmt.Println("童年")
	default:
		fmt.Println("unkonw")
	}

	fmt.Println("hello, loop")

}
