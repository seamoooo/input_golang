package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open("/Users/shimohozumahozuma/project")
	if err != nil {
		panic(err)
	}

	fileInfos, err := dir.Readdir(-1)

	if err != nil {
		panic(err)
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			fmt.Printf("[Dir] %s\n", fileInfo.Name())
		} else {
			fmt.Printf("[%s]\n", fileInfo.Name())
		}
	}
}
