package loger

import (
	"fmt"
	"github.com/loop-xxx/go-note/infra/day06/04loger/utils"
	"os"

)

type flog struct {
	log
	count     uint32
	splitsize int64
	logpath   string
	logname   string
}

func (f *flog) Debugln(objs ...interface{}) (err error) {
	f.check()
	err = f.log.Debugln(objs...)
	return
}
func (f *flog) Infoln(objs ...interface{}) (err error) {
	f.check()
	err = f.log.Infoln(objs...)
	return
}
func (f *flog) Waringln(objs ...interface{}) (err error) {
	f.check()
	err = f.log.Waringln(objs...)
	return
}
func (f *flog) Errorln(objs ...interface{}) (err error) {
	f.check()
	err = f.log.Errorln(objs...)
	return
}
func (f *flog) Debugfl(format string, objs ...interface{}) (err error) {
	f.check()
	err = f.log.Debugfl(format, objs...)
	return
}
func (f *flog) Infofl(format string, objs ...interface{}) (err error) {
	f.check()
	err = f.log.Infofl(format, objs...)
	return
}
func (f *flog) Waringfl(format string, objs ...interface{}) (err error) {
	f.check()
	err = f.log.Waringfl(format, objs...)
	return
}
func (f *flog) Errorfl(format string, objs ...interface{}) (err error) {
	f.check()
	err = f.log.Errorfl(format, objs...)
	return
}

func (f *flog) open() {
	if fp, err := os.OpenFile(fmt.Sprintf("%s/%s", f.logpath, "run.log"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err == nil {
		f.log.fp = fp
	} else {
		panic("logpath invaild")
	}
}
func (f *flog) close() {
	f.log.fp.Close()
}

func (f *flog) split() {
	f.close()
	if err := os.Rename(fmt.Sprintf("%s/%s", f.logpath, "run.log"), fmt.Sprintf("%s/%s%s.log", f.logpath, f.logname, utils.TimeInfoStr2())); err != nil {
		panic("rename faild")
	}
	f.open()
}
func (f *flog) check() {
	f.count++
	if f.count == 0x100 {
		f.count = 0
		fileInfo, _ := f.fp.Stat()
		if fileInfo.Size() > f.splitsize {
			f.split()
		}
	}
}
