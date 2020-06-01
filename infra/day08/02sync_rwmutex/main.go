package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type person struct {
	fristname  string
	secondname string
}

var global person

func writer(wgp *sync.WaitGroup, mutexp *sync.Mutex) {
	defer wgp.Done()
	number := rand.Intn(1000)
	mutexp.Lock()
	global.fristname = strconv.Itoa(number)
	time.Sleep(time.Second)
	global.secondname = strconv.Itoa(number + 1)
	mutexp.Unlock()
}

func reader(wgp *sync.WaitGroup, mutexp *sync.Mutex) {
	defer wgp.Done()
	mutexp.Lock()
	time.Sleep(time.Millisecond)
	fmt.Printf("%s %s\n", global.fristname, global.secondname)
	mutexp.Unlock()
}

func writer2(wgp *sync.WaitGroup, rwmp *sync.RWMutex) {
	defer wgp.Done()
	number := rand.Intn(1000)
	rwmp.Lock()
	global.fristname = strconv.Itoa(number)
	time.Sleep(time.Second)
	global.secondname = strconv.Itoa(number + 1)
	rwmp.Unlock()
}

func reader2(wgp *sync.WaitGroup, rwmp *sync.RWMutex) {
	defer wgp.Done()
	rwmp.RLock()
	time.Sleep(time.Millisecond)
	fmt.Printf("%s %s\n", global.fristname, global.secondname)
	rwmp.RUnlock()
}

func test() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := int(0); i < 0x4; i++ {
		wg.Add(1)
		go writer(&wg, &mutex)
	}

	for i := int(0); i < 0x100; i++ {
		wg.Add(1)
		go reader(&wg, &mutex)
	}
	wg.Wait()
}

func test2() {
	var wg sync.WaitGroup
	var rwm sync.RWMutex
	rand.Seed(time.Now().UnixNano())
	for i := int(0); i < 0x4; i++ {
		wg.Add(1)
		go writer2(&wg, &rwm)
	}
	for i := int(0); i < 0x100; i++ {
		wg.Add(1)
		go reader2(&wg, &rwm)
	}
	wg.Wait()
}

func main() {
	start := time.Now()
	rand.Seed(time.Now().UnixNano())
	test2()
	fmt.Println(time.Now().Sub(start))

}
