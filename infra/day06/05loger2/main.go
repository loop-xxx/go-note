package main

import (
	"github.com/loop-xxx/go-note/infra/day06/05loger2/log"
)

const (
	_ int64 = 1 << (iota * 10)
	kb
	mb
	gb
)

func test(loger log.Loger) {
	for {
		loger.Debugln("loop", "luna", "hello world")
		loger.Infoln("loop", "luna", "hello world")
		loger.Waringln("loop", "luna", "hello world")
		loger.Errorln("loop", "luna", "hello world")

		loger.Debugfl("%d:%d", 1, 0xff)
		loger.Infofl("%d:%d", 1, 0xff)
		loger.Waringfl("%d:%d", 1, 0xff)
		loger.Errorfl("%d:%d", 1, 0xff)
		//time.Sleep(2 * time.Second)
	}
}

func main() {

	//os.Stdout.Close()

	/**
	if loger, err := log.CommonLogerConstruction(log.INFO, "./testlog", "loop", 10*mb, 0); err == nil {
		test(loger)
	} else {
		fmt.Println(err)
	}
	*/
	{
		loger := log.ConsoleLogerConstruction(log.INFO)
		test(loger)
	}
}
