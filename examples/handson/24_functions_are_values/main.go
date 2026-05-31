package main

import "fmt"

func main() {
	fmt.Println("== 関数を変数に入れる ==")

	var op func(int, int) int
	op = add
	fmt.Printf("op(10, 20) = %d\n", op(10, 20))

	fmt.Println()
	fmt.Println("== 関数を引数として渡す ==")

	result := apply(6, 3, multiply)
	fmt.Printf("apply(6, 3, multiply) = %d\n", result)

	fmt.Println()
	fmt.Println("== 関数を戻り値として返す ==")

	selected := chooseOperation("add")
	fmt.Printf("selected(8, 4) = %d\n", selected(8, 4))
}

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func apply(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

func chooseOperation(name string) func(int, int) int {
	if name == "add" {
		return add
	}
	return multiply
}
