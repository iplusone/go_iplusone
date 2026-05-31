package main

import "fmt"

func main() {
	fmt.Println("== アドレスを取る ==")

	value := 10
	ptr := &value

	fmt.Printf("value = %d\n", value)
	fmt.Printf("&value = %p\n", &value)
	fmt.Printf("ptr = %p\n", ptr)

	fmt.Println()
	fmt.Println("== ポインタ経由で値を読む ==")

	fmt.Printf("*ptr = %d\n", *ptr)

	fmt.Println()
	fmt.Println("== ポインタ経由で値を書き換える ==")

	*ptr = 20
	fmt.Printf("value = %d\n", value)
	fmt.Printf("*ptr = %d\n", *ptr)

	fmt.Println()
	fmt.Println("== 関数にポインタを渡す ==")

	updateValue(&value)
	fmt.Printf("updateValue 後の value = %d\n", value)

	fmt.Println()
	fmt.Println("== nil ポインタ ==")

	var nilPtr *int
	fmt.Printf("nilPtr == nil -> %v\n", nilPtr == nil)
}

func updateValue(ptr *int) {
	*ptr = 99
}
