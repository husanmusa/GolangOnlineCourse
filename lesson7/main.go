package main

import "fmt"

type Driver struct {
	name       string
	gen        byte
	birthYear  int16
	hasLisence bool
}

type Vehicle struct {
	number string
	color  string
	year   int
	name   string
}

// Structure
type Car struct {
	Vehicle
	isNew  *bool
	driver *Driver
}

// struct nested x3.

// alias
// type son = int

func main() {
	driver := Driver{}
	driver.name = "kimdir"

	fmt.Println(driver)
	var car = new(Car)

	car.driver = &driver
	var b = true
	car.isNew = &b
	car.driver.name = "Hello"
	fmt.Println(car)
	return
	var damas = &Car{
		Vehicle: Vehicle{
			number: "01A111AA",
			color:  "white",
			year:   2000,
		},

		// isNew:  false,
		driver: &driver,
	}

	// var z = struct {
	// 	firstName string
	// 	lastName  string
	// 	age       int
	// }{
	// 	firstName: "Husan",
	// 	lastName:  "Musa",
	// 	age:       99,
	// }

	damas.name = "Blue line"
	damas.driver.hasLisence = true

	damas = nil
	if damas == nil {
		return
	}

	fmt.Println(damas)
	if damas != nil {
		fmt.Printf("%+v %s \n", damas, damas.number)
	}
	// fmt.Println(damas.driver)
}
