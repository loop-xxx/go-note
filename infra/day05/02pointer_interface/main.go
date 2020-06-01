package main

import "fmt"

type mover interface {
	move()
}
type eater interface {
	eat(food string)
}

type animal interface {
	mover
	eater
}

type person struct {
	name string
}

/*
func (p person) move() {
	fmt.Printf("%s is runing\n", p.name)
}

func (p person) eat(food string) {
	fmt.Printf("%s eat %s\n", p.name, food)
}
*/
func (pp *person) move() {
	fmt.Printf("%s is runing\n", pp.name)
}
func (pp *person) eat(food string) {
	fmt.Printf("%s eat %s\n", pp.name, food)
}

func main() {
	fmt.Println("hello, world")
	var p = person{"loop"}
	var a animal

	/**
	//当结构体方法的接收者为值类型, 那么这个结构体类型和结构体指针类型都实现了指针
	a = p
	fmt.Printf("%T, %+v\n", a, a)
	a = &p
	fmt.Printf("%T, %+v\n", a, a)
	*/

	//当结构体方法接受者为指针类型, 那么只有合资格结构体指针类型实现了指针
	a = &p
	fmt.Printf("%T, %+v\n", a, a)

}
