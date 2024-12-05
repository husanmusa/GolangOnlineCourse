package main

import "fmt"

func main() {
	// slice == dynamic array

	// leng, cap
	var a = make([]int16, 45)

	// a = append(a, "HELLO")
	// a = append(a, "HELLO")
	// a = append(a, "HELLO")
	// a[4] = "salom"

	// a[5] = "panic"
	a[9] = 5637
	fmt.Println(a, len(a), cap(a))

	a = append(a, 1324)

	fmt.Println(a, len(a), cap(a))
}
