package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello, world")
	var str string
	var n int
	fmt.Scanln(&str)
	fmt.Println(str)

	reader := bufio.NewReader(os.Stdin)
	//清空输入
	reader.ReadString('\n')

	fmt.Scanln(&n)
	fmt.Println(n)

	str, _ = reader.ReadString('\n')
	fmt.Println(str)

}
