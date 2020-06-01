package main

import "fmt"

func step(n, m uint) (count uint) {
	if n != 0 {
		for i := uint(1); i <= m; i++ {
			if n-i >= m {
				count += step(n-i, m)
			} else {
				count += step(n-i, n-i)
			}
		}
	} else {
		count = 1
	}
	return
}
func main() {
	fmt.Println(step(4, 2))
}
