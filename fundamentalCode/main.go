package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	fmt.Println(ctx.Deadline()) // 0001-01-01 00:00:00 +0000 UTC false

	fmt.Println(time.Now()) // 2021-08-22 20:03:53.352015 +0900 JST m=+0.000228979
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)
	fmt.Println(ctx.Deadline()) // 2021-08-22 20:03:55.352177 +0900 JST m=+2.000391584 true
}
