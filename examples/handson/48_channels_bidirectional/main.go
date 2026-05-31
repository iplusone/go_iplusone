package main

import "fmt"

func total(input chan int, result chan int) {
	n := <-input
	fmt.Printf("worker: n = %d を受信\n", n)

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	result <- sum
}

func main() {
	fmt.Println("== 2 本の channel を作る ==")

	input := make(chan int)
	result := make(chan int)

	go total(input, result)

	fmt.Println()
	fmt.Println("== main から worker に値を送る ==")

	input <- 100

	fmt.Println()
	fmt.Println("== worker から結果を受け取る ==")

	fmt.Printf("total = %d\n", <-result)

	fmt.Println()
	fmt.Println("メモ: input は main -> worker の通路")
	fmt.Println("メモ: result は worker -> main の通路")
}
