package other

import "fmt"

// Person type
type Person struct {
	name string
}
type person struct {
	Name string
}

// NewPerson 创建一个 person 变量, waring:返回值一个 person(unexport类型)类型的值
func NewPerson(name string) (p person) {
	p = person{name}
	return
}

func (p person) Hello() {
	fmt.Printf("%s:hello\n", p.Name)
}
