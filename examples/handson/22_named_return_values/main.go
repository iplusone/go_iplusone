package main

import "fmt"

func main() {
	fmt.Println("== 名前付き戻り値 ==")

	sum, diff := calc(30, 12)
	fmt.Printf("calc(%d, %d) = sum:%d diff:%d\n", 30, 12, sum, diff)

	fmt.Println()
	fmt.Println("== 名前付き戻り値に代入して return ==")

	message, ok := buildMessage("Go")
	fmt.Printf("buildMessage(%q) = %q, ok = %v\n", "Go", message, ok)

	message, ok = buildMessage("")
	fmt.Printf("buildMessage(%q) = %q, ok = %v\n", "", message, ok)
}

func calc(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func buildMessage(name string) (message string, ok bool) {
	if name == "" {
		message = "名前が空です"
		ok = false
		return
	}

	message = "こんにちは、" + name + "!"
	ok = true
	return
}
