package main

import (
	"fmt"
)

func main() {
	samples := []string{"hello", "apple_n!"}
	for _, s := range samples {
		for i, r := range s {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue
			}
		}
		fmt.Println("Prosecced outer loop...")
	}

	fmt.Println()

	outer:
	for _, s := range samples {
		for i, r := range s {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println("Prosecced outer loop...")
	}
}
