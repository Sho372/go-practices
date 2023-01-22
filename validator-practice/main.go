package main

import (
	"./model"
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	req := new(model.Request)
	//byt := []byte(`{"No":500,"Types":["a","b"],"Name":"p"}`)
	byt := []byte(`{"No":500,"Types":["a","b"]}`)
	// #1 Prepare for struct to be validate
	if err := json.Unmarshal(byt, &req); err != nil {
		panic(err)
	}

	// #2 Prepare a validate
	validate := validator.New()
	// #3 Validate the struct
	if err := validate.Struct(req); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf(`Field:[%v] Tag:[%v] Param:[%v] Actual:[%v]`, err.Field(), err.Tag(), err.Param(), err.Value())
			fmt.Println()
		}
		return
	}

	fmt.Println(req.No)
	fmt.Printf(`%T`, req.No)
}
