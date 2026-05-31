package main

import (
	"fmt"
	"time"
)

func total(n int, c chan int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		time.Sleep(10 * time.Millisecond)
	}
	c <- sum
}

func main() {
	fmt.Println("== channel を作る ==")

	c := make(chan int)
	fmt.Printf("c = %T\n", c)

	fmt.Println()
	fmt.Println("== goroutine から結果を送る ==")

	go total(10, c)
	fmt.Println("main は channel からの受信待ちに入ります")

	result := <-c
	fmt.Printf("受信した値 = %d\n", result)

	fmt.Println()
	fmt.Println("メモ: <-c は値が届くまで待つ")
	fmt.Println("メモ: channel は値の受け渡しと同期の両方に使える")
}
