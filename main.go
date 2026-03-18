package main

import (
	"fmt"
	"time"

	"github.com/ahmed-zein/go_rate_limiting/config"
	"github.com/ahmed-zein/go_rate_limiting/limiter"
	rl "github.com/ahmed-zein/go_rate_limiting/limiter/redis"
	"github.com/redis/go-redis/v9"
)

var L limiter.Limiter

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	cfg := &config.WindowBasedConfig{
		Limit:      10,
		WindowSize: 20 * time.Second,
	}
	L, _ = rl.NewSlidingWindowCounter("test", cfg, rdb)

	for range 10 {
		fmt.Printf("%+v\n", L)

		allow, _ := L.AllowN("1", 2)
		if allow {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")

		}

	}

}
