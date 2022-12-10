package main

import (
	"fmt"

	"github.com/seamoooo/input_golang/fundamentalCode/book"
)

func main() {
	// call factory of book structure
	b := book.NewAuthorlessBook("hi")
	fmt.Printf("%+v\n", b.Author)
}
