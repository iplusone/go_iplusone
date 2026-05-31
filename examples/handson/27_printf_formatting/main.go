package main

import "fmt"

func main() {
	fmt.Println("== 基本の書式指定 ==")

	age := 20
	price := 980.5
	name := "Go"

	fmt.Printf("age = %d\n", age)
	fmt.Printf("price = %.1f\n", price)
	fmt.Printf("name = %s\n", name)

	fmt.Println()
	fmt.Println("== 文字列の見え方を変える ==")

	message := "Go\nHands-on"
	fmt.Printf("%%s  : %s\n", message)
	fmt.Printf("%%q  : %q\n", message)

	fmt.Println()
	fmt.Println("== 型や値を確認する ==")

	fmt.Printf("age の値 = %v, 型 = %T\n", age, age)
	fmt.Printf("price の値 = %v, 型 = %T\n", price, price)
	fmt.Printf("name の値 = %v, 型 = %T\n", name, name)

	fmt.Println()
	fmt.Println("== 幅をそろえる ==")

	fmt.Printf("|%5d|\n", age)
	fmt.Printf("|%8.2f|\n", price)
	fmt.Printf("|%-10s|\n", name)
}
