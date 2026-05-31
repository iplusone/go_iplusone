package main

import "fmt"

func main() {
	fmt.Println("== 無限ループと break ==")

	count := 0
	for {
		fmt.Printf("count = %d\n", count)
		count++

		if count == 3 {
			fmt.Println("break でループを終了します")
			break
		}
	}

	fmt.Println()
	fmt.Println("== continue ==")

	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("i == 2 なので continue で次へ進みます")
			continue
		}

		fmt.Printf("i = %d\n", i)
	}

	fmt.Println()
	fmt.Println("== break と continue を組み合わせる ==")

	number := 0
	for {
		number++

		if number%2 == 1 {
			fmt.Printf("%d は奇数なので continue\n", number)
			continue
		}

		fmt.Printf("%d は偶数なので表示\n", number)

		if number >= 6 {
			fmt.Println("6 以上になったので break")
			break
		}
	}
}
