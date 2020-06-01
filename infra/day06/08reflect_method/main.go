package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Eat(number uint) {
	fmt.Println(number)
}

func (sp *student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	if t == v.Type() {
		fmt.Println("reflect.TypeOf(x) == reflect.ValueOf(x).Type()")
	}

	for i := 0; i < v.NumMethod(); i++ {
		fmt.Printf("method:%s\n", v.Method(i).Type())
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		if t.Method(i).Name != "Eat" {
			var args = []reflect.Value{}
			fmt.Printf("%v\n", v.Method(i).Call(args))
		} else {
			var parameter = uint(0xffffffff)

			var args = []reflect.Value{reflect.ValueOf(parameter)}
			v.Method(i).Call(args)
		}
	}

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", t.Method(i).Type)
	}
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}
	printMethod(stu1)
	fmt.Println("----------------------")
	printMethod(&stu1)
}
