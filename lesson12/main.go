package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type mutex struct {
	m sync.RWMutex
	a int
}

var mx mutex

// for range -> chan
// select case
// Mutex - mutual expression, lock unlock

func main() {

	// for v := range ch {
	// 	fmt.Println(len(ch), cap(ch))
	// 	time.Sleep(10 * time.Second)
	// 	fmt.Printf("v: %d\n", v)
	// }

	// select {
	// case <-ch:
	// 	fmt.Println(len(ch), a, cap(ch))
	// 	fmt.Printf("v: %d\n", <-ch)
	// 	time.Sleep(3 * time.Second)
	// case <-time.Tick(5 * time.Second):
	// 	fmt.Println(time.Now())
	// }
	// a := 0

	// wg := sync.WaitGroup{}

	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)

	// 	go func() {
	// 		mx.m.Lock()
	// 		a++
	// 		mx.m.Unlock()
	// 		wg.Done()
	// 	}()
	// }
	// fmt.Println(a)
	// wg.Wait()
	// fmt.Println(a)

	var ch = make(chan int)

	go test(ch)

	// <-ch
	_, ok := <-ch
	fmt.Println(ok)
	time.Sleep(1 * time.Second)
	fmt.Println(<-ch)
	close(ch)
	_, ok = <-ch
	fmt.Println(ok)
	time.Sleep(1 * time.Second)
}

func test(ch chan int) {
	ch <- 2
	ch <- 1
}

func isOdd(ch chan int) {
	fmt.Printf("%p\n", ch)

	time.Sleep(1 * time.Second)
	for {
		a := rand.Intn(10)
		fmt.Println(a)
		if a%2 == 1 {
			ch <- a
			time.Sleep(5 * time.Second)
		}
	}
}
