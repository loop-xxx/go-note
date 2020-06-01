package log

import (
	"fmt"
	"os"
)

// print level
const (
	ERROR = int(iota)
	WARING
	INFO
	//TRACE
	DEBUG
)

var level int
var writer *os.File

func init() {
	level = DEBUG
	/**
	fmt.Printf("%T, %v\n", ERROR, ERROR)
	fmt.Printf("%T, %v\n", WARING, WARING)
	fmt.Printf("%T, %v\n", INFO, INFO)
	fmt.Printf("%T, %v\n", DEBUG, DEBUG)
	*/
}

// SetLevel 设置打印等级
func SetLevel(l int) {
	level = l
}

// Open Open打印
func Open(path string) (errno int) {
	var err error

	errno = -1
	if writer, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644); err == nil {
		errno = 0
	}
	return
}

// Close Close打印
func Close() (errno int) {
	errno = -1
	if writer != nil {
		errno = 0
		writer.Close()
	}
	return
}

// Traceln trace level println
/**
func Traceln(obj ...interface{}) {
	if level >= TRACE {
		fmt.Fprint(writer, "[TRACE] ")
		fmt.Fprintln(writer, obj...)
	}
}
*/

// Tracef Trace level printf
/**
func Tracef(format string, obj ...interface{}) {
	if level >= TRACE {
		fmt.Fprint(writer, "[TRACE] ")
		fmt.Fprintf(writer, format, obj...)
	}
}
*/

// Debugln debug level println
func Debugln(obj ...interface{}) {
	if level >= DEBUG {
		fmt.Fprint(writer, "[DEBUG] ")
		fmt.Fprintln(writer, obj...)
	}
}

// Debugf debug level printf
func Debugf(format string, obj ...interface{}) {
	if level >= DEBUG {
		fmt.Fprint(writer, "[DEBUG] ")
		fmt.Fprintf(writer, format, obj...)
	}
}

// Infoln info level println
func Infoln(obj ...interface{}) {
	if level >= INFO {
		fmt.Fprint(writer, "[INFO] ")
		fmt.Fprintln(writer, obj...)
	}
}

// Infof info level printf
func Infof(format string, obj ...interface{}) {
	if level >= INFO {
		fmt.Fprint(writer, "[INFO] ")
		fmt.Fprintf(writer, format, obj...)
	}
}

// Waringln Waring level println
func Waringln(obj ...interface{}) {
	if level >= WARING {
		fmt.Fprint(writer, "[WARING] ")
		fmt.Fprintln(writer, obj...)
	}
}

// Waringf Waring level printf
func Waringf(format string, obj ...interface{}) {
	if level >= WARING {
		fmt.Fprint(writer, "[WARing] ")
		fmt.Fprintf(writer, format, obj...)
	}
}

// Errorln Error level println
func Errorln(obj ...interface{}) {
	if level >= ERROR {
		fmt.Fprint(writer, "[ERROR] ")
		fmt.Fprintln(writer, obj...)
	}
}

// Errorf Error level printf
func Errorf(format string, obj ...interface{}) {
	if level >= ERROR {
		fmt.Fprint(writer, "[ERROR] ")
		fmt.Fprintf(writer, format, obj...)
	}
}
