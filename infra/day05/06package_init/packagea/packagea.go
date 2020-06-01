package packagea

import "fmt"

import _ "github.com/loop-xxx/go-note/infra/day05/06package_init/packageb"

func init() {
	fmt.Println("packagea init")
}
