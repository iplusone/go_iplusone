package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== io.Reader / io.Writer によるストリーム処理 ===")
	demoStream()

	fmt.Println("\n=== bufio.Scanner による行単位の読み込み ===")
	demoScanner()

	fmt.Println("\n=== context によるタイムアウト ===")
	demoContext()
}

func demoStream() {
	src := strings.NewReader("Hello, streaming world!\nLine 2\nLine 3\n")
	var dst bytes.Buffer

	n, err := io.Copy(&dst, src)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Printf("転送バイト数: %d\n", n)
	fmt.Printf("内容: %s", dst.String())
}

func demoScanner() {
	text := "first line\nsecond line\nthird line"
	scanner := bufio.NewScanner(strings.NewReader(text))

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		fmt.Printf("  行 %d: %s\n", lineNum, scanner.Text())
	}
}

func demoContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	result := make(chan string, 1)

	go func() {
		// 50ms で完了する処理
		time.Sleep(50 * time.Millisecond)
		result <- "完了"
	}()

	select {
	case r := <-result:
		fmt.Println("処理結果:", r)
	case <-ctx.Done():
		fmt.Println("タイムアウト:", ctx.Err())
	}

	// タイムアウトする例
	ctx2, cancel2 := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel2()

	go func() {
		// 100ms かかる処理（タイムアウトする）
		time.Sleep(100 * time.Millisecond)
		result <- "遅い完了"
	}()

	select {
	case r := <-result:
		fmt.Println("処理結果:", r)
	case <-ctx2.Done():
		fmt.Println("タイムアウト（期待通り）:", ctx2.Err())
	}
}
