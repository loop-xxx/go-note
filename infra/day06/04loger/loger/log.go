package loger

import (
	"fmt"
	"github.com/loop-xxx/go-note/infra/day06/04loger/utils"
	"os"

)

// log 结构体
type log struct {
	level Level
	fp    *os.File
}

//Debugln ...
func (log log) Debugln(objs ...interface{}) (err error) {

	if DEBUG >= log.level {
		prefix := fmt.Sprintf("[DEBUG][%s][%s]", utils.TimeInfoStr(), utils.RunInfoStr())
		_, err = fmt.Fprintln(log.fp, prefix, objs)
	}

	return
}

//Debugfl ...
func (log log) Debugfl(format string, objs ...interface{}) (err error) {
	if DEBUG >= log.level {
		farme := fmt.Sprintf("[DEBUG][%s][%s] %s\n", utils.TimeInfoStr(), utils.RunInfoStr(), format)
		_, err = fmt.Fprintf(log.fp, farme, objs...)
	}
	return
}

//Infoln ...
func (log log) Infoln(objs ...interface{}) (err error) {

	if INFO >= log.level {
		prefix := fmt.Sprintf("[INFO][%s][%s]", utils.TimeInfoStr(), utils.RunInfoStr())
		_, err = fmt.Fprintln(log.fp, prefix, objs)
	}
	return
}

//Infofl ...
func (log log) Infofl(format string, objs ...interface{}) (err error) {
	if INFO >= log.level {
		farme := fmt.Sprintf("[INFO][%s][%s] %s\n", utils.TimeInfoStr(), utils.RunInfoStr(), format)
		_, err = fmt.Fprintf(log.fp, farme, objs...)
	}
	return
}

//Waringln ...
func (log log) Waringln(objs ...interface{}) (err error) {

	if WARING >= log.level {
		prefix := fmt.Sprintf("[WARING][%s][%s]", utils.TimeInfoStr(), utils.RunInfoStr())
		_, err = fmt.Fprintln(log.fp, prefix, objs)
	}
	return
}

//Waringfl ...
func (log log) Waringfl(format string, objs ...interface{}) (err error) {
	if WARING >= log.level {
		farme := fmt.Sprintf("[WARING][%s][%s] %s\n", utils.TimeInfoStr(), utils.RunInfoStr(), format)
		_, err = fmt.Fprintf(log.fp, farme, objs...)
	}
	return
}

//Errorln ...
func (log log) Errorln(objs ...interface{}) (err error) {

	if ERROR >= log.level {
		prefix := fmt.Sprintf("[ERROR][%s][%s]", utils.TimeInfoStr(), utils.RunInfoStr())
		_, err = fmt.Fprintln(log.fp, prefix, objs)
	}
	return
}

//Errorfl ...
func (log log) Errorfl(format string, objs ...interface{}) (err error) {
	if ERROR >= log.level {
		farme := fmt.Sprintf("[ERROR][%s][%s] %s\n", utils.TimeInfoStr(), utils.RunInfoStr(), format)
		_, err = fmt.Fprintf(log.fp, farme, objs...)
	}
	return
}
