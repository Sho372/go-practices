package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)


func main() {
	//inのデータ送信数と起動go routineの数を一致させたほうが効率が良くなる?
	in := tick(6)
	result := processAndGather(in, amplify, 6)
	fmt.Println(result)
}

func amplify(i int) int {
	return i * 2
}


func tick(n int) <-chan int {
	in := make(chan int, n)

	// inをcloseしないと、deadlockが起きる
	defer close(in)

	for i := 0; i < n; i++ {
		time.Sleep(1*time.Second)
		in <- i
		fmt.Printf("Send %v\n", i)
	}
	return in
}

// Lunch go routines with num times.
// Each go routines receive data out of in channel and then, process it.
func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
				fmt.Printf("receive: [%v] in #%v\n",v, goid())
			}
			fmt.Println("hhhhh")
		}()
	}
	go func() {
		fmt.Println("kkkkk")
		wg.Wait()
		fmt.Println("close 'out' channel")
		close(out)
	}()
	var result []int
	for v := range out {
		fmt.Printf("v: %v\n", v)
		result = append(result, v)
	}
	return result
}

// https://gist.github.com/metafeather/3615b23097836bc36579100dac376906
func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}