package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	ch := make(chan int, len(a))

	for _, v := range a {

		// goルーチンが外側の変数に依存している場合は、
		// 変数のシャドーか引数を無銘関数に渡す必要がある
		go func(val int) {
			ch <- val * 2 // 引数を渡さなければここに到達する時点のvの値が幾つになっているかわからない
		}(v)
	}

	for i := 0; i < len(a); i++ {
		fmt.Print(<-ch, " ")
	}
}
