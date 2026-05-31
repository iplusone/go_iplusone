package main

import "fmt"

type General interface{}

func showInt(v General) {
	n, ok := v.(int)
	if !ok {
		fmt.Printf("int ではない値です: %v (%T)\n", v, v)
		return
	}

	fmt.Printf("int として取り出せた値: %d\n", n)
}

func showString(v General) {
	s, ok := v.(string)
	if !ok {
		fmt.Printf("string ではない値です: %v (%T)\n", v, v)
		return
	}

	fmt.Printf("string として取り出せた値: %s\n", s)
}

func main() {
	fmt.Println("== 型アサーションの基本 ==")

	var value General

	value = 123
	showInt(value)
	showString(value)

	fmt.Println()

	value = "hello"
	showString(value)
	showInt(value)

	fmt.Println()
	fmt.Println("value.(型) で、interface の中身を元の型として取り出す")
}
