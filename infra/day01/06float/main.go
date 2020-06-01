package main

import "fmt"
import "math"

func main() {
	var number1 float64 = 0.1234
	number2 := float32(0.1234)

	fmt.Printf("%T,%T\n", number1, number2)
	number1 = float64(number2)
	number2 = float32(number1)
	fmt.Println(number1, number2)

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
}
