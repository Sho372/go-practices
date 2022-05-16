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
	original := []int{1, 2, 3, 4, 5}
	double := make([]int, 5)
	fmt.Printf("original: %v length: %v, capacity: %v\n", original, len(original), cap(original))
	fmt.Printf("double: %v length: %v, capacity: %v\n", double, len(double), cap(double))
	for i, v := range original {
		double[i] = 2 * v 
	}
	fmt.Printf("double: %v length: %v, capacity: %v\n", double, len(double), cap(double))

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

	var g []int
	h := make([]int, 0)
	var j = []int{}
	fmt.Printf("g: %v length: %v, capacity: %v\n", g, len(g), cap(g))
	fmt.Printf("h: %v length: %v, capacity: %v\n", h, len(h), cap(h))
	fmt.Printf("j: %v length: %v, capacity: %v\n", j, len(j), cap(j))
	fmt.Println(g == nil)
	fmt.Println(h == nil)
	fmt.Println(j == nil)
}
