package main

import (
	"fmt"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)
	go player(table)
	go player(table)

	fmt.Printf("Start Ping Pong: Ball is %v\n", Ball)
	table <- Ball                                    // mainさんがTable channelに球出し
	time.Sleep(1 * time.Second)                      // 1秒間ラリーを続ける (main go routineをブロック)
	fmt.Printf("End Ping Pong: Ball is %v", <-table) //1秒後にmainがtable channelから読取り、main go routine終了。そして、player go routine#1,#2も終了
}

func player(table chan int) {
	for {
		ball := <-table
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
