package main

import "fmt"

func main() {
	fmt.Println("== 値と変数 ==")

	const appName = "Go Hands-on"
	message := "Goでは値に型がある"

	var age int = 20
	score := 88

	fmt.Printf("定数 appName: %v (%T)\n", appName, appName)
	fmt.Printf("変数 message: %v (%T)\n", message, message)
	fmt.Printf("変数 age: %v (%T)\n", age, age)
	fmt.Printf("変数 score: %v (%T)\n", score, score)

	fmt.Println()
	fmt.Println("== 基本データ型 ==")

	var integer int = -42
	var unsignedInteger uint = 42
	var pointerNumber uintptr = 0
	var realNumber float64 = 3.14
	var complexNumber complex128 = 2 + 3i
	var isGoFun bool = true
	var text string = "Goはシンプルで読みやすい"

	fmt.Printf("整数 int: %v (%T)\n", integer, integer)
	fmt.Printf("符号なし整数 uint: %v (%T)\n", unsignedInteger, unsignedInteger)
	fmt.Printf("uintptr: %v (%T)\n", pointerNumber, pointerNumber)
	fmt.Printf("実数 float64: %v (%T)\n", realNumber, realNumber)
	fmt.Printf("複素数 complex128: %v (%T)\n", complexNumber, complexNumber)
	fmt.Printf("真偽値 bool: %v (%T)\n", isGoFun, isGoFun)
	fmt.Printf("文字列 string: %v (%T)\n", text, text)

	fmt.Println()
	fmt.Println("== ちょっと試す ==")
	fmt.Printf("integer + 10 = %d\n", integer+10)
	fmt.Printf("unsignedInteger * 2 = %d\n", unsignedInteger*2)
	fmt.Printf("realNumber / 2 = %.2f\n", realNumber/2)
	fmt.Printf("complexNumber * 2 = %v\n", complexNumber*2)
	fmt.Printf("!isGoFun = %v\n", !isGoFun)
	fmt.Printf("text + \"!\" = %s\n", text+"!")
}
