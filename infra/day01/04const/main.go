package main

import "fmt"

const pi float32 = 3.1415926

const (
	a1 int   = 100
	a2 uint8 = 50
	a3       = 10
)

/*
	const 多常量声明,如果本行(非第一行)没有明确指定值, 则本行变量的值与上一行变量相同
*/
const (
	b1 float32 = 100
	b2
	b3
)

/*
	iota是go语言的常量计数器，只能在常量的表达式中使用。

	iota在const代码块第一行常量声明语句中被重置为0(无论第一行常量声明是否用到iota), const中每新增一行常量声明将使iota计数一次
	iota可理解为const语句块中的行索引,使用iota能简化定义，在定义枚举时很有用。
*/
const (
	h1 = 100
	h2 = iota // iota = 1
	h3        // iota = 2
)

const (
	c1 float32 = iota
	c2
	c3
)

const (
	d1 = iota
	_
	d2
)

const (
	e1, e2 = iota + 1, iota + 2
	e3, e4 = iota + 3, iota + 4
)

const (
	f1, f2 = iota + 1, iota + 2
	_, _
	f3, f4 = iota + 3, iota + 4
)

const (
	g1 = iota
	g2 = 100
	g3
	g4 = iota
)

const (
	_  uint64 = 1 << (10 * iota)
	kb        //uint64 = 1 << (10 * iota)
	mb        //uint64 = 1 << (10 * iota)
	gb        //uint64 = 1 << (10 * iota)
	tb        //uint64 = 1 << (10 * iota)
	pb        //uint64 = 1 << (10 * iota)
)

func main() {

	fmt.Println("hello, loop")
	fmt.Println(pi)
	fmt.Println(a1, a2, a3)
	fmt.Println(b1, b2, b3)
	fmt.Println(c1, c2, c3)
	fmt.Println(d1, d2)
	fmt.Println(e1, e2, e3, e4)
	fmt.Println(f1, f2, f3, f4)
	fmt.Println(g1, g2, g3, g4)
	fmt.Println(h1, h2, h3)
	fmt.Printf("%#x, %#x, %#x, %#x, %#x", kb, mb, gb, tb, pb)
}
