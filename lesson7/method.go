package main

import "fmt"

type Electronic struct {
	Volt int
}

func (Electronic) printVolt() {
	fmt.Println("Default 5V")
}

type Computer struct {
	Electronic
	name        string
	company     string
	year        int
	displaySize float32
}

func (c Computer) printComp() {
	fmt.Printf("Name: %s\nCompany: %s\nYear: %d\nDisplay Size: %f", c.name, c.company, c.year, c.displaySize)
}

func (c *Computer) Init(nom, kampaniya string, yili int16, ekrani float64) {
	c.name = nom
	c.company = kampaniya
	c.year = int(yili)
	c.displaySize = float32(ekrani)
}

type son int

// method
// func (n *son) sqr() int {
// 	*n = (*n) * (*n)
// }

func main() {
	mac := Computer{}
	// a := son(34)
	mac.Init("Macbook", "Apple", 1234, 45)
	mac.Volt = 234
	mac.printVolt()
	mac.printComp()

	// a := son(34)
	// // sqr(&a)
	// a.sqr()
	// fmt.Println(a)
}
