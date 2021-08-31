package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)
}
