package main

import "fmt"

// 具体的な実装を提供しない抽象型
type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	ID   string
}

func (p Person) ToString() string {
	return fmt.Sprintf("%v", p.Name)
}

type Dog struct {
	Name string
	ID   string
}

func (d Dog) ToString() string {
	return fmt.Sprintf("%v", d.Name)
}

func main() {
	// 例えばfor分で各クラスごとのNameを出力したい場合など、
	// 同じToStringを同じ形で扱うことが難しいのでinterface

	// interfaceを定義すると下記のようにsliceで同じように扱うことができる
	// 型安全のダックタイプ
	val := []Stringfy{
		&Person{Name: "John", ID: "21"},
		&Dog{Name: "dog", ID: "22"},
	}

	for _, v := range val {
		fmt.Println(v.ToString())
	}
}
