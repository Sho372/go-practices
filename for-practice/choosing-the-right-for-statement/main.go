package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		if i == 0 || i == len(evenVals)-1 {
			continue
		}
		fmt.Println(i, v)
	}

	fmt.Println()

	for i := 1; i < len(evenVals)-1; i++ {
		fmt.Println(i, evenVals[i])
	}
}
