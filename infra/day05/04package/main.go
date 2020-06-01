package main

import "fmt"
import lother "github.com/loop-xxx/go-note/infra/day05/04package/other" //lother起别名

func main() {
	fmt.Println("hello, loop")

	//var p 接收一个unexport类型的返回值
	var p = lother.NewPerson("loop")
	fmt.Printf("%+v, %s\n", p, p.Name) //在other package 中unexport类型的export字段可以正常使用
	p.Hello()                          //在other package 中unexport类型的export方法可以正常使用

	var ep = lother.Person{} //不能初始化和使用Person.name字段
	fmt.Printf("%+v\n", ep)
}
