package main

import "fmt"

var packageLevelMessage = "パッケージレベル変数"

func main() {
	fmt.Println("== 変数宣言 ==")

	var name string
	name = "Go"

	var age int = 20
	score := 95

	fmt.Printf("name = %q (%T)\n", name, name)
	fmt.Printf("age = %d (%T)\n", age, age)
	fmt.Printf("score = %d (%T)\n", score, score)

	fmt.Println()
	fmt.Println("== 値の代入 ==")

	message := "はじめまして"
	fmt.Printf("代入前 message = %q\n", message)

	message = "Go を学習中です"
	fmt.Printf("代入後 message = %q\n", message)

	age = age + 1
	fmt.Printf("age を更新 = %d\n", age)

	fmt.Println()
	fmt.Println("== 変数の利用範囲 ==")
	fmt.Printf("packageLevelMessage = %q\n", packageLevelMessage)

	showScope()

	{
		blockMessage := "この変数はこのブロック内だけで使える"
		fmt.Printf("blockMessage = %q\n", blockMessage)
	}

	fmt.Println()
	fmt.Println("== 変数名の付け方 ==")

	userName := "Ishii"
	totalCount := 3
	isActive := true

	fmt.Printf("userName = %q\n", userName)
	fmt.Printf("totalCount = %d\n", totalCount)
	fmt.Printf("isActive = %v\n", isActive)

	fmt.Println("分かりやすい名前にすると、値の意味が読み取りやすくなる")
}

func showScope() {
	functionMessage := "この変数は showScope 関数の中だけで使える"
	fmt.Printf("functionMessage = %q\n", functionMessage)
}
