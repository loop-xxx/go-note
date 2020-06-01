package main

import (
	"fmt"
	"runtime"
	"sync"
)

func suba(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 0x10; i++ {
		fmt.Printf("a:%d\n", i)
	}
}

func subb(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 0x10; i++ {
		fmt.Printf("b:%d\n", i)
	}
}

// G goroutinue
// P (PROCS)管理着一组goroutine队列, 用于调度goroutinue
// M (machine)是Go运行时(runtime)对操作系统内核线程的虚拟, M与内核线程是一一映射的关系, groutine最终是要放到M上执行的
// P与M一般也是一一对应的. P管理着一组G挂载在M上运行.
// 设置GOMAXPROCS, 相当于间接设置开启操作系统内核线程的个数

func main() {
	var wg sync.WaitGroup
	fmt.Println(runtime.NumCPU()) //机器上的CPU核心数(intel 核心支持的超线程技术也算一个核心)
	runtime.GOMAXPROCS(1)         //设置GOMAXPROCS,相当于间接设置开启内核线程的个数 //默认值是机器上的CPU核心数
	wg.Add(2)
	go suba(&wg)
	go subb(&wg)

	wg.Wait()
}
