package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	fmt.Println(len(ch), cap(ch))

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 4
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// for range -> chan
	// select case
	// Mutex - mutual expression , lock unlock
	
	

	fmt.Println(len(ch), cap(ch))
}
