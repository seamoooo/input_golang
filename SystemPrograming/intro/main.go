package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer

	buffer.Write([]byte("hello world"))

	fmt.Println(buffer.String())
}
