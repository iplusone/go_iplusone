package main

import (
	"fmt"
	"strings"
)

type UserName string
type Scores []int

func (u UserName) Upper() string {
	return strings.ToUpper(string(u))
}

func (u UserName) Greeting() string {
	return "こんにちは、" + string(u) + "!"
}

func (s Scores) Total() int {
	total := 0
	for _, score := range s {
		total += score
	}
	return total
}

func (s Scores) Average() float64 {
	if len(s) == 0 {
		return 0
	}
	return float64(s.Total()) / float64(len(s))
}

func main() {
	fmt.Println("== 既存の型を元に独自型を作る ==")

	name := UserName("ishii")
	fmt.Printf("name = %q (%T)\n", name, name)
	fmt.Printf("name.Upper() = %q\n", name.Upper())
	fmt.Printf("name.Greeting() = %q\n", name.Greeting())

	fmt.Println()
	fmt.Println("== slice を元に独自型を作る ==")

	scores := Scores{80, 90, 100}
	fmt.Printf("scores = %v (%T)\n", scores, scores)
	fmt.Printf("scores.Total() = %d\n", scores.Total())
	fmt.Printf("scores.Average() = %.2f\n", scores.Average())

	fmt.Println()
	fmt.Println("メモ: Go では既存型を直接変更するのではなく、新しい型を定義して振る舞いを足す")
}
