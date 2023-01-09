package main

import "testing"

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)

	if result {
		t.Errorf("expexted false")
	}

	if msg != "0 id is not prime" {
		t.Error("wrong message received:", msg)
	}
}
