package main

import (
	"github.com/loop-xxx/go-note/infra/day05/05package_bidirectional_import/bullet"
	"github.com/loop-xxx/go-note/infra/day05/05package_bidirectional_import/tank"
)

//golang中不支持包与包之间互相import
func main() {
	tank.Fire()
	bullet.Attack()
}
