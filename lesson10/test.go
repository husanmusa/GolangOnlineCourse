package main

import "fmt"

func main() {
	var son int
	var son2 int
	var sum int
	var nat int
	fmt.Print("enter  num:")
	fmt.Scanln(&son)
	fmt.Println("enter sec num:")
	fmt.Scanln(&son2)
	if son > son2 {
		nat = son
		son = son2
		son2 = nat
	}
	for i := son; i < son2; i++ {
		if i%3 == 0 {
			sum += i
		}

	}
	fmt.Println(sum)
}
