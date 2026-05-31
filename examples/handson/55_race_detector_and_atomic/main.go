package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("=== atomic カウンタ（安全）===")
	var atomicCount int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&atomicCount, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("atomic: count = %d (期待値: 1000)\n", atomic.LoadInt64(&atomicCount))

	fmt.Println("\n=== Mutex カウンタ（安全）===")
	var mu sync.Mutex
	mutexCount := 0
	var wg2 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			mutexCount++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Printf("mutex: count = %d (期待値: 1000)\n", mutexCount)
}
