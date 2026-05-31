package main

import "fmt"

func main() {
	fmt.Println("== 値を引数に渡す ==")

	score := 80
	fmt.Printf("呼び出し前 score = %d\n", score)
	updateByValue(score)
	fmt.Printf("呼び出し後 score = %d\n", score)

	fmt.Println()
	fmt.Println("== ポインタを引数に渡す ==")

	fmt.Printf("呼び出し前 score = %d\n", score)
	updateByPointer(&score)
	fmt.Printf("呼び出し後 score = %d\n", score)

	fmt.Println()
	fmt.Println("== nil チェック付きのポインタ引数 ==")

	updateIfNotNil(nil)
	updateIfNotNil(&score)
	fmt.Printf("最終的な score = %d\n", score)
}

func updateByValue(value int) {
	value = 100
	fmt.Printf("updateByValue 内の value = %d\n", value)
}

func updateByPointer(ptr *int) {
	*ptr = 100
	fmt.Printf("updateByPointer 内の *ptr = %d\n", *ptr)
}

func updateIfNotNil(ptr *int) {
	if ptr == nil {
		fmt.Println("ptr が nil なので更新しません")
		return
	}

	*ptr += 10
	fmt.Printf("updateIfNotNil 内の *ptr = %d\n", *ptr)
}
