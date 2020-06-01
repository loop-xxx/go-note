package main

import (
	"github.com/loop-xxx/go-note/infra/day05/11log/log"
)

func test() {
	//log.Traceln("hello, loop")
	log.Debugln("hello, loop")
	log.Infoln("hello, loop")
	log.Waringln("hello, loop")
	log.Errorln("hello, loop")

	//log.Tracef("%d\n", 0xff)
	log.Waringf("%d\n", 0xff)
	log.Infof("%d\n", 0xff)
	log.Errorf("%d\n", 0xff)
}

func main() {

	if errno := log.Open("./log.txt"); errno == 0 {
		defer log.Close()

		log.SetLevel(log.DEBUG)
		test()
	}
}
