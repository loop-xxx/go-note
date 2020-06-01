package main

import "fmt"

//定义类型
type lint int

//类型给类型起别名
type oint = int

type person struct {
	name   string
	age    uint32
	gender bool
}

/**
func (l lint) String() (text string) {
	text = "hello, loop"
	return
}
*/
func (l *lint) String() (text string) {
	text = "hello, loop"
	return
}
func stringMethodTest() {
	var variable lint
	fmt.Println(variable)
	fmt.Printf("%#v, %v\n", variable, variable)
	fmt.Println(&variable)
	fmt.Printf("%#v, %v\n", &variable, &variable)
}

func typeTest() {
	var char rune
	char = 'l'
	fmt.Printf("%T, %c\n", char, char)
	var variable1 lint
	var variable2 oint
	fmt.Printf("%T, %T\n", variable1, variable2)

}

func structTest() {
	var obj1 struct {
		name string
	}
	var obj2 = struct {
		name string
	}{"loop"}
	fmt.Printf("%T, %+v\n", obj1, obj2)

	var obj3 person
	var obj4 = person{"loop", 25, true}
	fmt.Printf("%T, %+v\n", obj3, obj4)

	var obj5 = person{gender: true, age: 25}
	fmt.Printf("%#v", obj5)
}

func main() {
	fmt.Println("hello, world")
	//typeTest()
	//structTest()
	stringMethodTest()
}
