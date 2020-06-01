package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// DirIsValid ...
func DirIsValid(dirpath string) (IsValid bool) {

	IsValid = false
	if fileInfo, err := os.Stat(dirpath); err == nil {
		IsValid = fileInfo.IsDir()
	}
	return
}

//FileNameIsValid ...
func FileNameIsValid(name string) (IsValid bool) {
	IsValid = false
	if strings.Index(name, "/") == -1 &&
		strings.Index(name, "\\") == -1 &&
		strings.Index(name, ":") == -1 {
		IsValid = true
	}

	return
}

/**
修改日志命名格式
*/
const filenameTimeformat = string("20060102150405")

//GenerateFileName ...
func GenerateFileName(dir, name string, last, now time.Time) (filename string) {
	filename = fmt.Sprintf("%s/%s[%s-%s].log", dir, name, last.Format(filenameTimeformat), now.Format(filenameTimeformat))
	return
}

/**
修改日志前缀格式
*/
const logTimeFormat = string("2006-01-02 15:04:05.000")

//GenerateLogPrefix ...
func GenerateLogPrefix(level string, nowtime time.Time) (logprefix string) {
	logprefix = fmt.Sprintf("[%s][%s]", level, nowtime.Format(logTimeFormat))
	return
}
