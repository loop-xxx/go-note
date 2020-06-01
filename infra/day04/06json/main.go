package main

import "fmt"
import "encoding/json"

type animal struct {
	Name string `json:"name" fmt:"name" math:"name"`
	Age  uint32 `json:"age"`
}
type dog struct {
	animal
	Food string `json:"food"`
}

func main() {
	fmt.Println("hello, world")
	/**
	var d1 = dog{
		animal: animal{"white", 0x10},
		Food:   "bone",
	}
	*/
	var d1 = dog{animal{"white", 0x10}, "bone"}

	if data, err := json.Marshal(d1); err == nil {
		fmt.Printf("%s\n", data)
	}

	var context = "{\"name\":\"black\", \"age\":17, \"food\":\"bone\" }"
	var d2 dog
	if err := json.Unmarshal([]byte(context), &d2); err == nil {

		fmt.Printf("%#v\n", d2)
	}
}
