package bullet

import "fmt"

//import "github.com/loop/day05/05package_bidirectional_import/tank"

// Fly func
func Fly() {
	fmt.Println("fly")
}

//Attack func
func Attack() {
	//tank.Destroy()
	fmt.Println("tank.Destroy():can not double import")
}
