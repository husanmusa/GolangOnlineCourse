package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	err := rdb.Set(ctx, "key", "val", 1*time.Minute).Err()
	if err != nil {
		panic(err)
	}
	//fmt.Println(val)

	//val, err := rdb.Do(ctx, "set", "key", "1").Text()
	//fmt.Println(val, err)
}
