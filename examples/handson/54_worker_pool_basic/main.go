package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// ワーカー 3 本を起動
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range jobs {
				fmt.Printf("worker %d: job %d を処理\n", id, j)
			}
		}(w)
	}

	// ジョブを 10 件送信
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs) // ここで close するとワーカーの range が終了する

	wg.Wait()
	fmt.Println("全ジョブ完了")
}
