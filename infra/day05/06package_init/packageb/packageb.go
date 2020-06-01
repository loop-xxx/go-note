package packageb

import "fmt"

import _ "github.com/loop-xxx/go-note/infra/day05/06package_init/packagec"

func init() {
	fmt.Println("packageb init")
}
