package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	// wg.Add(3)
	wg.Add(4)
	go func ()  {
		defer wg.Done()
		doThing()
	}()
	go func() {
		defer wg.Done()
		doThing2()
	}()
	go func() {
		defer wg.Done()
		doThing3()
	}()
	go func(v *sync.WaitGroup) {
		defer v.Done()
		doThing4()
	}(&wg)
	wg.Wait()
	fmt.Println("main is finished")
}


func doThing() {
	time.Sleep(1 * time.Second)
	fmt.Println("#1 is finished")	
}


func doThing2() {
	time.Sleep(2 * time.Second)
	fmt.Println("#2 is finished")	
}

func doThing3() {
	time.Sleep(3 * time.Second)
	fmt.Println("#3 is finished")	
}

func doThing4() {
	time.Sleep(4 * time.Second)
	fmt.Println("#4 is finished")	
}