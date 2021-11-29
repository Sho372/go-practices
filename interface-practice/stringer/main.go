package main

import (
	"fmt"
)

// Interface definition
type Stringer interface {
	String() string
}

// Implement type #1
type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

// Implement type #2
type Oct int

func (o Oct) String() string {
	return fmt.Sprintf("%x", int(o))
}

// Implement type #3
type Bin int

func (b Bin) String() string {
	return fmt.Sprintf("%b", int(b))
}

// Implement type #4
type Cp int

func (cp Cp) String() string {
	return fmt.Sprintf("%c", int(cp))
}

//Usage of Stringer interface -> Abstraction
func F(s Stringer) {

	// type switch
	switch v := s.(type) {
	case Hex:
		fmt.Println("Int:", int(v), "Hex:", v.String())
	case Oct:
		fmt.Println("Int:", int(v), "Oct:", v.String())
	case Bin:
		fmt.Println("Int:", int(v), "Bin:", v.String())
	case Cp:
		fmt.Println("Code point (Decimal):", int(v), "Character:", v.String())
	}
}

func main() {
	var i int = 100
	var h Hex = Hex(i)
	F(h)

	var b Bin = Bin(i)
	F(b)

	var o Oct = Oct(i)
	F(o)

	var c = 128512
	var cp Cp = Cp(c)
	F(cp)
}
