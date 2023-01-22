package main

import "fmt"

func main() {

	var s *string
	fmt.Printf("s -> Type:[%T] Value:[%v] \n", s, s)
	fmt.Println(s == nil)

	var i interface{}
	fmt.Printf("i -> Type:[%T] Value:[%v] \n", i, i)
	fmt.Println(i == nil)

	i = s
	fmt.Printf("i -> Type:[%T] Value:[%v] \n", i, i)
	fmt.Println(i == nil)

}
