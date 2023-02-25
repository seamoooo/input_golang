package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running")

	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	buffer := make([]byte, 1000)

	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("receive from %v: %v \n", remoteAddress, length)

		_, err = conn.WriteTo([]byte("hello server"), remoteAddress)

		if err != nil {
			panic(err)
		}
	}
}
