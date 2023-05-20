package main

import "time"

// fan-in pattern
// producer : consumer = n : 1

// producer
func producer(ch chan int, d time.Duration) {
	var i int
	for {
		ch <- i
		i++
		time.Sleep(d)
	}
}

// consumer
func reader(out chan int) {
	for {
		<-out
	}
}

func main() {
	ch, out := make(chan int), make(chan int)
	// sent data with defferent rates
	go producer(ch, 100*time.Millisecond)
	go producer(ch, 300*time.Millisecond)
	go reader(out)
	for i := range ch {
		out <- i
	}
}
