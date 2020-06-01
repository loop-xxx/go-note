package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  uint
}
type unsignedint = uint32
type unsignedlongint uint64

func reflectType(parameter interface{}) {
	t := reflect.TypeOf(parameter)
	fmt.Printf("%v\n", t)
	fmt.Printf("%#v, %v\n", t.Name(), t.Name())
	fmt.Printf("%#v, %v\n", t.Kind(), t.Kind())
}

func reflectValue(parameter interface{}) {
	v := reflect.ValueOf(parameter)
	fmt.Printf("%#v, %v\n", v, v)
	fmt.Printf("%#v, %v\n", v.Kind(), v.Kind())
	if v.Kind() == reflect.Uint {
		fmt.Println(v.Uint())
	}
}

func reflectSetValue1(parameter interface{}) {
	v := reflect.ValueOf(parameter)
	fmt.Printf("%#v, %v\n", v, v)
	fmt.Printf("%#v, %v\n", v.Kind(), v.Kind())
	if v.Kind() == reflect.Uint {
		//v.SetInt(0x10) //修改的是副本(interface{}的value域), reflect包会引发panic
	}
}
func reflectSetValue2(parameter interface{}) {
	//pt := reflect.TypeOf(parameter)
	//pt.Elem() //reflect.Type 返回指针指向的类型 (指针是什么类型的指针)
	pv := reflect.ValueOf(parameter)
	if pv.Kind() == reflect.Ptr {
		v := pv.Elem() //refelect.Value 返回指针指向的值
		fmt.Printf("%#v, %v\n", v, v)
		fmt.Printf("%#v, %v\n", v.Kind(), v.Kind())
		if v.Kind() == reflect.Uint {
			v.SetUint(0x10)
		}
	}
}

func isNil0isValid() {
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}

func main() {
	//var variable uint
	fmt.Println("hello, loop")
	//reflectType(uint64(0))
	reflectType(person{"loop", 25})
	//reflectType(&person{"loop", 25})
	//reflectType(rune('a'))
	//reflectType(unsignedint(0x0))
	//reflectType(unsignedlongint(0x0))
	//reflectType(&variable)

	//reflectValue(person{"loop", 25})
	//reflectValue(&variable)
	//reflectValue(variable)

	/**
	fmt.Printf("-----%d-----\n", variable)
	reflectSetValue1(variable)
	fmt.Printf("-----%d-----\n", variable)
	reflectSetValue2(&variable)
	fmt.Printf("-----%d-----\n", variable)
	*/
	//isNil0isValid()
}
