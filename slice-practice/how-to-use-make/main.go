package main

import "fmt"

func main() {
	x := make([]int, 5)
	fmt.Printf("%v length: %v, capacity: %v\n", x, len(x), cap(x))

	// Common beginner mistake is to try to populate those inital elements useing append.
	x = append(x, 99)
	fmt.Printf("%v length: %v, capacity: %v\n", x, len(x), cap(x))

	x = make([]int, 5, 10)
	fmt.Printf("%v length: %v, capacity: %v\n", x, len(x), cap(x))

	// You can also create a slice with zero length, but a capacity that's greater than zero
	x = make([]int, 0, 10)
	fmt.Printf("%v length: %v, capacity: %v\n", x, len(x), cap(x))

	// We can't directly index into it, but we can append values to it;
	x = append(x, 5, 6, 7, 8)
	fmt.Printf("%v length: %v, capacity: %v\n", x, len(x), cap(x))
}
