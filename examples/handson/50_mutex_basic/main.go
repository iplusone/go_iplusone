package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func increment(label string, times int, delay time.Duration, counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < times; i++ {
		counter.mu.Lock()
		counter.value++
		current := counter.value
		counter.mu.Unlock()

		fmt.Printf("%s updated counter to %d\n", label, current)
		time.Sleep(delay)
	}
}

func main() {
	fmt.Println("== 共有データを Mutex で保護する ==")

	counter := &Counter{}

	var wg sync.WaitGroup
	wg.Add(2)

	go increment("worker-A", 5, 60*time.Millisecond, counter, &wg)
	go increment("worker-B", 5, 90*time.Millisecond, counter, &wg)

	wg.Wait()

	fmt.Println()
	fmt.Printf("final counter = %d\n", counter.value)
	fmt.Println("メモ: 共有データを書き換える直前に Lock し、終わったらすぐ Unlock する")
}
