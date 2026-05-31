package main

import "fmt"

func main() {
	fmt.Println("== 暗黙的型宣言 ==")

	name := "Go"
	age := 20
	height := 168.5
	isLearning := true

	fmt.Printf("name = %q (%T)\n", name, name)
	fmt.Printf("age = %d (%T)\n", age, age)
	fmt.Printf("height = %.1f (%T)\n", height, height)
	fmt.Printf("isLearning = %v (%T)\n", isLearning, isLearning)

	fmt.Println()
	fmt.Println("== 値から型が決まる ==")

	count := 10
	price := 980.0
	message := "こんにちは"

	fmt.Printf("count = %v (%T)\n", count, count)
	fmt.Printf("price = %v (%T)\n", price, price)
	fmt.Printf("message = %v (%T)\n", message, message)

	fmt.Println()
	fmt.Println("== 再代入 ==")

	score := 80
	fmt.Printf("更新前 score = %d (%T)\n", score, score)

	score = 95
	fmt.Printf("更新後 score = %d (%T)\n", score, score)

	fmt.Println()
	fmt.Println("== 複数の暗黙的型宣言 ==")

	firstName, lastName := "Shinichi", "Ishii"
	x, y := 10, 20

	fmt.Printf("firstName = %q (%T)\n", firstName, firstName)
	fmt.Printf("lastName = %q (%T)\n", lastName, lastName)
	fmt.Printf("x = %d (%T), y = %d (%T)\n", x, x, y, y)
}
