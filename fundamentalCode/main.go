package main

import (
	"fmt"
	"time"
)

type Book struct {
	Title     string
	Author    string
	Publisher string
	Day       time.Time
	ISBN      string
}

func main() {

	jst, _ := time.LoadLocation("Asia/Tokyo")

	book := Book{
		Title:     "海の向こう",
		Author:    "johe",
		Publisher: "johe",
		Day:       time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
		ISBN:      "hogeoge",
	}

	fmt.Println(book)

	b1 := new(Book)
	var b2 Book
	// pointerの初期化、初期値が設定できる
	b3 := &Book{
		Title:     "海の向こう",
		Author:    "johe",
		Publisher: "johe",
		Day:       time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
		ISBN:      "hogeoge",
	}

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)

	b3.Title = "丘の向こうに"
	fmt.Println(b3)
}
