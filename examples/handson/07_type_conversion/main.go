package main

import "fmt"

type Score int

func main() {
	fmt.Println("== 整数と実数の型変換 ==")

	count := 12
	rate := 2.5

	floatCount := float64(count)
	intRate := int(rate)

	fmt.Printf("count = %v (%T)\n", count, count)
	fmt.Printf("float64(count) = %v (%T)\n", floatCount, floatCount)
	fmt.Printf("rate = %v (%T)\n", rate, rate)
	fmt.Printf("int(rate) = %v (%T)\n", intRate, intRate)

	fmt.Println()
	fmt.Println("== 計算で型をそろえる ==")

	total := float64(count) * rate
	fmt.Printf("float64(count) * rate = %v (%T)\n", total, total)

	fmt.Println()
	fmt.Println("== 文字コードと文字列 ==")

	codePoint := 65
	character := string(rune(codePoint))

	fmt.Printf("codePoint = %v (%T)\n", codePoint, codePoint)
	fmt.Printf("string(rune(codePoint)) = %q (%T)\n", character, character)

	fmt.Println()
	fmt.Println("== 独自型への変換 ==")

	rawScore := 95
	playerScore := Score(rawScore)

	fmt.Printf("rawScore = %v (%T)\n", rawScore, rawScore)
	fmt.Printf("Score(rawScore) = %v (%T)\n", playerScore, playerScore)

	fmt.Println()
	fmt.Println("メモ: Go では多くの場合、型変換は明示的に書く")
}
