package main

import (
	"fmt"
	"sync"
)

func handler(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := 0; i < 0x100; i++ {
		ch <- i
	}
}

func worker(wg *sync.WaitGroup, ch1, ch2 chan int) {
	defer wg.Done()
	for number := range ch1 {
		ch2 <- number * 2
	}
}

//发布任务 goroutine组
func handlerpool(ch chan int) {
	var wwg sync.WaitGroup
	for i := 0; i < 0x2; i++ {
		wwg.Add(1)
		go handler(&wwg, ch)
	}
	wwg.Wait()
	close(ch)
}

//执行任务 goroutine组
func workpool(ch1, ch2 chan int) {
	var rwg sync.WaitGroup
	for i := 0; i < 0x4; i++ {
		rwg.Add(1)
		go worker(&rwg, ch1, ch2)
	}
	rwg.Wait()
	close(ch2)
}

//goroutine泄露
func blockproc() {
	var chint chan int //未初始化管道
	//chint <- 0x10      //未初始化管道, 读取和写入都会阻塞
	<-chint //未初始化管道, 读取和写入都会阻塞
	fmt.Println("blockproc")
}

/**
发布任务流程groutine(组), 执行任务流程groutine(组), 取结果流程groutine(组), 三个groutine(组)是并行
但是结束是有顺序的, 而且是一定的
(1)发布任务groutine(组), 发布完所有任务, 关闭发布任务管道, 发布任务流程结束
(2)执行任务groutine(组), 在发布任务管道关闭后, 执行完管道中剩余的任务, 并将结果全部推送到结果管道, 关闭结果管道, 执行任务流程结束
(3)取结果流程groutine(组), 在结果管道关闭后, 取完管道中剩余的任务,并将结果取完, 取结果流程结束
*/
//此时main goroutine 是取结果流程的groutine, 最后结束, 所以不用wait handlerpool(发布任务groutine(组)) 和 workpool(执行任务groutine(组))
func main() {
	var ch1 = make(chan int, 0x10)
	var ch2 = make(chan int, 0x10)
	go blockproc() //goroutine泄露

	go workpool(ch1, ch2)
	go handlerpool(ch1)

	//取结结果goroutine
	for number := range ch2 {
		fmt.Print(number, " ")
	}
	fmt.Println()
}
