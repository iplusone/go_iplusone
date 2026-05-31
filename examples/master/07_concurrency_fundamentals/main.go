package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== goroutine の基本 ===")
	demoGoroutine()

	fmt.Println("\n=== unbuffered channel による同期 ===")
	demoUnbufferedChannel()

	fmt.Println("\n=== buffered channel ===")
	demoBufferedChannel()

	fmt.Println("\n=== select による多重化 ===")
	demoSelect()

	fmt.Println("\n=== channel を range で受信 ===")
	demoRangeChannel()
}

func demoGoroutine() {
	done := make(chan struct{})

	go func() {
		fmt.Println("goroutine 実行中")
		close(done)
	}()

	<-done
	fmt.Println("goroutine 完了を確認")
}

func demoUnbufferedChannel() {
	ch := make(chan int)

	go func() {
		fmt.Println("送信前")
		ch <- 42 // 受信側が準備できるまでブロックする
		fmt.Println("送信後")
	}()

	time.Sleep(50 * time.Millisecond) // 送信側が送信待ちになる時間
	fmt.Println("受信前")
	v := <-ch
	fmt.Printf("受信: %d\n", v)
}

func demoBufferedChannel() {
	ch := make(chan int, 3) // バッファ 3 → 受信側がいなくても 3 件送れる

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Printf("バッファに入れた件数: %d\n", len(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func demoSelect() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		time.Sleep(30 * time.Millisecond)
		ch1 <- "ch1 から"
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch2 <- "ch2 から"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("受信:", msg)
		case msg := <-ch2:
			fmt.Println("受信:", msg)
		}
	}
}

func demoRangeChannel() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i * i
	}
	close(ch) // close しないと range がブロックし続ける

	for v := range ch {
		fmt.Printf("  %d\n", v)
	}
}
