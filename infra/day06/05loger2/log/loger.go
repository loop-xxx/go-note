package log

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/loop-xxx/go-note/infra/day06/05loger2/print"
	"github.com/loop-xxx/go-note/infra/day06/05loger2/utils"
)

//loger print level
const (
	DEBUG print.Level = iota
	INFO
	WARING
	ERROR
)

const runLog = string("run.log")
const minSize = int64(0x1000)
const minCheckduration = 10 * time.Second
const logerMaxPrintLevl = ERROR

// Loger interface
type Loger interface {
	Debugln(...interface{})
	Infoln(...interface{})
	Waringln(...interface{})
	Errorln(...interface{})
	Debugfl(string, ...interface{})
	Infofl(string, ...interface{})
	Waringfl(string, ...interface{})
	Errorfl(string, ...interface{})
	Close()
}

// CommonLoger ...
type CommonLoger struct {
	printer       print.Printer
	splitsize     int64
	logdir        string
	logname       string
	checkduration time.Duration
	lasttime      time.Time
}

// ConsoleLoger ...
type ConsoleLoger struct {
	printer print.Printer
}

func (lp *CommonLoger) splitlog(nowtime time.Time) {
	if err := os.Rename(path.Join(lp.logdir, runLog), utils.GenerateFileName(lp.logdir, lp.logname, lp.lasttime, nowtime)); err == nil {
		lp.lasttime = nowtime
	} else {
		panic(err)
	}
}

