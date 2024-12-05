package main

import (
	"fmt"
)

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}
func update(px *int) {
	*px = 20
}

func main() {
	x := 10
	var c *int = &x
	failedUpdate(c)
	fmt.Println(x) // prints 10
	update(&x)
	fmt.Println(x, *c) // prints 20
}
