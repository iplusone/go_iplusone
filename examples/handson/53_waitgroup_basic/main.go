package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("goroutine %d 完了\n", n)
		}(i)
	}

	wg.Wait()
	fmt.Println("全 goroutine 完了")
}
