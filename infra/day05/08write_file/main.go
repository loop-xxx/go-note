package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writetext(filename string) (errno int) {
	errno = -1
	//O_WRONLY|O_CREATE|O_TURNC 文件不存在创建新文件, 创建权限0664, 存在则truncate
	//O_WRONLY|O_CREATE|O_APPEND 文件不存在创建新文件 创建权限0664, 存在则append
	if fp, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664); err == nil {
		errno = 0 //程序内部可以处理的错误,交给上级处理函数处理
		defer fp.Close()

		if _, err := fp.Write([]byte("hello, world\n")); err != nil {
			panic(err)
		}
		if _, err := fp.WriteString("hello, loop\n"); err != nil {
			panic(err)
		}
	}

	return
}

func writetext2(filename string) (errno int) {
	errno = -1
	//O_WRONLY|O_CREATE|O_TURNC 文件不存在创建新文件, 创建权限0664, 存在则truncate
	//O_WRONLY|O_CREATE|O_APPEND 文件不存在创建新文件 创建权限0664, 存在则append
	if fp, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664); err == nil {
		errno = 0 //程序内部可以处理的错误,交给上级处理函数处理
		defer fp.Close()

		writer := bufio.NewWriter(fp)
		//写入到缓冲区, 缓冲区满再flush到缓冲区
		for i := int(0); i < 0x10; i++ {
			_, err := writer.WriteString("hello, world2\n")
			if err != nil {
				panic(err)
			}
		}
		//刷新缓冲区, 将缓冲区数据写入到文件
		writer.Flush()
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

	if errno := writetext("./loop.txt"); errno != 0 {
		//log.wang()
		if errno = writetext("C:/Users/Tempest/Desktop/loop.txt"); errno != 0 {
			panic("writetext error")
		}
	}

	if errno := writetext2("./loop.txt"); errno != 0 {
		//log.wang()
		if errno = writetext2("C:/Users/Tempest/Desktop/loop.txt"); errno != 0 {
			panic("writetext2 error")
		}
	}

	//默认以O_WRONLY|O_CREATE|O_TURNC方式打开文件
	if err := ioutil.WriteFile("./loop.txt", []byte("hello, world"), 0664); err != nil {
		if err := ioutil.WriteFile("C:/Users/Tempest/Desktop/loop.txt", []byte("hello, world"), 0664); err != nil {
			panic(err)
		}
	}
}
