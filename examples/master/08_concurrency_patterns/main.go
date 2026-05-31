package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("=== Mutex による排他制御 ===")
	demoMutex()

	fmt.Println("\n=== WaitGroup による完了待ち ===")
	demoWaitGroup()

	fmt.Println("\n=== Worker Pool ===")
	demoWorkerPool()

	fmt.Println("\n=== atomic による安全なカウンタ ===")
	demoAtomic()
}

func demoMutex() {
	var mu sync.Mutex
	count := 0

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("Mutex あり: count = %d (期待値: 1000)\n", count)
}

func demoWaitGroup() {
	var wg sync.WaitGroup
	results := make([]int, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			results[n] = n * n
		}(i)
	}

	wg.Wait()
	fmt.Println("結果:", results)
}

func demoWorkerPool() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// ワーカー 3 本を起動
	for w := 0; w < 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range jobs {
				fmt.Printf("  worker %d: job %d 処理\n", id, j)
			}
		}(w)
	}

	// ジョブを送信
	for j := 0; j < 9; j++ {
		jobs <- j
	}
	close(jobs) // close すると worker の range が終了する

	wg.Wait()
}

func demoAtomic() {
	var count int64

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("atomic: count = %d (期待値: 1000)\n", atomic.LoadInt64(&count))
}
