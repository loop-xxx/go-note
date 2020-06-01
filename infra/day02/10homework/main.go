package main

import (
	"fmt"
	"strings"
	"unicode"
)

func calcWrodCount(text string) {
	var words = strings.Split(text, " ")
	var wordMap = make(map[string]int, len(words))

	for _, word := range words {
		if value, isok := wordMap[word]; isok {
			wordMap[word] = value + 1
		} else {
			wordMap[word] = 1
		}
	}
	fmt.Println(len(wordMap), wordMap)
}

func calcHanCount(text string) {
	var total uint32
	for _, val := range text {
		if unicode.Is(unicode.Han, val) {
			total++
		}
	}
	fmt.Printf("hanCount:%#v\n", total)
}

func isHui(text string) bool {
	var isHui bool //默认初始化为false

	chars := strings.Split(text, "")
	//fmt.Println(chars)
	if end := len(chars); end > 0 {
		start := int(0)
		end -= int(1)

		for start < end {
			if chars[start] != chars[end] {
				goto ret
			}
			start++
			end--
		}
		isHui = true
	ret:
	}
	return isHui

}
func isHui2(text string) bool {
	var isHui bool //默认初始化为false

	/**
	//错误的比较string类型的[]操作符返回的是byte,如果字符串全部是ascii码字符则没有问题
	if end := len(text); end > 0 {
		end -= int(1)
		for i := range text {
			if text[i] != text[end-i] {
				fmt.Println(i)
				goto over
			}
		}
	over:
	}
	*/
	var runeSlice = []rune(text)
	if end := len(runeSlice); end > 0 {
		end -= int(1)
		for i := range runeSlice {
			if runeSlice[i] != runeSlice[end-i] {
				fmt.Println("----->", i, "<-----")
				goto over
			}
		}
		isHui = true
	over:
	}

	return isHui

}
func main() {
	fmt.Println("hello, loop")
	calcWrodCount("how do you do")
	calcHanCount("你好世界 hello, world")

	var str1 = string("hello, 你好世界")
	fmt.Println(str1, isHui2(str1))

	str1 = "黄山落叶松叶落山黄"
	fmt.Println(str1, isHui2(str1))
	str1 = "山西煤运车运煤西山"
	fmt.Println(str1, isHui2(str1))

}
