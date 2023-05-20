package main

import (
	"fmt"
	"time"
)

func main() {

	datetime := "20230515"

	now, err := time.Parse("20060102", datetime)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 10; i++ {
		tomorrow := now.AddDate(0, 0, i)
		s := tomorrow.Format("20060102")
		fmt.Println(s)
	}

}
