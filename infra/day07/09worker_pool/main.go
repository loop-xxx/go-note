package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type id = uint64
type person struct {
	id   uint64
	name string
}

func id2str(id uint64) (str string) {
	for id > 0 {
		str += fmt.Sprintf("%d", id%10)
		id /= 10
	}
	return
}

func workersub(idchan <-chan id, pechan chan<- person) {
	var pe person
	for id := range idchan {
		pe.id = id
		pe.name = id2str(id)
		pechan <- pe
	}
}

func worker(wgp *sync.WaitGroup, idchan <-chan id, pechan chan<- person) {
	defer wgp.Done()
	workersub(idchan, pechan)
}

//单项通道, 限制只读(<-chan)或只写(chan<-), 防止代码中误操作
func workerpool(extra int, idchan <-chan id, pechan chan<- person) {
	var wwg sync.WaitGroup
	wwg.Add(extra)
	for i := int(0); i < extra; i++ {
		go worker(&wwg, idchan, pechan)
	}
	//workerpool goroutine 除了负责开启其他worker,本身也是个worker
	workersub(idchan, pechan)
	//等待所有worker执行结束后关闭输出管道
	wwg.Wait()
	close(pechan)
}

func handlersub(idchan chan<- id) {
	for i := 0; i < 0x100; i++ {
		idchan <- rand.Uint64()
	}
}
func handler(wgp *sync.WaitGroup, idchan chan<- id) {
	defer wgp.Done()
	handlersub(idchan)
}

//单项通道, 限制只读(<-chan)或只写(chan<-), 防止代码中误操作
func handlerpool(extra int, idchan chan<- id) {
	var hwg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	hwg.Add(extra)
	for i := int(0); i < extra; i++ {
		go handler(&hwg, idchan)
	}

	//handlerpool goroutine 除了负责开启其他handler,本身也是个handler
	handlersub(idchan)

	//等待所有handler执行结束后关闭输出管道
	hwg.Wait()
	close(idchan)
}

func receiversub(pechan <-chan person) {
	for pe := range pechan {
		fmt.Printf("%d, %#v\n", pe.id, pe.name)
	}
}
func receiver(wgp *sync.WaitGroup, pechan <-chan person) {
	defer wgp.Done()
	receiversub(pechan)
}

func receiverpool(extra int, pechan <-chan person) {
	var rwg sync.WaitGroup
	rwg.Add(extra)
	for i := int(0); i < extra; i++ {
		go receiver(&rwg, pechan)
	}
	receiversub(pechan)
	rwg.Wait()
}

/**
发布任务流程groutine(组), 执行任务流程groutine(组), 取结果流程groutine(组), 三个groutine(组)是并行
但是结束是有顺序的, 而且是一定的
(1)发布任务groutine(组), 发布完所有任务, 关闭发布任务管道, 发布任务流程结束
(2)执行任务groutine(组), 在发布任务管道关闭后, 执行完管道中剩余的任务, 并将结果全部推送到结果管道, 关闭结果管道, 执行任务流程结束
(3)取结果流程groutine(组), 在结果管道关闭后, 取完管道中剩余的任务,并将结果取完, 取结果流程结束
*/
//此时main goroutine 是取结果流程groutine组的groutine, 最后结束, 所以不用wait handlerpool(发布任务groutine(组))和 workerpool(执行任务groutine(组))

func main() {
	var idchan chan id
	idchan = make(chan id)

	var pechan chan person
	pechan = make(chan person)

	//执行任务流程gorotinue组
	go workerpool(0x4, idchan, pechan)
	//发布任务流程gorotinue组
	go handlerpool(0x1, idchan)

	//接收结果流程 gorotinue组 (将main gorotinue作为接收结果流程组组长)
	receiverpool(0x0, pechan)
}
