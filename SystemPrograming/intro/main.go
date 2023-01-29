package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer1 bytes.Buffer
	bytes.NewBuffer([]byte{0x10})
	buffer3 := bytes.NewBufferString("初期文字列")
	fmt.Println(buffer1)
	fmt.Println(buffer3)
}
