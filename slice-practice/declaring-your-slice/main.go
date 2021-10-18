package main

import (
	"fmt"
)

func main() {

	//#1 var
	var x []int
	fmt.Printf("x: %v length: %v, capacity: %v\n", x, len(x), cap(x))
	fmt.Println(x == nil)

	//nil is an identifer that represents the lack of a value for some types.
	x = []int{}
	fmt.Printf("x: %v length: %v, capacity: %v\n", x, len(x), cap(x))
	fmt.Println(x == nil)

	//#2 literal
	data := []int{2, 4, 6, 8}
	fmt.Printf("data: %v length: %v, capacity: %v\n", data, len(data), cap(data))

	//#3 make
	//#3-1. specify a nonzero length
	//	buf := make([]byte, 2048)
	//fmt.Printf("%v length: %v, capacity: %v\n", buf, len(buf), cap(buf))

	//#3-2. This is often done when transforming values in one slice and storing them in a second.
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5)
	fmt.Printf("s1: %v length: %v, capacity: %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v length: %v, capacity: %v\n", s2, len(s2), cap(s2))
	for i := range s2 {
		s2[i] = 2 * s1[i]
	}
	fmt.Printf("s2: %v length: %v, capacity: %v\n", s2, len(s2), cap(s2))

	//#3-3.
	//Comparing with zero lenght slice #1
	a := make([]int, 0, 5)
	var b []int
	fmt.Printf("a: %v length: %v, capacity: %v\n", a, len(a), cap(a))
	fmt.Printf("b: %v length: %v, capacity: %v\n", b, len(b), cap(b))
	for i := 1; i < 6; i++ {
		a = append(a, i)
		b = append(b, i)
	}
	fmt.Printf("a: %v length: %v, capacity: %v\n", a, len(a), cap(a))
	fmt.Printf("b: %v length: %v, capacity: %v\n", b, len(b), cap(b))

	//Comparing with non-zero length slice #3-2
	//If the number of items turns out to be smaller, you won't have an extraneous zero value at the end.
	c := make([]int, 0, 5)
	d := make([]int, 5)
	fmt.Printf("c: %v length: %v, capacity: %v\n", c, len(c), cap(c))
	fmt.Printf("d: %v length: %v, capacity: %v\n", d, len(d), cap(d))
	for i := 0; i < 3; i++ {
		c = append(c, i+1)
		d[i] = i + 1
	}
	fmt.Printf("c: %v length: %v, capacity: %v\n", c, len(c), cap(c))
	fmt.Printf("d: %v length: %v, capacity: %v\n", d, len(d), cap(d))

	e := make([]int, 0, 5)
	f := make([]int, 5)
	fmt.Printf("e: %v length: %v, capacity: %v\n", e, len(e), cap(e))
	fmt.Printf("f: %v length: %v, capacity: %v\n", f, len(f), cap(f))
	for i := 0; i < 6; i++ {
		e = append(e, i+1)
		// Panic happens. It's out of index.
		//f[i] = i + 1
	}
	fmt.Printf("e: %v length: %v, capacity: %v\n", e, len(e), cap(e))
	fmt.Printf("f: %v length: %v, capacity: %v\n", f, len(f), cap(f))
}
