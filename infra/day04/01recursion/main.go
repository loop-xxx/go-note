package main

import "fmt"

/**
有n阶台阶, 一次可以走两次一次可以走2次, 可以有几种走法
*/
func recursion01(n uint32) (sum uint32) {
	/**
	if n == 1 { //最终剩余台阶数量(1/2)
		return 1
	} else if n == 2 {
		return 2
	}
	*/
	if n == 0 { //最终台阶剩余数量(0/1)
		sum = 1
	} else if n == 1 {
		sum = recursion01(n - 1)
	} else {
		sum = recursion01(n-1) + recursion01(n-2)
	}

	return
}

/**
有n阶台阶, 一次可以走(1->m)次, 可以有几种走法
*/
func recursion02(n uint32, m uint32) (sum uint32) {

	var loop uint32

	if n == 0 {
		sum = 1
	} else {
		if n >= m {
			loop = m
		} else {
			loop = n
		}
		for i := uint32(1); i <= loop; i++ {
			//fmt.Printf("recursion02(%d, %d)\n", n-i, m)
			sum += recursion02(n-i, m)
		}
	}
	return
}

func min(a uint32, b uint32) (ret uint32) {
	ret = a
	if b < a {
		ret = b
	}
	return
}

/**
有n阶台阶, 一次可以走(m1, m2)次, 可以有几种走法
*/
func recursion03(n uint32, m1 uint32, m2 uint32) (sum uint32) {
	if n < min(m1, m2) {
		sum = 1
	} else {
		if n >= m1 {
			sum += recursion03(n-m1, m1, m2)
		}
		if n >= m2 {
			sum += recursion03(n-m2, m1, m2)
		}
	}
	return
}

func main() {
	fmt.Println("hello, loop")
	fmt.Println(recursion01(6))
	fmt.Println(recursion02(6, 2))
	fmt.Println(recursion03(6, 1, 2))
}
