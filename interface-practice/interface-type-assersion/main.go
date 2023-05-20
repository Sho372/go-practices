package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Status       int `json:"status"`
	Responsetime string
}

func main() {

	b := []interface{}{
		`{
			"status": 0,
			"responsetime": "2019-03-19T02:15:06.055Z"
		}`,
		`{
			"status": 1,
			"responsetime": "2019-03-19T02:15:06.055Z"
		}`,
		`{
			"status": 2,
			"responsetime": "2019-03-19T02:15:06.055Z"
		}`,
		`{
			"status": 3,
			"responsetime": "2019-03-19T02:15:06.055Z"
		}`,
	}

	var c interface{}
	// assign []interface{} to interface{}
	c = b
	//OK
	fmt.Printf("%T", c)

	// -> string NG
	// -> []byte NG
	// -> []interface{} OK
	d := c.([]interface{})
	fmt.Println("%T", d)

	for _, v := range d {
		//interface{} -> string OK
		s := v.(string)
		fmt.Printf("%v\n", s)

		var r response
		json.Unmarshal([]byte(s), &r)
		fmt.Printf("status %v\n", r.Status)
	}
}
