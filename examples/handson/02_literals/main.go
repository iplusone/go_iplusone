package main

import "fmt"

func main() {
	fmt.Println("== 整数リテラル ==")

	decimal := 10
	octal := 0o12
	hexadecimal := 0xA

	fmt.Printf("10進数 10 = %v (%T)\n", decimal, decimal)
	fmt.Printf("8進数 0o12 = %v (%T)\n", octal, octal)
	fmt.Printf("16進数 0xA = %v (%T)\n", hexadecimal, hexadecimal)

	fmt.Println()
	fmt.Println("== 実数リテラル ==")

	floatNumber := 3.14
	exponentNumber := 1.2e3

	fmt.Printf("実数 3.14 = %v (%T)\n", floatNumber, floatNumber)
	fmt.Printf("指数表記 1.2e3 = %v (%T)\n", exponentNumber, exponentNumber)

	fmt.Println()
	fmt.Println("== テキストリテラル ==")

	interpreted := "Go\nHands-on"
	raw := `Go
Hands-on`

	fmt.Printf("解釈付き文字列 = %q\n", interpreted)
	fmt.Printf("生文字列 = %q\n", raw)

	fmt.Println()
	fmt.Println("== 真偽値リテラル ==")

	isGoSimple := true
	isJavaScriptTypeSafe := false

	fmt.Printf("true = %v (%T)\n", isGoSimple, isGoSimple)
	fmt.Printf("false = %v (%T)\n", isJavaScriptTypeSafe, isJavaScriptTypeSafe)
}
