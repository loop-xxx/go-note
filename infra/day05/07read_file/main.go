package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
return:
	-1 open file error
	0 success
*/
func readtext(filename string) (errno int) {
	errno = -1
	if fp, err := os.Open(filename); err == nil {
		defer fp.Close()

		//打开文件失败, 程序内部可以处理的错误, 交给上级调用函数处理
		errno = 0

		var buf [0x1000]byte
		//read函数出错一般是系统或硬件出问题, 程序内部无法处理的错误建议直接panic
		for n, err := fp.Read(buf[:]); err != io.EOF; n, err = fp.Read(buf[:]) {
			fmt.Print(string(buf[:n]))
		}

	}
	return
}

func readtext2(filename string) (errno int) {
	errno = -1
	if fp, err := os.Open(filename); err == nil {
		defer fp.Close()
		//打开文件失败, 交给上级调用函数处理
		errno = 0

		reader := bufio.NewReader(fp)

		//read函数出错一般是系统或硬件出问题, 程序内部无法处理的错误建议直接panic
		for line, err := reader.ReadString('\n'); err != io.EOF; line, err = reader.ReadString('\n') {
			fmt.Print(line)
		}
	}

	return
}

func main() {
	fmt.Println("hello, world")

	defer func() { //最终处理打印err日志, 不让程序崩溃
		if err := recover(); err != nil {
			//log.error()
			fmt.Println(err)
		}
	}()

	/*
		if errno := readtext("./main.go"); errno == -1 {
			//打印log.wang()
			//程序内部处理error
			if errno := readtext("C:/Users/Tempest/go/src/github.com/loop/day05/07file/main.go"); errno != 0 {
				panic("readtext error")
			}
			panic("readtext error")
		}
	*/

	if errno := readtext2("./main.go"); errno == -1 {
		//打印log.wang()
		//程序内部处理error
		if errno := readtext2("C:/Users/Tempest/go/src/github.com/loop/day05/07file/main.go"); errno != 0 {
			panic("readtext2 error")
		}
	}

	/*
		content, err := ioutil.ReadFile("./main.go")
		if err != nil {
			content, err = ioutil.ReadFile("C:/Users/Tempest/go/src/github.com/loop/day05/07file/main.go")
			if err != nil {
				panic("ioutil.readFile error")
			}
		}
		fmt.Print(string(content))
	*/
}
