package main

import "fmt"

func main() {
	fmt.Println("== スライスの要素を更新する ==")

	numbers := []int{10, 20, 30}
	fmt.Printf("更新前 numbers = %v\n", numbers)
	updateElement(numbers)
	fmt.Printf("更新後 numbers = %v\n", numbers)

	fmt.Println()
	fmt.Println("== append でスライス本体を更新する ==")

	values := []int{1, 2, 3}
	fmt.Printf("更新前 values = %v\n", values)
	appendByValue(values)
	fmt.Printf("appendByValue 後 values = %v\n", values)

	fmt.Println()
	fmt.Println("== ポインタでスライス本体を更新する ==")

	appendByPointer(&values)
	fmt.Printf("appendByPointer 後 values = %v\n", values)
}

func updateElement(items []int) {
	items[0] = 99
	fmt.Printf("updateElement 内 items = %v\n", items)
}

func appendByValue(items []int) {
	items = append(items, 4)
	fmt.Printf("appendByValue 内 items = %v\n", items)
}

func appendByPointer(items *[]int) {
	*items = append(*items, 4)
	fmt.Printf("appendByPointer 内 *items = %v\n", *items)
}
