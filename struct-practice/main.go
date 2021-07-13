package main

import "fmt"

type request struct {
	id   int
	name string
}

var (
	r1 = request{1, "BTC"}
	r2 = request{id: 2}
	r3 = request{} // initlize the request struct by each zero values.
)

func main() {
	fmt.Println(r1, r2, r3)

	if r3.id == 0 && r3.name == "" {
		fmt.Println("id and name fields are ZERO value")
	}
}
