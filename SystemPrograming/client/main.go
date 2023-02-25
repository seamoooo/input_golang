package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp4", "localhost:8888")
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Println("send to request")

	_, err = conn.Write([]byte("hello world client"))
	if err != nil {
		panic(err)
	}

	fmt.Println("recibe from srrver")
	buffer := make([]byte, 1500)

	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("recibe from %s\n", string(buffer[:length]))
}
