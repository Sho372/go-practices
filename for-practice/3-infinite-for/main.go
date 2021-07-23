package main

import(
	"fmt"
)

func main() {
	for {
		fmt.Println("Hello")
	}
//	printHello()
}

// https://www.educative.io/collection/page/6151088528949248/4547996664463360/6292303276670976
func printHello() {
	fmt.Println("Hello")
	printHello()
}
