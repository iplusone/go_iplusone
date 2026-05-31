package main

import "fmt"

const appName = "Go Hands-on"
const version = 1.0

func main() {
	fmt.Println("== 定数の基本 ==")

	const message = "定数は変更しない値"
	const year int = 2026

	fmt.Printf("appName = %q (%T)\n", appName, appName)
	fmt.Printf("version = %v (%T)\n", version, version)
	fmt.Printf("message = %q (%T)\n", message, message)
	fmt.Printf("year = %v (%T)\n", year, year)

	fmt.Println()
	fmt.Println("== 複数の定数宣言 ==")

	const (
		statusOK    = 200
		statusFound = 302
		statusError = 500
	)

	fmt.Printf("statusOK = %v\n", statusOK)
	fmt.Printf("statusFound = %v\n", statusFound)
	fmt.Printf("statusError = %v\n", statusError)

	fmt.Println()
	fmt.Println("== 定数を計算に使う ==")

	const taxRate = 0.10
	price := 1000.0
	total := price + (price * taxRate)

	fmt.Printf("price = %v (%T)\n", price, price)
	fmt.Printf("taxRate = %v (%T)\n", taxRate, taxRate)
	fmt.Printf("total = %v (%T)\n", total, total)

	fmt.Println()
	fmt.Println("メモ: 定数は再代入できない")
	fmt.Println("メモ: 変わらない値には const を使うと意図が伝わりやすい")
}
