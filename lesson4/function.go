package main

import "fmt"

func main() {
	// var x, y = 65, 645
	// greating(x, y)

	// r :=
	// fmt.Println(sqr(6))
	f, l := separate("Husan Musa")
	fmt.Printf("Ism: %s, fam: %s", f, l)



}

func greating(x int, y int) {
	// x := 5
	// y := 65
	fmt.Println("Hello World", x+y)
}

func sqr(a int) int {
	//...
	//...
	return a * a
}

func separate(fullname string) (string, string) {
	for i, v := range fullname {
		if v == ' ' { //   [:)            []
			return fullname[:i], fullname[i+1:]
		}
	}

	return "", ""
}

func sep(slc []int) ([]int, []int) {
}
