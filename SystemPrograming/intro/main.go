package main

import "fmt"

// インターフェイスの定義、interfaceは構造体などの具体型が持つべきメソッドを表現するための言語機能
type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("hello my name id %s\n", g.name)
}

func main() {
	// GreeterはTalkerインターフェイスを満たすので、インターフェイス型の変数にポインターを代入することができる
	var talker Talker = &Greeter{"wozozo"}

	talker.Talk()
}
