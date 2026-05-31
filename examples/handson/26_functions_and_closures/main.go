package main

import "fmt"

func main() {
	fmt.Println("== 無名関数 ==")

	double := func(value int) int {
		return value * 2
	}
	fmt.Printf("double(5) = %d\n", double(5))

	fmt.Println()
	fmt.Println("== クロージャ ==")

	counter := makeCounter()
	fmt.Printf("counter() = %d\n", counter())
	fmt.Printf("counter() = %d\n", counter())
	fmt.Printf("counter() = %d\n", counter())

	fmt.Println()
	fmt.Println("== 外側の変数を使うクロージャ ==")

	base := 10
	addBase := func(value int) int {
		return base + value
	}
	fmt.Printf("addBase(5) = %d\n", addBase(5))

	base = 100
	fmt.Printf("base を 100 に変更した後の addBase(5) = %d\n", addBase(5))
}

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
