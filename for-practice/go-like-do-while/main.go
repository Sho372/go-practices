package main

import (
	"fmt"
)

func main() {

	/*
		do keyword in Java, C, and JavaScripts

		do {
			// things to do in the loop
		} while (CONDITION)
	*/

	// infinite for loop that ends with an if statement
	i := 1000
	for {
		fmt.Println(i)
		if i >= 100 {
			break
		}
	}
}
