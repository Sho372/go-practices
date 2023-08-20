package main

import (
	"fmt"
)

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
		//time.Sleep(time.Second) //vの更新を遅らせれれば期待した動きをする
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
