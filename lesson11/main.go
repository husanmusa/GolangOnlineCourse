package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered channel
	ch := make(chan int)

	// write from channel
	go sqr(ch)
	fmt.Println(1)
	ch <- 12
	ch <- 12
	fmt.Println(2)

	fmt.Println(<-ch)

}

func sqr(ch chan int) {
	a := <-ch
	a = <-ch
	time.Sleep(1 * time.Second)
	fmt.Println(3)
	a *= a
	ch <- a
}
