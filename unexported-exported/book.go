package book

import (
	"fmt"
	"person"
)

type Book struct {
	Author *person.Person
}

func (b *Book) AuthorName() string {
	return fmt.Sprint("%s %s", b.Author.GetFirstName(), b.Author.GetLastName())
}
