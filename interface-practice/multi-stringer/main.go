package main

import "fmt"

//Interface
type Stringer interface {
	String() string
}

//Interface
type Stringer2 interface {
	String() string
}

type Bin int

func (b Bin) String() string {
	return fmt.Sprintf("%b", int(b))
}

func main() {

	var i int = 10
	var a Stringer
	var b Stringer2

	//type BinはStringerとStringer2インターフェースを実装している
	//Stringerの変数に代入しても、Stringer2に変数に代入しても正常に動作する
	a = Bin(i)
	fmt.Println(a.String())

	b = Bin(i)
	fmt.Println(b.String())
}
