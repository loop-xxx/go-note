package main

import "fmt"

type dog struct {
	name string
	age  uint64
}

func (d dog) call() {
	fmt.Printf("%s:www~\n", d.name)
}

//d不是调用该方法的对象, 仅仅只是growUp1()方法的一个参数, struct变量传参拿到的只是调用该方法的对象的副本
func (d dog) growUp1() {
	d.age++
}

//d为调用该方法的变量的指针, 有点类似其他面向对象语言中this指针
func (d *dog) growUp2() {
	(*d).age++
}

func newDog1(name string, age uint64) (retd dog) {
	retd = dog{name, age}
	return
}

func newDog2(name string, age uint64) (retpd *dog) {
	var d = dog{name, age}
	retpd = &d
	return
}

/**
retrun_value为常量, 不能取地址, 也不能修改或间接修改

//内置类型map, 无法获取map中储存的值所在的地址, &map[key]不被允许, map的[],[]= 运算符的行为类似于方法调用
func (m map[KTYPE]VTYPE)operator[](key KTYPE)(value VTYPE){
	...
}
func (m map[KTYPE]VTYPE)operator[]=(key KTYPE, value VTYPE){
	...
}

var str = string("hello")
fmt.Printf("%p\n", &str[0])
//内置类型string, &string[index]不被允许, string[index]不能修改,  string []运算符的行为类似方法调用
func (str string)operator[](index int)(byte b){
	...
}

//内置类型 array []运算符底层是通过指针偏移来索引数据 允许&array[index], 允许直接或间接修改
//内置类型 slice []运算符底层也是通过指针偏移来索引数据 允许&slice[index], 允许直接或间接修改
*/

func main() {
	fmt.Println("hello, loop")
	var dd = dog{"white", 0x10}
	dd.call()
	fmt.Println(dd)
	dd.growUp1()
	fmt.Println(dd)
	dd.growUp2()
	fmt.Println(dd)

	//返回结构体变量
	//newDog1("white", 0x10).name = "small white"
	newDog1("white", 0x11).growUp1()
	//newDog1("white", 0x11).growUp2()

	//返回结构体变量地址
	newDog2("black", 0x10).name = "small black"
	newDog2("black", 0x10).growUp1()
	newDog2("black", 0x10).growUp2()

	var dogMap = make(map[uint32]dog, 0x1)
	dogMap[0] = dd
	//内置数据类型map []运算符类似于map类型的方法
	fmt.Println(dogMap[0].name)
	dogMap[0].growUp1()
	//dogMap[0].growUp2()
	//dogMap[0].name = "black"

	//内置数据类型slice []运算符底层也是和array类型类型相同通过指针偏移
	var s = []dog{dd}
	fmt.Println(s)
	s[0].growUp1()
	s[0].growUp2()
	s[0].name = "black"
	fmt.Println(s)

	//函数有类型,方法不是函数, 没有类型
	//fmt.Printf("%T, %p\n", growUp1, growUp1)
	//fmt.Printf("%T, %p\n", growUp2, growUp2)
}
