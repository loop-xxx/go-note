package main

import "fmt"

//空接口类型可以储存任意类型
var gmap map[string]interface{}

//空接口类型可以接受任意类型的参数
func typeAssert(x interface{}) {
	//类型断言猜测类型
	v, is := x.(int)
	if is {
		fmt.Printf("%d", v)
	} else {
		fmt.Println("not int")
	}

	//类型断言 switch语法支持
	switch x.(type) {
	case int:
		fmt.Println("is int")
	case string:
		fmt.Println("is string")
	case bool:
		fmt.Println("is bool")
	}
	switch val := x.(type) {
	case int:
		fmt.Println("is int", val)
	case string:
		fmt.Println("is string", val)
	case bool:
		fmt.Println("is bool", val)
	}

}
func main() {
	fmt.Println("hello, loop")
	gmap = make(map[string]interface{}, 0x10)
	gmap["name"] = "loop"
	gmap["age"] = 25
	gmap["gender"] = true
	fmt.Printf("%T, %#v\n", gmap, gmap)
	typeAssert("aaa")
}
