package main

import "testing"

func Test_isPrime(t *testing.T) {

	// methodへのparameterごとでテストを書いていくと冗長になってう
	// そこでsliceでparameterとreturnを事前に定義してtesttableを設定する
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number"},
		{"zoro", 0, false, "0 is not prime"},
		{"negative", -1, false, "negative number are not prime"},
		{"not prime", 10, false, "number is not prime"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s expected true but get false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
