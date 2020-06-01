package main

import (
	"fmt"
	"sync"
)

var number int

func addsub(mutexp *sync.Mutex) {
	for i := int(0); i < 0x500000; i++ {
		mutexp.Lock()
		number++
		mutexp.Unlock()
	}
}
func add(wgp *sync.WaitGroup, mutexp *sync.Mutex) {
	defer wgp.Done()
	addsub(mutexp)
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(1)
	go add(&wg, &mutex)
	addsub(&mutex)
	wg.Wait()
	fmt.Printf("%#x\n", number)
}
