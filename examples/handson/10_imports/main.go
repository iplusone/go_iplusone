package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("== import の基本 ==")
	fmt.Println("import は別のパッケージの機能を使うために書く")

	fmt.Println()
	fmt.Println("== fmt パッケージ ==")

	name := "Go"
	fmt.Printf("name = %q\n", name)

	fmt.Println()
	fmt.Println("== strings パッケージ ==")

	message := "go hands-on"
	upper := strings.ToUpper(message)
	contains := strings.Contains(message, "hand")

	fmt.Printf("message = %q\n", message)
	fmt.Printf("strings.ToUpper(message) = %q\n", upper)
	fmt.Printf("strings.Contains(message, %q) = %v\n", "hand", contains)

	fmt.Println()
	fmt.Println("== strconv パッケージ ==")

	text := "42"
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Atoi error:", err)
		return
	}

	fmt.Printf("text = %q\n", text)
	fmt.Printf("strconv.Atoi(text) = %d\n", number)

	fmt.Println()
	fmt.Println("メモ: 使っていない import はコンパイルエラーになる")
}
