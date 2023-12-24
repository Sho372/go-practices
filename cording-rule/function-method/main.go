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

	fmt.Printf("%#v\n", result)

	result2, err := search2()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", result2)

	//Pattern#1 Bad
	changeName1(&result)
	fmt.Printf("%#v\n", result)

	//Pattern2 Good
	result2 = changeName2(result2)
	fmt.Printf("%#v\n", result2)
}

// Pattern#1
func search() (Result,error){

	var ret Result // zero value

	name, err := doSearch()
	if err != nil {
		return ret, err
	}

	ret.name = name
	return ret, nil

}

// Pattern#2
func search2() (Result,error){


	name, err := doSearch()
	if err != nil {
		return Result{}, err
	}

	ret := Result{
		name: name,	
	}

	return ret, nil

}

func doSearch() (string, error){
	res := "aaaa"
	return res, nil
} 

//Bad
func changeName1(r *Result) {
	r.name = r.name +"bbbb";
}


//Good
func changeName2(r Result) Result{

	ret := Result{
		name: r.name + "bbbb",
	}

	return ret
}

