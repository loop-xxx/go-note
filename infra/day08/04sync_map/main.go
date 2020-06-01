package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var data sync.Map

func store(key string, value int) {

	data.Store(key, value)
}
func load(key string) {
	if value, ok := data.Load(key); ok {

		fmt.Println(value)
	}
}

var data2 = make(map[string]int, 0x20)

func store2(key string, value int) {
	data2[key] = value
}

func load2(key string) {
	fmt.Println(data2[key])
}

func main() {
	for i := 0; i < 0x20; i++ {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			store("loop", value)
			load("loop")
		}(i)
	}
	wg.Wait()
}
