package main

import (
	"fmt"
	"time"
)

func count(label string, max int, delay time.Duration, c chan string) {
	for i := 1; i <= max; i++ {
		time.Sleep(delay)
		c <- fmt.Sprintf("%s %d", label, i)
	}
}

func main() {
	fmt.Println("== 3 つの goroutine を起動する ==")

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go count("* first", 3, 110*time.Millisecond, c1)
	go count("** second", 3, 70*time.Millisecond, c2)
	go count("*** third", 3, 40*time.Millisecond, c3)

	fmt.Println("select で、値が届いた channel の処理に入ります")
	fmt.Println()

	for i := 0; i < 9; i++ {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		case msg := <-c3:
			fmt.Println(msg)
		}
	}

	fmt.Println()
	fmt.Println("メモ: select は受信可能な case を選んで実行する")
	fmt.Println("メモ: 出力順は毎回同じとは限らない")
}
