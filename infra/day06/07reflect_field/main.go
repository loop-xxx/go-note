package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func printField(x interface{}) {

	t := reflect.TypeOf(x)
	fmt.Println(t.Kind())           // student struct
	if t.Kind() == reflect.Struct { //只有结构体类型可以使用以下方法, 结构体指针类型不可以
		// 通过for循环遍历结构体的所有字段信息
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
		}

		// 通过字段名获取指定结构体字段信息
		if scoreField, ok := t.FieldByName("Score"); ok {
			fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
		}
	}

	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())
	if v.Kind() == reflect.Struct {
		fmt.Println(v.NumField())
		scoreV := v.FieldByName("Score")
		fmt.Println(scoreV)
		fmt.Println(v.Field(0))

		//scoreV.SetInt(100) //修改的是副本(interface{}的value域), reflect包会引发panic
	}

}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}
	printField(stu1)
	fmt.Println("-----------------")
	//printField(&stu1)
}
