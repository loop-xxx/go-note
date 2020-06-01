package loger

import (
	"fmt"
	"os"
)

// Level log打印日志等级类型
type Level = uint8

//Level常量
const (
	DEBUG Level = iota
	INFO
	WARING
	ERROR
)

// Loger ...
type Loger interface {
	Debugln(...interface{}) error
	Infoln(...interface{}) error
	Waringln(...interface{}) error
	Errorln(...interface{}) error
	Debugfl(string, ...interface{}) error
	Infofl(string, ...interface{}) error
	Waringfl(string, ...interface{}) error
	Errorfl(string, ...interface{}) error
}

// NewConsoleLoger ...
func NewConsoleLoger(level Level) (l Loger) {
	if level > ERROR {
		level = DEBUG
		fmt.Println("[WARING] level out of range -> set level = [DEBUG]")
	}
	l = log{level: level, fp: os.Stdout}
	return
}

// NewCommonLoger ...
func NewCommonLoger(level Level, logpath, logname string, splitsize int64) (l Loger) {
	if level > ERROR {
		level = DEBUG
		fmt.Println("[WARING] level out of range -> set level = [DEBUG]")
	}
	fl := flog{log: log{level: level}, splitsize: splitsize, logpath: logpath, logname: logname}
	fl.open()
	l = &fl
	return
}


