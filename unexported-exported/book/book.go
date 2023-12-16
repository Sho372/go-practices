package book

import (
	"fmt"

	"github.com/Sho372/go-practices/unexported-exported/person"
)

type Book struct {
	Author *person.Person
}

func (b *Book) AuthorName() string {
	return fmt.Sprint("%s %s", b.Author.GetFirstName(), b.Author.GetLastName())
}
