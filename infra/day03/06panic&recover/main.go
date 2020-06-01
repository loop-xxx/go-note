package main

import "fmt"

func f0() {
	fmt.Println("f0")
}
func f1(code int) {
	defer func() {
		fmt.Println("defer")
		/**
		var err = recover()
		if err != nil {
			fmt.Printf("%T, %#v\n", err, err)
		}
		*/
	}()
	if code < 0 {
		panic("code < 0")
	}
	fmt.Println("f1")
}
func f2() {
	fmt.Println("f2")
}

func proc1() {
	defer func() {
		var err = recover()
		if err != nil {
			fmt.Printf("%T, %#v\n", err, err)
		}
	}()
	f0()
	f1(-1)
	f2()
}

/**
defer函数中可以panic()
*/
func deferTest(code uint32) {
	defer func() {
		fmt.Println("0x1")
	}()
	defer func() {
		fmt.Println("0x11")
	}()
	defer func() {
		fmt.Println("0x1111")
		if code == 0 {
			panic("code == 0")
		}
		fmt.Println("0x111")
	}()
	defer func() {
		fmt.Println("0x11111")
	}()
	defer func() {
		fmt.Println("0x111111")
	}()
}
func proc2() {
	defer func() {
		var err = recover()
		if err != nil {
			fmt.Printf("%T, %#v\n", err, err)
		}
	}()
	deferTest(0)
}

func main() {
	fmt.Println("hello, loop")
	proc2()
	fmt.Println("bye loop")

}
