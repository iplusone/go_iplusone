package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== タイムアウト内に完了するケース ===")
	runWithTimeout(200*time.Millisecond, 50*time.Millisecond)

	fmt.Println("\n=== タイムアウトするケース ===")
	runWithTimeout(50*time.Millisecond, 200*time.Millisecond)

	fmt.Println("\n=== 手動キャンセル ===")
	runWithCancel()
}

func runWithTimeout(timeout, workDuration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan string, 1)
	go func() {
		time.Sleep(workDuration)
		done <- "処理完了"
	}()

	select {
	case result := <-done:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("タイムアウト:", ctx.Err())
	}
}

func runWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(30 * time.Millisecond)
		cancel() // 外部から明示的にキャンセル
	}()

	select {
	case <-ctx.Done():
		fmt.Println("キャンセル:", ctx.Err())
	}
}
