package main

import "fmt"

func main() {
	/*
		nil map
	*/
	var nilMap map[string]int
	fmt.Printf("len: %v \n", len(nilMap))
	// read ok -> zero value of int
	fmt.Printf("%v \n", nilMap["apple"])
	// write ng -> panic
	//nilMap["apple"] = 10

	/*
		empty map literal
	*/
	emptyMapLiteral := map[string]int{}
	// read ok -> zero value of int
	fmt.Printf("%v \n", emptyMapLiteral["apple"])
	// write ok
	emptyMapLiteral["apple"] = 10
	fmt.Printf("%v \n", emptyMapLiteral["apple"])
	fmt.Printf("len: %v \n", len(emptyMapLiteral))

	/*
		nonempty map literal
	*/

	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Fred", "Ralph", "Bijou"},
		"Kittens": {"Fred", "Ralph", "Bijou"},
	}

	fmt.Printf("len: %v \n", len(teams))

	/*
		comma ok idiom
	*/
	timeZone := map[string]int{
		"UTC": 0 * 60 * 60,
		"EXT": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	if seconds, ok := timeZone["PST"]; ok {
		fmt.Printf("found:%v, %v", ok, seconds)
	}
}
