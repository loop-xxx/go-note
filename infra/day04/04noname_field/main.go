package main

import "fmt"

//匿名字段
type dog struct {
	string
	uint32
}

type address struct {
	province string
	city     string
}

//结构体类型的匿名字段
type person struct {
	name string
	age  uint32
	address
}

func main() {
	var d1 = dog{"white", 0x10}
	fmt.Println(d1)
	var d2 = dog{string: "white", uint32: 0x10}

	fmt.Printf("%s, %d; %s, %d\n", d1.string, d1.uint32, d2.string, d2.uint32)

	var p1 = person{"loop", 25, address{"河北", "邯郸"}}
	var p2 = person{
		name:    "loop",
		age:     25,
		address: address{"河南", "安阳"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	//结构体类型匿名字段语法糖:省略匿名结构体字段, 直接.匿名结构体字段的字段
	fmt.Printf("%s, %s; %s:%s", p1.address.province, p1.address.city, p2.province, p2.city)

}
