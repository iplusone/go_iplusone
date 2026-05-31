package main

import "fmt"

func main() {
	fmt.Println("== 複数の戻り値 ==")

	sum, diff := calc(20, 8)
	fmt.Printf("calc(%d, %d) = sum:%d diff:%d\n", 20, 8, sum, diff)

	fmt.Println()
	fmt.Println("== 戻り値の一部だけ使う ==")

	result, _ := calc(15, 5)
	fmt.Printf("sum だけ使う = %d\n", result)

	fmt.Println()
	fmt.Println("== 状態も一緒に返す ==")

	value, ok := findScore("Go")
	fmt.Printf("findScore(%q) = %d, ok = %v\n", "Go", value, ok)

	value, ok = findScore("Ruby")
	fmt.Printf("findScore(%q) = %d, ok = %v\n", "Ruby", value, ok)
}

func calc(a int, b int) (int, int) {
	return a + b, a - b
}

func findScore(name string) (int, bool) {
	scores := map[string]int{
		"Go":     100,
		"Python": 90,
	}

	score, ok := scores[name]
	return score, ok
}
