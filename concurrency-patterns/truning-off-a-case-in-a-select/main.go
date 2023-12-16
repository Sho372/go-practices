package main

import (
	"fmt"
)

// receiver
func main() {


	in := hello() 
	in2 := world()

	//無限ループで、channelのselectをし続ける
	//in, in2からデータが送られなくなったら、!okでin,in2をnil (channleを解放) -> これがTurn off
	//in, in2の読み取りのcaseはfor内でskipされる

	//無限ループでin, in2チャンネルをselectで監視
	for {
		select {
		case v, ok := <-in: // read ok:trueならチャンネルはクローズされておらず読み取り可能、falseの場合はすでにcloseされている
			if !ok {
				// inをcloseすると、zero value（空文字）とfalse		
				in = nil //[turn off!] the case will never succeed again!
				fmt.Printf("in [%v] ok [%v]\n", v, ok)
				continue
			}
			// process the v that was read from in
			fmt.Printf("in [%v] ok [%v]\n", v, ok)

		case v, ok := <-in2: // read 
			if !ok {
				// in2をcloseすると、zero value（空文字）とfalse		
				in2 = nil //[turn off!] the case will never succeed again!
				fmt.Printf("in2 [%v], ok [%v]\n", v, ok)
				continue
			}
			// process the v that was read from in2
			fmt.Printf("in2 [%v], ok [%v]\n", v, ok)
		default:
			//両方のchannelがcloseした時にブロックしないように必要（両channneがまだcloseしてなくても通る）
			fmt.Printf("do nothing... \n")
		}
		// 両方のチャンネルがクローズされた場合、ループを終了します。
		if in == nil && in2 == nil {
			break
		}

		//doneチャンネルを使って、明示的に2つのチャンネルがcloseしたことを伝えるには、waig groupとか必要な気がする
	}
}


// sender -> return in and in2 channel
// <-chan型は読取り専用のチャンネル
func hello() <-chan string {

	ch := make(chan string)
	go func() {
		data := []string{"H","E","L","L","O"}

		for _, v := range data {
			ch <-v //send
		}
		close(ch)
	}()
	return ch
}

func world() <-chan string {

	ch := make(chan string)
	go func() {
		data := []string{"W","O","R","L","D"}

		for _, v := range data {
			ch <-v //send
		}
		close(ch)
	}()
	return ch
}

//func main() {
//// チャンネル1とチャンネル2を作成します。
//ch1 := make(chan int)
//ch2 := make(chan int)

//// ゴルーチン1：チャンネル1から値を送信します。
//go func() {
//for i := 1; i <= 5; i++ {
//ch1 <- i
//time.Sleep(time.Second)
//}
//close(ch1) // チャンネル1をクローズします。
//}()

//// ゴルーチン2：チャンネル2から値を送信します。
//go func() {
//for i := 100; i <= 105; i++ {
//ch2 <- i
//time.Sleep(time.Second)
//}
//close(ch2) // チャンネル2をクローズします。
//}()

//// メインゴルーチン：select文を使用して2つのチャンネルから値を受信します。
//for {
//select {
//case val, ok := <-ch1:
//if !ok {
//fmt.Println("チャンネル1がクローズされました。")
//ch1 = nil // チャンネル1のケースをオフにします。
//} else {
//fmt.Println("チャンネル1から値を受信:", val)
//}
//case val, ok := <-ch2:
//if !ok {
//fmt.Println("チャンネル2がクローズされました。")
//ch2 = nil // チャンネル2のケースをオフにします。
//} else {
//fmt.Println("チャンネル2から値を受信:", val)
//}
//}

//// 両方のチャンネルがクローズされた場合、ループを終了します。
//if ch1 == nil && ch2 == nil {
//break
//}
//}
//}
