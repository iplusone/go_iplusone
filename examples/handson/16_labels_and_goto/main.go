package main

import "fmt"

func main() {
	fmt.Println("== goto の基本 ==")

	step := 1
	if step == 1 {
		fmt.Println("step 1 を実行")
		goto Next
	}

	fmt.Println("この行は実行されません")

Next:
	fmt.Println("Next ラベルへ移動しました")

	fmt.Println()
	fmt.Println("== ラベル付き break ==")

	found := false

Outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if i == 1 && j == 1 {
				found = true
				fmt.Println("見つかったので Outer を break します")
				break Outer
			}
		}
	}

	fmt.Printf("found = %v\n", found)

	fmt.Println()
	fmt.Println("メモ: goto は flow を追いにくくしやすいので、必要な場面だけで使う")
}
