package main

import (
	"sync"
)

var dataOnce sync.Once
var data map[string]string

func loaddata() {
	data = map[string]string{"loop": "hi", "luna": "bey"}
}
func get(key string) string {
	dataOnce.Do(loaddata)
	return data[key]
}

var dataMutex sync.Mutex
var data2 map[string]string //假设有volatile 关键字修饰

func loaddata2() {
	temp := map[string]string{"loop": "hi", "luna": "bey"}
	data2 = temp //假设这句是原子操作
}

func get2(key string) string {
	if data2 == nil {
		dataMutex.Lock()
		if data2 == nil {
			loaddata()
		}
		dataMutex.Unlock()
	}

	return data[key]
}

func main() {
	get("loop") // 相当于get2()假设成立的情况
}
