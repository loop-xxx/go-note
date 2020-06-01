package main

import (
	"github.com/loop-xxx/go-note/infra/day06/04loger/loger"
)

const (
	_ int64 = 1 << (10 * iota)
	kb
	mb
	gb
)

func main() {

	cl := loger.NewConsoleLoger(loger.DEBUG)
	cl.Debugln("hello", "world")
	cl.Infoln("hello", "world")
	cl.Waringln("hello", "world")
	cl.Errorln("hello", "world")

	fl := loger.NewCommonLoger(loger.DEBUG, "./testlog", "loger", 10*mb)
	for {
		fl.Debugln("hello", "world")
		fl.Infoln("hello", "world")
		fl.Waringln("hello", "world")
		fl.Errorln("hello", "world")
	}
}
