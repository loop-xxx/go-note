package main

import "fmt"

type animal struct {
	name string
	age  uint32
}

func (pa *animal) move() {
	fmt.Printf("%s move~\n", (*pa).name)
}

type dog struct {
	animal
	food string
}

func (pd *dog) call() {
	fmt.Printf("%s wang~\n", (*pd).name)
}

func main() {
	fmt.Println("hello, loop")
	var dd = dog{
		animal: animal{"white", 0x10},
		food:   "bone",
	}
	fmt.Println(dd)
	fmt.Printf("%s, %d, %s\n", dd.name, dd.age, dd.food)
	dd.move()
	dd.call()
}
