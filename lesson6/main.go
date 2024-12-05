package main

import "fmt"

func main() {
	// var q *int

	// fmt.Println(q)

	// var arr = [5]int{1, 2, 3, 4, 5}

	// var p *[5]int

	// p = &arr

	// justFunc(p)

	// fmt.Printf("%p %p\n", &arr, &arr[0])

	slc := make([]int, 3, 5)

	justFunc(slc)
	fmt.Println(slc)
	fmt.Printf("%v %p %p %p\n", slc, slc, &slc, &slc[0])
	// slc = append(slc, 4)
	// fmt.Printf("%v %p %p %p", slc, slc, &slc, &slc[0])
}

func justFunc(n []int) {
	n[2] = 12
	fmt.Printf("%v %p %p %p\n", n, n, &n, &n[0])
	n = append(n, 3)
	fmt.Println(n)
	fmt.Printf("%v %p %p %p\n", n, n, &n, &n[0])
}
