package main

import "fmt"

func main() {
	fmt.Println("== if / else ==")

	age := 20
	if age >= 18 {
		fmt.Printf("age = %d なので成人です\n", age)
	} else {
		fmt.Printf("age = %d なので未成年です\n", age)
	}

	fmt.Println()
	fmt.Println("== 比較演算子 ==")

	a := 10
	b := 20

	fmt.Printf("%d == %d -> %v\n", a, b, a == b)
	fmt.Printf("%d != %d -> %v\n", a, b, a != b)
	fmt.Printf("%d < %d -> %v\n", a, b, a < b)
	fmt.Printf("%d <= %d -> %v\n", a, b, a <= b)
	fmt.Printf("%d > %d -> %v\n", a, b, a > b)
	fmt.Printf("%d >= %d -> %v\n", a, b, a >= b)

	fmt.Println()
	fmt.Println("== ショートステートメント付き if ==")

	if score := 78; score >= 80 {
		fmt.Printf("score = %d なので合格ラインを十分に超えています\n", score)
	} else if score >= 60 {
		fmt.Printf("score = %d なので合格です\n", score)
	} else {
		fmt.Printf("score = %d なので再挑戦です\n", score)
	}

	fmt.Println()
	fmt.Println("== switch / case ==")

	day := "sat"
	switch day {
	case "mon":
		fmt.Println("月曜日です")
	case "tue":
		fmt.Println("火曜日です")
	case "sat", "sun":
		fmt.Println("週末です")
	default:
		fmt.Println("平日です")
	}

	fmt.Println()
	fmt.Println("== 条件のない switch ==")

	temperature := 28
	switch {
	case temperature >= 30:
		fmt.Println("かなり暑いです")
	case temperature >= 20:
		fmt.Println("過ごしやすいです")
	case temperature >= 10:
		fmt.Println("少し涼しいです")
	default:
		fmt.Println("寒いです")
	}
}
