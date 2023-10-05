package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {
	result, err := timeLimit()
	if err != nil {
		log.Fatal(err, result)
		fmt.Printf("err: [%v], result: [%v]", err, result)
	}

	fmt.Printf("result: [%v]", result)
}

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()
	select {
	case <- done:
		return result, err
	case <- time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}
}

func doSomeWork() (int, error) {

	time.Sleep(3 * time.Second) // time out
	// time.Sleep(1 * time.Second) // success

	return 1, nil
}