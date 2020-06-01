package main

import (
	"fmt"
	"github.com/loop-xxx/go-note/infra/day06/09homework/str2objs"
)

type person struct {
	Name string
	Age  string
}

func main() {
	//var p = person{"luna", "28"}
	var p person
	fmt.Println(p)
	str2objs.Str2obj("person{Name:loop, Age:25}", &p)
	fmt.Println(p)
}