func (lp *CommonLoger) installprinter(level print.Level) {
	if head, err := os.OpenFile(path.Join(lp.logdir, runLog), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err == nil {
		lp.printer = print.PrinterConstruction(level, head)
	} else {
		panic(err)
	}
}

func (lp *CommonLoger) replaceprinter() {
	if head, err := os.OpenFile(path.Join(lp.logdir, runLog), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err == nil {
		lp.printer = lp.printer.Dup(head)
	} else {
		panic(err)
	}
}

func (lp *CommonLoger) build() (nowtime time.Time) {
	nowtime = time.Now()
	if nowtime.Sub(lp.lasttime) >= lp.checkduration {
		if fi, _ := lp.printer.Headinfo(); fi.Size() >= lp.splitsize {
			//close old printer
			lp.printer.Close()
			//split log
			lp.splitlog(nowtime)
			//replace new printer
			lp.replaceprinter()
		}
	}

	return
}

//Close ...
func (lp *CommonLoger) Close() {
	lp.printer.Close()
	lp.splitlog(time.Now())
	fmt.Println("[ConmonLoger INFO]:closed...")
}

//Close ...
func (lp *ConsoleLoger) Close() {
	fmt.Println("[ConsoleLoger INFO]:closed...")
}

//CommonLogerConstruction ...
func CommonLogerConstruction(level print.Level, logdir string, logname string,
	splitsize int64, checkduration time.Duration) (loger Loger, err error) {
	//print level default = DEBUG
	if level > logerMaxPrintLevl {
		level = DEBUG
		fmt.Println("[CommonLoger WARING]:level out of range, set level = DEBUG")
	}
	if splitsize < minSize {
		splitsize = minSize
		fmt.Printf("[CommonLoger WARING]:splitsize < %v, set splitsize = %v\n", minSize, minSize)
	}
	if checkduration < minCheckduration {
		checkduration = minCheckduration
		fmt.Printf("[CommonLoger WARING]:checkduration < %v, set checkduration = %v\n", minCheckduration, minCheckduration)
	}

	//check logdir
	err = errors.New("logdir is invaild")
	if utils.DirIsValid(logdir) {

		//check logname
		err = errors.New("log name is invaild")
		if utils.FileNameIsValid(logname) {

			err = nil
			lp := &CommonLoger{
				logdir:        logdir,
				logname:       logname,
				splitsize:     splitsize,
				checkduration: checkduration,
				lasttime:      time.Now(),
			}
			//install pointer
			lp.installprinter(level)
			loger = lp
		}
	}

	return
}

//ConsoleLogerConstruction ...
func ConsoleLogerConstruction(level print.Level) (loger Loger) {
	if level > logerMaxPrintLevl {
		level = DEBUG
		fmt.Println("[ConsoleLoger WARING]:level out of range, set level = DEBUG")
	}
	lp := &ConsoleLoger{}
	lp.printer = print.PrinterConstruction(level, os.Stdout)
	loger = lp

	return
}

//Debugln ...
func (lp *CommonLoger) Debugln(objs ...interface{}) {
	nowtime := lp.build()
	lp.printer.Println(DEBUG, utils.GenerateLogPrefix("DEBUG", nowtime), objs...)
}

//Debugln ...
func (lp *ConsoleLoger) Debugln(objs ...interface{}) {
	lp.printer.Println(DEBUG, utils.GenerateLogPrefix("DEBUG", time.Now()), objs...)
}

//Infoln ...
func (lp *CommonLoger) Infoln(objs ...interface{}) {
	nowtime := lp.build()
	lp.printer.Println(INFO, utils.GenerateLogPrefix("INFO", nowtime), objs...)
}

//Infoln ...
func (lp *ConsoleLoger) Infoln(objs ...interface{}) {
	lp.printer.Println(INFO, utils.GenerateLogPrefix("INFO", time.Now()), objs...)
}

//Waringln ...
func (lp *CommonLoger) Waringln(objs ...interface{}) {
	nowtime := lp.build()
	lp.printer.Println(WARING, utils.GenerateLogPrefix("WARING", nowtime), objs...)
}

//Waringln ...
func (lp *ConsoleLoger) Waringln(objs ...interface{}) {
	lp.printer.Println(WARING, utils.GenerateLogPrefix("WARING", time.Now()), objs...)
}

//Errorln ...
func (lp *CommonLoger) Errorln(objs ...interface{}) {
	nowtime := lp.build()
	lp.printer.Println(ERROR, utils.GenerateLogPrefix("ERROR", nowtime), objs...)
}

//Errorln ...
func (lp *ConsoleLoger) Errorln(objs ...interface{}) {
	lp.printer.Println(ERROR, utils.GenerateLogPrefix("ERROR", time.Now()), objs...)
}

//Debugfl ...
func (lp *CommonLoger) Debugfl(format string, objs ...interface{}) {
	nowtime := lp.build()
	lp.printer.Printfl(DEBUG, utils.GenerateLogPrefix("DEBUG", nowtime), format, objs...)
}

//Debugfl ...
func (lp *ConsoleLoger) Debugfl(format string, objs ...interface{}) {
	lp.printer.Printfl(DEBUG, utils.GenerateLogPrefix("DEBUG", time.Now()), format, objs...)
}

//Infofl ...
func (lp *CommonLoger) Infofl(format string, objs ...interface{}) {
	nowtime := lp.build()
	lp.printer.Printfl(INFO, utils.GenerateLogPrefix("INFO", nowtime), format, objs...)
}

//Infofl ...
func (lp *ConsoleLoger) Infofl(format string, objs ...interface{}) {
	lp.printer.Printfl(INFO, utils.GenerateLogPrefix("INFO", time.Now()), format, objs...)
}

//Waringfl ...
func (lp *CommonLoger) Waringfl(format string, objs ...interface{}) {
	nowtime := lp.build()
	lp.printer.Printfl(WARING, utils.GenerateLogPrefix("WARING", nowtime), format, objs...)
}

//Waringfl ...
func (lp *ConsoleLoger) Waringfl(format string, objs ...interface{}) {
	lp.printer.Printfl(WARING, utils.GenerateLogPrefix("WARING", time.Now()), format, objs...)
}

//Errorfl ...
func (lp *CommonLoger) Errorfl(format string, objs ...interface{}) {
	nowtime := lp.build()
	lp.printer.Printfl(ERROR, utils.GenerateLogPrefix("ERROR", nowtime), format, objs...)
}

//Errorfl ...
func (lp *ConsoleLoger) Errorfl(format string, objs ...interface{}) {
	lp.printer.Printfl(ERROR, utils.GenerateLogPrefix("ERROR", time.Now()), format, objs...)
}
