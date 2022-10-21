package main

import "fmt"

type Person struct {
	LastName  string
	FirstName string
	Age       int
}

// レシーバは先頭一文字で定義するのが慣習
func (p Person) String() string {
	return fmt.Sprintf("%s %s :年齢%d歳", p.LastName, p.FirstName, p.Age)
}

func main() {

	takashi := Person{
		LastName:  "takashi",
		FirstName: "inoue",
		Age:       45,
	}

	fmt.Println(takashi.String())
}
