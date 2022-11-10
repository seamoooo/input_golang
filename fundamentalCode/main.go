package main

import (
	"errors"
	"fmt"
	"time"
)

func timeLimit() (int, error) {
	var result int
	var err error

	done := make(chan struct{})

	go func() {
		reslut, err := doSomeWork()
		if err != nil {
			fmt.Println(reslut)
		}
		close(done)
	}()

	select {
	case <-done:
		return result, err
	case <-time.After(2 * time.Second):
		return 0, errors.New("timeout")
	}
}

func doSomeWork() (result int, err error) {
	// time.Sleep(time.Second * 10)
	return 0, nil
}

func main() {
	fmt.Println(timeLimit())
}
