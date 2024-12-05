package main

import (
	"fmt"
)

func main() {
	// words := "gbhsdfyu4i5yg346fv9q805gv73458943hfj54983242598ytre90dvh8j2398452"
	// letters, numbers := counterLetter(words)

	// fmt.Printf("in %s, we have %d letters and %d numbers", words, letters, numbers)

	fmt.Println(factorial(5))
}

func counterLetter(word string) (int, int) {
	l, n := 0, 0

	for _, v := range word {
		
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			l++
		} else if v >= '0' && v <= '9' {
			n++
		}
	}

	return l, n
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
