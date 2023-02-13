package main

import (
	"fmt"
	"math"
)

func printNumber() chan int {
	result := make(chan int)

	go func() {
		result <- 2
		for i := 3; i < 100000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			fount := false
			for j := 3; j < l; j += 2 {
				if i%j == 0 {
					fount = true
					break
				}
			}
			if !fount {
				result <- 1
			}
		}
		close(result)
	}()
	return result
}

func main() {
	pn := printNumber()
	for n := range pn {
		fmt.Println(n)
	}
}
