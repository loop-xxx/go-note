package print

import (
	"fmt"
	"os"
)

//Level print_level
type Level = uint

//Printer ...
type Printer struct {
	level Level
	head  *os.File
}

//Println ...
func (p Printer) Println(textlevel Level, prefix string, objs ...interface{}) (err error) {

	if textlevel >= p.level {
		_, err = fmt.Fprintln(p.head, prefix, objs)
	}
	return
}

//Printfl ...
func (p Printer) Printfl(textlevel Level, prefix string, format string, objs ...interface{}) (err error) {
	if textlevel >= p.level {
		_, err = fmt.Fprintf(p.head, fmt.Sprintf("%s %s\n", prefix, format), objs...)
	}
	return
}

// PrinterConstruction ...
func PrinterConstruction(level Level, head *os.File) (p Printer) {
	p.level = level
	p.head = head
	return
}

// Dup ...
func (p Printer) Dup(head *os.File) (new Printer) {
	new.level = p.level
	new.head = head
	return
}

// Headinfo ...
func (p Printer) Headinfo() (fi os.FileInfo, err error) {
	fi, err = p.head.Stat()
	return
}

// Close ...
func (p Printer) Close() (err error) {
	err = p.head.Close()
	return
}
