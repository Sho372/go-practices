package main

import (
	"fmt"
)

func processChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc) // buffered channle
	for i := 0; i < conc; i++ {
		go func(v int) {
			results <- process(v)
		}(i)
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func process(i int) int {
	return i * 2
}

func main() {

	ch := make(chan int)
	out := processChannel(ch)

	for i, v := range out {
		fmt.Printf("i: [%v], v: [%v]\n", i, v)
	}

}
