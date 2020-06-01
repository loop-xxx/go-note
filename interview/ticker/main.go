package main

import (
	"fmt"
	"time"
)

func main(){
	/*
	tick := time.Tick(time.Millisecond)
	for t := range tick{
		time.Sleep(time.Second * 4)
		fmt.Println(t)
	}
	 */
	timer := time.After(time.Second)
	t := <- timer
	fmt.Println(t)
}
