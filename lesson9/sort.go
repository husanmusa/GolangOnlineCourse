package main

import (
	"fmt"
	"sort"
)

type sonlar []float32

func (s sonlar) Len() int {
	return len(s)
}
func (s sonlar) Less(a int, b int) bool {
	if s[a] < s[b] {
		return true
	}
	return false
}
func (s sonlar) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}



func main() {

	var n = sonlar{1, 34, 4, 234, 23, 144, 3534}

	var a interface{} = sonlar{1,2,3,4}

	sort.Sort(n)

	fmt.Println(n, a.)
}
