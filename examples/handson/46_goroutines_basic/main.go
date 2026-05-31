package main

import (
	"fmt"
	"sync"
	"time"
)

func say(label string, delay time.Duration, count int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= count; i++ {
		fmt.Printf("%s %d\n", label, i)
		time.Sleep(delay)
	}
}

func main() {
	fmt.Println("== goroutine を起動する ==")

	var wg sync.WaitGroup
	wg.Add(2)

	go say("goroutine", 80*time.Millisecond, 5, &wg)
	go say("main-side", 120*time.Millisecond, 5, &wg)

	fmt.Println("2 つの goroutine を起動しました")
	fmt.Println("出力順は毎回同じとは限りません")

	wg.Wait()

	fmt.Println()
	fmt.Println("メモ: go を付けると関数は別 goroutine で実行される")
	fmt.Println("メモ: WaitGroup を使うと終了待ちを明示できる")
}
