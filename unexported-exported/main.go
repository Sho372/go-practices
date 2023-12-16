package main

import (
	"fmt"

	"github.com/Sho372/go-practices/unexported-exported/book"
	"github.com/Sho372/go-practices/unexported-exported/person"
)

func main() {
	var p person.Person
	var b book.Book
	
	fmt.Println(p.FirstName, b.Author)
}
