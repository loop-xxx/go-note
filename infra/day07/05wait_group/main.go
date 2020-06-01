package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

//rand包使用
func randtest() {
	rand.Seed(time.Now().UnixNano())
	for i := int(0); i < 0x10; i++ {
		//rand.Int()
		//rand.Int63n(range)
		fmt.Println(rand.Uint64())
	}
}

func f(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(0x400)))
	fmt.Println(i)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f(i)
	}
	wg.Wait()
}
