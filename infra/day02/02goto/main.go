package main

import "fmt"

func main() {
	fmt.Println("hello, world")

	//没有goto语句跳出多层循环
	var flag = bool(false)
	for i := uint32(0); i < 0x10; i++ {
		for j := uint32(0); j < 0x10; j++ {
			if i*0x10+j == 100 {
				flag = true
				break
			}
			fmt.Printf("%2d ", i<<4+j)
		}

		if flag {
			break
		}
		fmt.Println()
	}

	fmt.Println()

	//goto 跳出多层循环
	for i := uint32(0); i < 0x10; i++ {
		for j := uint32(0); j < 0x10; j++ {
			if i*0x10+j == 100 {
				goto over
			}
			fmt.Printf("%2d ", i<<4+j)
		}
		fmt.Println()
	}
over:
	fmt.Println()

	//break 跳出外层循环
outer1:
	for i := uint32(0); i < 0x10; i++ {
		for j := uint32(0); j < 0x10; j++ {
			if i*0x10+j == 100 {
				break outer1
			}
			fmt.Printf("%2d ", i<<4+j)
		}
		fmt.Println()
	}
	fmt.Println()

	//continue 外层循环, 不能达到目的
outer2:
	for i := uint32(0); i < 0x10; i++ {
		for j := uint32(0); j < 0x10; j++ {
			if i*0x10+j == 100 {
				continue outer2
			}
			fmt.Printf("%2d ", i<<4+j)
		}
		fmt.Println()
	}
	fmt.Println()
}
