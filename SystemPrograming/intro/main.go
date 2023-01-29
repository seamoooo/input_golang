package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	data := []byte{0x0, 0x0, 0x27, 0x27}

	var i int32
	var o int32
	// byte oderはlitte edianかbig edian かを指定
	binary.Read(bytes.NewReader(data), binary.LittleEndian, &i)
	binary.Read(bytes.NewReader(data), binary.BigEndian, &o)
	fmt.Printf("data: %d\n", i)
	fmt.Printf("data: %d\n", o)
}
