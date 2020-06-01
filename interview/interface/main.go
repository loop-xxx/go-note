package main

import (
	"fmt"
)

type UserName struct{
	firstName string
	secondName string
}

type UserInfo1 struct{
	name UserName
	age uint
}
type UserInfo2 struct{
	name *UserName
	age uint
}


func main(){
	u1 := UserName{"li", "loop"}
	u2 := UserName{"li", "loop"}
	i1 := interface{}(u1)
	i2 := interface{}(u2)

	fmt.Println(i1 == i2)

	ui11 := UserInfo1{
		name: u1,
		age:  0,
	}

	ui12 := UserInfo1{
		name: u2,
		age:  0,
	}
	fmt.Println(ui11 == ui12 )

	ui21 := UserInfo2{
		name: &u1,
		age:  0,
	}

	ui22 := UserInfo2{
		name: &u2,
		age:  0,
	}
	fmt.Println(ui21 == ui22 )

	str := "loop"
	str1 := "li loop"
	str2 := fmt.Sprintf("%s %s", "li", str)
	i3 := interface{}(str1)
	i4 := interface{}(str2)

	fmt.Println(i3 == i4)
}