package main

import (
	"fmt"
	"strings"
)

func main() {
	// Ichma-ich funksiya
	transformer := func(input string) func(op1 func(string) string) func(op2 func(string) string) string {
		return func(op1 func(string) string) func(op2 func(string) string) string {
			// Birinchi funksiya qo‘llanadi
			intermediate := op1(input)
			return func(op2 func(string) string) string {
				// Ikkinchi funksiya qo‘llanadi
				final := op2(intermediate)
				return final
			}
		}
	}

	// Harf manipulyatsiyasi funksiyalari
	toUpperCase := func(s string) string {
		return strings.ToUpper(s)
	}

	reverseString := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	// Input so'z
	word := "helloa"

	// Transformatsiya jarayoni
	result := transformer(word)(toUpperCase)(reverseString)

	// Natijani chiqarish
	fmt.Printf("Input: %s\n", word)
	fmt.Printf("Natija: %s\n", result)
}
