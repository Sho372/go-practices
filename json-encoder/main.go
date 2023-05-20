package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {

	type person struct {
		name string
		age  int
	}
	type persons []person

	data := []string{
		`{"name": "Fred", "age": 40}`,
		`{"name": "Mary", "age": 21}`,
		`{"name": "Fred", "age": 30}`,
	}

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.Encode(data)
	fmt.Println(b.String())
	fmt.Println(b.Bytes())

	var ps persons
	json.Unmarshal([]byte(b.String()), &ps)
	fmt.Println(ps[0].name)

}
