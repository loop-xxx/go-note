package main

import "fmt"

func main() {
	var intchan1 chan int
	var intchan2 chan int
	intchan1 = make(chan int, 0x1)
	intchan2 = make(chan int, 0x1)
	intchan1 <- 0x10

	for i := int(0); i < 0x10; i++ {

		/**
		select 同时操作多个管道, 标签的判断顺序随机, 流程如下:
		进入select先随机一个标签进行判断:
		标签为读取管道:
			(1)可读, 进入case代码块执行结束后break, select执行结束
			(2)channel为空, 再随机下一个标签进行判断, 如果该标签为最后一个标签,
			若有default标签, 执行default代码块结束select, 否则select阻塞
			(3)channel为空且关闭, 会进入case代码块执行结束后break, select执行结束
		标签为写入管道
			(1)可写, 进入case代码块执行结束后break, select执行结束
			(2)channel已满, 再随机下一个标签进行判断, 如果该标签为最后一个标签,
			若有default标签, 执行default代码块结束select, 否则select阻塞
			(3)channel关闭, 直接panic
		*/

		//select {} //select 如果没有标签可以匹配时, 会阻塞

		select {
		case number, ok := <-intchan1:
			if ok {
				fmt.Println("intchan1", number, i)
			}
		case intchan1 <- i:
		case intchan2 <- i:
		case number, ok := <-intchan2:
			if ok {
				fmt.Println("intchan2", number, i)
			}

		}
		if i == 0 {
			close(intchan1)
			intchan1 = nil //使该管道无效, 所有操作该管道的标签都会阻塞, select不会再进入操作该管道的标签
		}
	}
}
