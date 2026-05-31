package main

import "fmt"

func main() {
	fmt.Println("== 数値の演算 ==")

	a := 20
	b := 6

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)

	fmt.Println()

	x := 7.5
	y := 2.0

	fmt.Printf("%.1f + %.1f = %.1f\n", x, y, x+y)
	fmt.Printf("%.1f - %.1f = %.1f\n", x, y, x-y)
	fmt.Printf("%.1f * %.1f = %.1f\n", x, y, x*y)
	fmt.Printf("%.1f / %.1f = %.2f\n", x, y, x/y)

	fmt.Println()
	fmt.Println("== 文字列の演算 ==")

	first := "Go"
	second := "Hands-on"
	space := " "

	full := first + space + second

	fmt.Printf("%q + %q + %q = %q\n", first, space, second, full)
	fmt.Printf("%q の文字数 = %d\n", full, len(full))
}
