package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)

	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
				// case out <- num: これが時間がかかっているという想定
			}
		}
		close(out)
		fmt.Printf("generator closed")
	}()

	return out
}

func main() {
	// context.WithDeadline関数を使うことで、
	//指定された時刻に自動的にDoneメソッドチャネルがcloseされるcontextを作成することができます。
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	gen := generator(ctx, 1)

	wg.Add(1)

LOOP:
	for i := 0; i < 5; i++ {
		select {
		case result, ok := <-gen:
			if ok {
				fmt.Println(result)
			} else {
				fmt.Println("timeout")
				break LOOP
			}
		}
	}
	// すでにDoneメソッドチャネルがcloseされているときに呼ばれたら、何もしない
	cancel()

	wg.Wait()
}
