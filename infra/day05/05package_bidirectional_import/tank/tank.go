package tank

import (
	"fmt"
	"github.com/loop-xxx/go-note/infra/day05/05package_bidirectional_import/bullet"
)

// Destroy func
func Destroy() {
	fmt.Println("boom")
}

// Fire func
func Fire() {
	bullet.Fly()
}
