package main

import "fmt"

func main() {
	fmt.Println("hello, loop")
	for i := int(1); i < 10; i++ {
		for j := int(1); j <= i; j++ {
			fmt.Printf("%d*%d=%2d\t", j, i, j*i)
		}
		fmt.Println()
	}

	/*
		for i := int(9); i > 0; i-- {
			for j := int(1); j <= i; j++ {
				fmt.Printf("%d*%d=%2d\t", j, i, j*i)
			}
			fmt.Println()
		}
	*/
	/*
		for i := int(1); i < 10; i++ {
			for j := 9; j >= i; j-- {
				fmt.Printf("%d*%d=%2d\t", j, i, j*i)
			}
			fmt.Println()
		}
	*/
}
