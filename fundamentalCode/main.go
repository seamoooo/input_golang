package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	st := []string{"hello", "world", "hoge"}
	gh := [...]string{"hello", "world", "hoge"}

	type grpc struct {
		proto  string
		method int
	}

	// g := grpc{
	// 	proto:  "tcp",
	// 	method: 3,
	// }

	for i, v := range st {
		fmt.Println(i, v)
	}
	for _, v := range gh {
		fmt.Println(v)
	}
	//　構造体は利用不可
	// for i, v := range g {
	// 	fmt.Println(i, v)
	// }

	m := map[string]int{
		"a": 1,
		"b": 2,
		"x": 3,
	}

	// mapの取り出しは順序を保証しない
	// hash dos の防止
	for i, v := range m {
		fmt.Println(i, v)
	}

	text := "i move to okinawa"

	// stringは1文字ごとのruneを返す
	// 先頭から順にアクセスする
	for i, v := range text {
		fmt.Println(i, v)
	}
}
