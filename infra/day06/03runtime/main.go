package main

import (
	"fmt"
	"path"
	"runtime"
)

func test02() {
	if pc, file, line, ok := runtime.Caller(0); ok {
		//fmt.Printf("%#x\n", pc)
		function := runtime.FuncForPC(pc)
		fmt.Println(function.Name())
		//fmt.Println(file)
		fmt.Println(path.Base(file))
		fmt.Println(line)
	}

}
func test1() {
	test02()
}

func main() {
	test1()
}
