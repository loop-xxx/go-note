package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func stdLog() {
	for i := int(0); i < 0x10; i++ {
		time.Sleep(2 * time.Second)
		log.Println("hello, world")
	}
}

func fileLog() {
	if fp, err := os.OpenFile("./log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644); err == nil {
		defer fp.Close()
		log.SetOutput(fp)

		for i := int(0); i < 0x10; i++ {
			time.Sleep(2 * time.Second)
			log.Println("hello, world")
		}

		log.SetOutput(os.Stdin)
	}
}
func main() {
	fmt.Println("hello, loop")
	fileLog()
	stdLog()
}
