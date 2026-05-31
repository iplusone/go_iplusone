package main

import "fmt"

func main() {
	fmt.Println("== for の初期化・条件・後処理 ==")

	for i := 0; i < 3; i++ {
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println()
	fmt.Println("初期化: i := 0")
	fmt.Println("条件: i < 3")
	fmt.Println("後処理: i++")

	fmt.Println()
	fmt.Println("== 後処理を変える例 ==")

	for n := 10; n > 0; n -= 2 {
		fmt.Printf("n = %d\n", n)
	}

	fmt.Println()
	fmt.Println("== 外で宣言した変数を使う例 ==")

	count := 1
	for ; count <= 3; count++ {
		fmt.Printf("count = %d\n", count)
	}

	fmt.Printf("ループ終了後の count = %d\n", count)
}
