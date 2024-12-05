package main

import (
	"fmt"

	"github.com/husanmusa/goOnline/lesson8/mathem/rand"

	"github.com/husanmusa/goOnline/lesson8/mathem"
)

func main() {
	fmt.Println(mathem.Sqr(rand.RandInt(100)))
}
