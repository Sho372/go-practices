package main

import (
	"fmt"
	"time"
)

// sender
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 1; i <= 15; i++ {
		c <- x // write to channel (send)
		fmt.Printf("#%v send %v, len %v/%v\n", i, x, len(c), cap(c))
		//		time.Sleep(1 * time.Second)
		x, y = y, x+y
	}
	close(c)
}

// receiver
func main() {
	c := make(chan int, 10) // buffered channel
	go fibonacci(cap(c), c)
	time.Sleep(1 * time.Second)
	cnt := 1
	for i := range c { // The loop for i := range c receives values from the channel repeatedly until it is closed.
		fmt.Printf("#%v receive %v, len %v\n", cnt, i, len(c))
		time.Sleep(1 * time.Second)
		cnt += 1
	}
}
