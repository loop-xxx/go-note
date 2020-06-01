/*
//只有包下的main模块的main函数, 才会作为可执行程序的入口函数进行编译
package mobule

import "fmt"

func main(){

	fmt.Println("hello, main")
}
*/

/*
//function main is undeclared in the main package
package main

import "fmt"

func start(){
	fmt.Println("hello, main")
}
*/

package main

import "fmt"

func main(){
	fmt.Println("hello, main")
}