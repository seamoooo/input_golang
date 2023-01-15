package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

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

func Test_promt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	prompt()

	_ = w.Close()

	os.Stdout = oldOut
	out, _ := io.ReadAll(r)

	if string(out) != "->" {
		t.Errorf("incoreect promt: expected -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut
	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Is it Prime?") {
		t.Errorf("expectd Is it Prime? actually get  %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	test := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "pleace enter number"},
		{name: "prime number", input: "7", expected: "7 is a prime number"},
		{name: "not prime number", input: "8", expected: "number is not prime"},
		{name: "quit", input: "q", expected: ""},
		{name: "quit", input: "-1", expected: "negative number are not prime"},
	}

	for _, v := range test {
		intput := strings.NewReader(v.input)
		reader := bufio.NewScanner(intput)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, v.expected) {
			t.Errorf("if input %s and expected %s incorrect value returuned: %s", v.input, v.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
