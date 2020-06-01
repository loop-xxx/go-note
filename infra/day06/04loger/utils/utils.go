package utils

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

//RunInfoStr 获取当前运行函数被调用的位置,以string类型返回
func RunInfoStr() (runInfo string) {
	if pc, file, line, ok := runtime.Caller(3); ok {
		runInfo = fmt.Sprintf("%s:%s(%d)", path.Base(file), strings.Split(runtime.FuncForPC(pc).Name(), ".")[1], line)
	}
	return
}

//TimeInfoStr 获取当前系统时间以string类型返回
func TimeInfoStr() (timeInfo string) {
	timeInfo = time.Now().Format("2006-01-02 15:04:05.000")
	return
}

//TimeInfoStr2 获取当前系统时间以string类型返回
func TimeInfoStr2() (timeInfo string) {
	timeInfo = time.Now().Format("2006-01-02-150405")
	return
}
