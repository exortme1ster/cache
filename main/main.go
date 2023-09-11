package main

import (
	"context"
	"fmt"
	"github.com/exortme1ster/cache"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*10)
	c := cache.New()
	c.Set("1", 20, time.Second*3)

	for {
		select {
		case <-time.After(time.Second * 1):
			c.Set("2", 20, time.Second*10)
			fmt.Println(c)
			fmt.Println(c.Get("1"), c.Get("2"))
			c.Delete("2")
			fmt.Println(c.Get("1"), c.Get("2"))

		case <-ctx.Done():
			fmt.Println("deadline exceeded")
			return
		}
	}

}
