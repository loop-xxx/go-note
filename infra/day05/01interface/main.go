package main

import "fmt"

type speaker interface {
	speak()
}
type person struct {
}

type dog struct {
}

type cat struct {
}

type luna struct {
	person
}

func (p person) speak() {
	fmt.Println("aaa~")
}
func (d dog) speak() {
	fmt.Println("wangwangwang~")
}

func (c cat) speak() {
	fmt.Println("miaomiaomiao~")
}

func attack(target speaker) {
	target.speak()
}

func main() {
	fmt.Println("hello, world")
	var d1 = dog{}
	attack(d1)
	attack(cat{})
	attack(person{})

	var ss speaker

	fmt.Printf("%T, %v, %p\n", ss, ss, &ss)
	ss = d1
	fmt.Printf("%T, %v,%p\n", ss, ss, &ss)

	var s1 = speaker(d1)
	fmt.Printf("%T, %v\n", s1, s1)

	//"继承"了person类
	attack(luna{})
}
