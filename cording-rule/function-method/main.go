package main

import "fmt"

type Result struct {
	name string
}

func main() {

	result, err := search()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v", result)
}

func search() (Result,error){

	var ret Result // zero value

	name, err := doSearch()
	if err != nil {
		return ret, err
	}

	ret.name = name
	return ret, nil

}

func doSearch() (string, error){
	return fmt.Sprint("aaaa"), nil
} 

