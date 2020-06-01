package main

import "fmt"

type person struct {
	name string
}

type student struct {
	person
	score uint32
}

func (p person) show() {
	fmt.Printf("I'm %s\n", p.name)
}

type stuManager map[uint64]student

func (s stuManager) addStu() {
	fmt.Println("add student")

	var id uint64
	fmt.Scanln(&id)
	if _, isExist := s[id]; !isExist {
		var stu student
		fmt.Scanln(&stu.name)
		fmt.Scanln(&stu.score)
		s[id] = stu
	}
}
func (s stuManager) delStu() {
	fmt.Println("delete student")

	var id uint64
	fmt.Scanln(&id)
	delete(s, id)
}
func (s stuManager) modifyStu() {
	fmt.Println("modify student")

	var id uint64
	fmt.Scanln(&id)
	if _, isExist := s[id]; isExist {
		var stu student
		fmt.Scanln(&stu.name)
		fmt.Scanln(&stu.score)
		s[id] = stu
	}
}

func (s stuManager) listStu() {
	fmt.Println("list student")

	for id, stu := range s {
		stu.show()
		fmt.Printf("%d:%d\n", id, stu.score)
	}
}

func getOpt() (opcode int) {
	fmt.Println(`	welcome student manager system!
	1 add student
	2 delete student
	3 modify student
	4 list student
	5 exit`)
	fmt.Scanln(&opcode)
	return
}

func main() {
	fmt.Println("hello loop")
	var stum stuManager = make(map[uint64]student, 0x10)

	for opcode := getOpt(); opcode != 5; opcode = getOpt() {
		switch opcode {
		case 1:
			stum.addStu()
		case 2:
			stum.delStu()
		case 3:
			stum.modifyStu()
		case 4:
			stum.listStu()
		default:
			fmt.Println("opcode invaild!")
		}
	}

}
