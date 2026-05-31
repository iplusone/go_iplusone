package main

import "fmt"

func main() {
	fmt.Println("== 可変長引数の基本 ==")

	total := sum(10, 20, 30)
	fmt.Printf("sum(10, 20, 30) = %d\n", total)

	fmt.Println()
	fmt.Println("== 引数が 0 個のとき ==")

	emptyTotal := sum()
	fmt.Printf("sum() = %d\n", emptyTotal)

	fmt.Println()
	fmt.Println("== スライスを展開して渡す ==")

	numbers := []int{5, 15, 25}
	sliceTotal := sum(numbers...)
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("sum(numbers...) = %d\n", sliceTotal)

	fmt.Println()
	fmt.Println("== 可変長引数を range で処理する ==")

	showMessages("Go", "Python", "Rust")
}

func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func showMessages(messages ...string) {
	for index, message := range messages {
		fmt.Printf("index = %d, message = %q\n", index, message)
	}
}
