package main

import "fmt" // No.3 file block

var status int // No.2 package block

func main() {
	fmt.Println("aa")

	switch s := "hello"; s  { //No.4 local block (if, for, and switch statement)
	case "hello":
		fmt.Println(s)
	default:
		ss := "hhhh" // No.5 local block (Each clause in a "switch" or "select" statement)
		fmt.Println(ss)
	}

	fmt.Println(status) 
	// fmt.Println(s) // undefined
	// fmt.Println(ss) // undefined
}