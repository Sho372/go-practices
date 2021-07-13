package main

import (
	"encoding/json"
	"fmt"
)

type RequestOmittable struct {
	Field1 int    `json:"id"`
	Field2 string `json:"name,omitempty"`
}

type RequestNotOmittable struct {
	Field1 int    `json:"id"`
	Field2 string `json:"name"`
}

func main() {

	/*
	   Marshal: struct -> string
	*/

	// #1
	req := RequestOmittable{1, "BTC"}
	bs, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}
	// The fields name is changed {"id":1,"name":"BTC"}
	bs1 := bs
	fmt.Println(string(bs1))

	// #2
	req = RequestOmittable{Field1: 1}
	bs, err = json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}
	// Field2(name) is omitted {"id":1}
	bs2 := bs
	fmt.Println(string(bs2))

	req = RequestOmittable{}
	bs, err = json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}

	// #3
	bs3 := bs
	fmt.Println(string(bs3))

	/*
	   Unmarshal: string -> struct
	*/
	var reqOmittable RequestOmittable
	if err = json.Unmarshal(bs1, &reqOmittable); err != nil {
		fmt.Println(err)
	}
	fmt.Println(reqOmittable)

	var reqOmittable2 RequestOmittable
	if err = json.Unmarshal(bs2, &reqOmittable2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(reqOmittable2)
	if reqOmittable2.Field2 == "" {
		fmt.Println("Field2 is Empty string")
	}

	var reqNotOmittable RequestOmittable
	if err = json.Unmarshal(bs2, &reqNotOmittable); err != nil {
		fmt.Println(err)
	}
	fmt.Println(reqNotOmittable)
}
