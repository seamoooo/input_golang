package main

import "fmt"

func main() {
	n := 13

	_, msg := isPrime(n)

	fmt.Println(msg)
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d id is not prime", n)
	}

	if n < 0 {
		return false, "negative number are not prime"
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("number is not prime")
		}
	}

	return true, fmt.Sprintf("%d is a prime number", n)
}
