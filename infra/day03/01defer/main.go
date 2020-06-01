package main

import "fmt"

type person struct {
	name   string
	age    uint32
	gender bool
}

/**
Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，
也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
*/
func deferTest1() {
	defer fmt.Println("start")
	defer fmt.Println("0x1")
	defer fmt.Println("0x11")
	defer fmt.Println("0x111")
	defer fmt.Println("0x1111")
	defer fmt.Println("0x11111")
	defer fmt.Println("0x111111")
	defer fmt.Println("end")
}

/**
defer可以嵌套使用
*/
func deferTest2() {
	fmt.Println("start0")
	defer func() {
		fmt.Println("start1")
		defer func() {
			fmt.Println("start2")
			fmt.Println("end2")
		}()
		fmt.Println("end1")
	}()
	fmt.Println("end0")
}

func deferdemo1() int {
	var retcode = int(1)
	defer func() {
		retcode++
	}()
	return retcode
}

func deferdemo2() (retcode int) {
	retcode = 1
	defer func() {
		retcode++
	}()
	return retcode
}

func deferdemo3() (retcode int) {
	defer func() {
		retcode++
	}()
	return 2
}
func deferdemo4() (retperson person) {
	defer func() {
		retperson.name = "loop"
	}()
	retperson = person{"tu", 25, true}
	return
}

/**
记录被调用函数返回值的变量, 是被调用函数的局部变量,不是外部的调用函数中用来接收返回值的变量
*/
func variableRetruned1() (retcode int) {
	retcode = int(0)

	fmt.Printf("&retcode:%p\n", &retcode)
	return
}

func variableRetruned2() (retperson person) {
	retperson = person{"loop", 25, true}
	fmt.Printf("%+v, %p\n", retperson, &retperson)
	return
}

/**
推测 go语言函数返回流程:
go 语言无论是否对返回值命名, 都存在一个用作记录返回值的局部变量$(ret)
deferdemo1:
return retcode ===> $(ret) = retcode
defer func(){
	retcode++    ===> retcdoe++
}()

//真正返回代码
ret $(ret){
...
}

deferdemo2://对返回值命名 retcode 就代表局部变量$(ret)
retcode = 1  ===> $(ret) = 1
return retcode    ===> $(ret) = $(ret)
defer func() {
	retcode++    ===> $(ret) ++
}()

//真正返回代码
ret $(ret){
...
}

deferdemo3://对返回值命名 retcode 就代表变量$(ret)
return 2 ===> $(ret) = 2
defer func() {
	retcode++    ===> $(ret) ++
}()

//真正返回代码
ret $(ret){
...
}
*/
func main() {
	fmt.Println("hello, loop")

	fmt.Println(deferdemo1()) //1
	fmt.Println(deferdemo2()) //2
	fmt.Println(deferdemo3()) //3
	fmt.Println(deferdemo4())

	var retcode = variableRetruned1()
	fmt.Printf("%#v, %p\n", retcode, &retcode)

	var retperson = variableRetruned2()
	fmt.Printf("%+v, %p\n", retperson, &retperson)
}
