package main

func f1(str string) {
	defer func() {
		panic("f1")
	}()
}

func f2(s []int) {
	defer func() {
		panic("f2")
	}()
}

func f3(s []int) (str string) {
	str = "loop"
	defer func() {
		panic("f3")
	}()
	return
}

func f4(ch chan int) {
	defer func() {
		panic("f3")
	}()
}
func main() {
	f1("loop")
	var arr = [4]int{0x1, 0x2, 0x3, 0x4}
	//f2(arr[:])
	f3(arr[:])

	//var ch = make(chan int)
	//f4(ch)
}
