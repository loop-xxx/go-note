package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func chr(ch chan int) {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	/**
	number := <-ch
	fmt.Println(number)
	*/

	/**
	for number, ok := <-ch; ok; number, ok = <-ch {
		fmt.Println(number)
	}
	*/
	/**
	for range 循环等价于上述写法, 循环从channel中读取数据, 如果channel中的数据被读完, 则等待数据, 所以会死锁
	当只有管道关闭,且管道中的数据被读取完:
	(1)for range 循环才会退出
	(1) number, ok:=<-ch, ok才会 == false
	*/
	for number := range ch {
		fmt.Println(number)
		//close(ch)
	}

}

func chw(ch chan int) {
	defer wg.Done()
	ch <- 8
}

func main() {
	var ch chan int
	//(1)channel必须初始化才能使用,
	//(2)channel的写入和读取是原子操作
	/*(3)channel维护一块可以保存对应类型数据的内存$(cache), 但当有goroutinue写入管道时并不会直接将数据写入$(cache)而是先被阻塞,
	当有其他goroutine读取管道时(允许写标志位被置true), 数据此时写入到$(cache)(允许写标志位被置false,允许读标志为被置true),
	管道从$(cache)中读取, 完成后(允许读标志位被置false)
	*/
	/*(4)当channel申请了固定大小缓存之后, channel底层会额外维护一个指定大小环形队列, 往管道写入数据时, 会先缓存到环形队列里,
	环形队列满后阻塞, 当(允许写标志位被置true)时, 从队列中取一个数据, 写入到$(cache)(允许写标志位被置false,允许读标志为被置true),
	管道从$(cache)中读取, 完成后(允许读标志位被置false)*/

	//ch = make(chan int)
	ch = make(chan int, 0x1)

	wg.Add(2)
	go chr(ch)
	go chw(ch)

	ch <- 8
	time.Sleep(time.Second)

	close(ch)
	//向已关闭的管道写入数据则会触发panic
	//ch <- 19
	//var ch2 chan int
	//close(ch2) //关闭nil channel会触发panic
	wg.Wait()
}
