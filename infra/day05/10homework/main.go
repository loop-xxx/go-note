package main

import (
	"fmt"
	"io"
	"os"
)

func copyfile1(src string, dest string) (errno int) {
	//目标文件已存在
	errno = -1
	if _, err := os.Stat(dest); err != nil {
		//input文件打开失败
		errno = -2
		if infp, err := os.Open(src); err == nil {
			defer infp.Close()

			//output文件打开失败
			errno = -3
			if outfp, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0662); err == nil {
				defer outfp.Close()

				errno = 0
				var buf [0x1000]byte
				for n, err := infp.Read(buf[:]); err != io.EOF; n, err = infp.Read(buf[:]) {
					outfp.Write(buf[:n])
				}
			}
		}
	}

	return
}

func copyfile3(src string, dest string) (errno int) {
	//目标文件已存在
	errno = -1
	if _, err := os.Stat(dest); err != nil {
		//input文件打开失败
		errno = -2
		if infp, err := os.Open(src); err == nil {

			//output文件打开失败
			errno = -3
			if outfp, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0662); err == nil {

				errno = 0
				var buf [0x1000]byte
				for n, err := infp.Read(buf[:]); err != io.EOF; n, err = infp.Read(buf[:]) {
					outfp.Write(buf[:n])
				}

				outfp.Close()
			}

			infp.Close()
		}
	}

	return
}

func copyfile2(src string, dest string) int {
	if _, err := os.Stat(dest); err == nil {
		return -1 //文件存在
	}

	infp, err := os.Open(src)
	if err != nil {
		return -2 //打开src文件错误
	}
	defer infp.Close()

	outfp, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0622)
	if err != nil {
		return -3 //打开dest文件错误
	}
	defer outfp.Close()

	var buf [0x1000]byte
	for n, err := infp.Read(buf[:]); err != io.EOF; n, err = infp.Read(buf[:]) {
		fmt.Println(n)
		outfp.Write(buf[:n])
	}

	return 0
}

func copyfile4(src string, dest string) int {
	if _, err := os.Stat(dest); err == nil {
		return -1 //文件存在
	}

	infp, err := os.Open(src)
	if err != nil {
		return -2 //打开src文件错误
	}

	outfp, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0622)
	if err != nil {
		infp.Close()
		return -3 //打开dest文件错误
	}

	var buf [0x1000]byte
	for n, err := infp.Read(buf[:]); err != io.EOF; n, err = infp.Read(buf[:]) {
		outfp.Write(buf[:n])
	}

	infp.Close()
	outfp.Close()

	return 0
}

func main() {
	fmt.Println("hello, loop")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(copyfile1("./main.go", "./main1.txt"))
	fmt.Println(copyfile3("./main.go", "./main2.txt"))
}
