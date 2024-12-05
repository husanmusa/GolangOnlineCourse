package main

import "fmt"

type addd func(a, b int) int

// anonymous functions
func main() {

	// fmt.Printf("%T, %v\n", res, res)
	// fmt.Printf("%T, %v", res(123,64), res(12,43))

	// greeting := func () {
	// 	fmt.Println("Hello")
	// }

	// greeting()

	// a := greeting1

	// a()
	// greeting1()

	// a()

	// a()
	// add := func(a, b int) int {
	// 	return a + b
	// }

	// mul := func(a, b int) int {
	// 	return a * b
	// }

	// sqrPrintwithDouble(add)()

	a := []byte{'a', 'b', 'c', 'd'}

	c := a
	a = append(a, 'z')
	z := 5
	b := z

	z = 2

	fmt.Println(b, z)

	fmt.Println(c, a)

	// RespFunc(4,5, "mul")() // -> 20
	// fmt.Println(2,34,43,45,546,56,67,657,657,65,7,65)
	sum(true, 4, 5, 54, 43, 43, 534, 534, 543)
	// fmt.Printf("", 2,32,42,34,234,4,5455,645)
}

func sum(check bool, a ...int) {
	z := []int{}
	fmt.Println(a)

	z = append(z, a...)
}

// func sqrPrintwithDouble(num addd) func() int {
// 	a := num(4, 5)
// 	return func() int {
// 		fmt.Println(a * a)

// 		return a + a
// 	}
// }

// func greeting() {
// 	fmt.Println("Hello1")
// }

// func RespFunc (a, b int, c string) func () int {

// }
