package main

import "fmt"

type Product struct {
	name, category string
	price          float32
}

func (p Product) getName() string {
	return p.name
}

type Service struct {
	description string
	duration    int
	price       float32
}

func (s Service) getName() string {
	return s.description
}

func (s Service) getCost(isYear bool) float32 {
	if isYear {
		return float32(s.duration) * s.price * 0.90
	}
	return s.price
}

type Expence interface {
	// getName() string
}

func main() {
	
	expences := []any{
		Product{"Ko'ylak", "kiyim", 10},
		Product{"Iphone", "Elektronika", 800},
		Service{"Gilam Yuvish", 1, 100},
		12,
		"SAlom",
		true,
	}

	for _, expence := range expences {
		switch p := expence.(type) {
		case Product:
			// fmt.Println(expence.getName())
			fmt.Println(p.category, p.name)
		case Service:
			fmt.Println("Boshqa service")
		}
	}
}
