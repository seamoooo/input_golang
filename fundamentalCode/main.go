package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Print("無銘関数:", v, "", v2, "\n")
	}()

	v := 2
	var v2 int
	// mainもランタイム開始時にゴールーチンとして起動されるので、
	// selectを挟まないとデッドロックになる
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}

	fmt.Print("mainの最後:", v, "", v2, "\n")
}
