package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// TODO hi smth
	ctx := context.Background()

	ctx = context.WithValue(ctx, "hello", "world")

	// ctx, cancel := context.WithCancel(ctx)
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	cancel()
	// }()

	// f(ctx)
	ctx, cancelFunc := context.WithTimeout(ctx, 2*time.Second)
	time.Now().AddDate(0, 0, 2)
	defer cancelFunc()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("after 3 second")
	case <-ctx.Done():
		fmt.Print(ctx.Err())
	}
}

func f(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("in ticker", t)
		case <-ctx.Done():
			fmt.Print("Done: ", ctx.Err())
			return
		}
	}
}
