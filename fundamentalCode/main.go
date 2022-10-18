package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("ファイルが指定されてません")
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	data := make([]byte, 2048)

	for {
		count, err := f.Read(data)

		os.Stdout.Write(data[:count])

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			break
		}
	}
}
