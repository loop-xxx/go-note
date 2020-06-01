package main

import "fmt"

func main() {
	fmt.Println("hello, world")

	//slice类型变量分配内存的四种方式

	var slice1 = make([]int, 0, 0x10)
	fmt.Println(len(slice1), cap(slice1), slice1)

	var arr [4]int
	var slice2 = arr[:1]
	fmt.Println(len(slice2), cap(slice2), slice2)

	var slice3 = []int{1, 2, 3, 4}
	fmt.Println(len(slice3), cap(slice3), slice3)

	var slice4 = slice3[:2]
	fmt.Println(len(slice4), cap(slice4), slice4)

	var map1 = make(map[int]string, 0x10)
	fmt.Println(len(map1), map1)

	//map类型的数组和切片
	var mapArr [2]map[int]string
	for _, v := range mapArr {
		fmt.Printf("%#v ", v)
	}
	fmt.Println()
	var mapSlice = mapArr[:]
	fmt.Println(mapSlice)

	//value值为数组或切片的map
	var datMap1 map[int][4]int
	datMap1 = make(map[int][4]int, 0x10)
	datMap1[0] = [4]int{1, 2, 3, 4}
	fmt.Println(datMap1)
	var datMap2 map[int][]int
	datMap2 = make(map[int][]int, 0x10)
	v, isok := datMap1[0]
	if isok {
		datMap2[0] = v[:]
	}

	fmt.Println(datMap2)
}
