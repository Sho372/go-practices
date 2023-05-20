package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() //counter--
	for {
		task, ok := <-tasks //check if channel is closed
		if !ok {
			return //close
		}
		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Println("processing task", task)
	}
}

//50個のtaskをn人で分担して処理
func pool(wg *sync.WaitGroup, n int) {
	tasks := make(chan int)

	//workerをn個起動
	for i := 0; i < n; i++ {
		go worker(tasks, wg)
	}

	//50個のタスクをworkerに送信（書込み）
	for i := 0; i < 50; i++ {
		tasks <- i
	}
	close(tasks)
}

func main() {
	var wg sync.WaitGroup
	//wait groupのcounterに36を追加 = 36人のworker
	wg.Add(36)
	go pool(&wg, 36)
	wg.Wait() //counter==0になるまで、mainのgoroutineをblock
}
