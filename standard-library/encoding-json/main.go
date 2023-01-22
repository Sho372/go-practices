package main

import (
	"json"
	"time"
)

var (
	Order struct {
		ID          string
		DateOrdered time.Time
		CustormID   string
		Items       []Item
	}

	Item struct {
		ID   string
		Name string
	}
)

func main() {

}
