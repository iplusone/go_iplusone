package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("== 文字列から整数へ ==")

	textNumber := "123"
	intValue, err := strconv.Atoi(textNumber)
	if err != nil {
		fmt.Println("Atoi error:", err)
		return
	}

	fmt.Printf("textNumber = %q (%T)\n", textNumber, textNumber)
	fmt.Printf("strconv.Atoi(textNumber) = %v (%T)\n", intValue, intValue)

	fmt.Println()
	fmt.Println("== 文字列から実数へ ==")

	textFloat := "3.14"
	floatValue, err := strconv.ParseFloat(textFloat, 64)
	if err != nil {
		fmt.Println("ParseFloat error:", err)
		return
	}

	fmt.Printf("textFloat = %q (%T)\n", textFloat, textFloat)
	fmt.Printf("strconv.ParseFloat(textFloat, 64) = %v (%T)\n", floatValue, floatValue)

	fmt.Println()
	fmt.Println("== 整数から文字列へ ==")

	score := 95
	scoreText := strconv.Itoa(score)

	fmt.Printf("score = %v (%T)\n", score, score)
	fmt.Printf("strconv.Itoa(score) = %q (%T)\n", scoreText, scoreText)

	fmt.Println()
	fmt.Println("== 実数から文字列へ ==")

	price := 980.5
	priceText := strconv.FormatFloat(price, 'f', 1, 64)

	fmt.Printf("price = %v (%T)\n", price, price)
	fmt.Printf("strconv.FormatFloat(price, 'f', 1, 64) = %q (%T)\n", priceText, priceText)

	fmt.Println()
	fmt.Println("メモ: Go では文字列と数値を直接足し算できない")
}
