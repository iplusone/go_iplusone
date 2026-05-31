package main

import "fmt"

func main() {
	fmt.Println("== map を作る ==")

	scores := map[string]int{
		"Go":     95,
		"Python": 88,
	}

	fmt.Printf("scores = %v (%T)\n", scores, scores)
	fmt.Printf("len(scores) = %d\n", len(scores))

	fmt.Println()
	fmt.Println("== 値を追加・更新する ==")

	scores["Java"] = 82
	scores["Go"] = 100

	fmt.Printf("scores = %v\n", scores)

	fmt.Println()
	fmt.Println("== 値を取り出す ==")

	goScore := scores["Go"]
	rubyScore := scores["Ruby"]

	fmt.Printf("scores[%q] = %d\n", "Go", goScore)
	fmt.Printf("scores[%q] = %d\n", "Ruby", rubyScore)

	fmt.Println()
	fmt.Println("== 値の有無を確認する ==")

	value, ok := scores["Python"]
	fmt.Printf("scores[%q] = %d, ok = %v\n", "Python", value, ok)

	value, ok = scores["PHP"]
	fmt.Printf("scores[%q] = %d, ok = %v\n", "PHP", value, ok)

	fmt.Println()
	fmt.Println("== 値を削除する ==")

	delete(scores, "Java")
	fmt.Printf("scores = %v\n", scores)

	fmt.Println()
	fmt.Println("== range で回す ==")

	for key, value := range scores {
		fmt.Printf("key = %q, value = %d\n", key, value)
	}
}
