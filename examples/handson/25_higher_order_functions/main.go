package main

import "fmt"

func main() {
	fmt.Println("== 関数を引数に取る高階関数 ==")

	numbers := []int{1, 2, 3, 4, 5}
	doubled := mapInts(numbers, func(value int) int {
		return value * 2
	})
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("doubled = %v\n", doubled)

	fmt.Println()
	fmt.Println("== 条件関数を渡す高階関数 ==")

	evens := filterInts(numbers, func(value int) bool {
		return value%2 == 0
	})
	fmt.Printf("evens = %v\n", evens)

	fmt.Println()
	fmt.Println("== 関数を返す高階関数 ==")

	addTen := makeAdder(10)
	addHundred := makeAdder(100)
	fmt.Printf("addTen(5) = %d\n", addTen(5))
	fmt.Printf("addHundred(5) = %d\n", addHundred(5))
}

func mapInts(values []int, transformer func(int) int) []int {
	result := make([]int, 0, len(values))
	for _, value := range values {
		result = append(result, transformer(value))
	}
	return result
}

func filterInts(values []int, predicate func(int) bool) []int {
	result := make([]int, 0, len(values))
	for _, value := range values {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

func makeAdder(base int) func(int) int {
	return func(value int) int {
		return base + value
	}
}
