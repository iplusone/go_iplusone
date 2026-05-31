package main

import "fmt"

func main() {
	fmt.Println("== 配列 array ==")

	numbers := [3]int{10, 20, 30}
	fmt.Printf("numbers = %v (%T)\n", numbers, numbers)
	fmt.Printf("len(numbers) = %d\n", len(numbers))

	fmt.Println()
	fmt.Println("== array を range で回す ==")

	for index, value := range numbers {
		fmt.Printf("index = %d, value = %d\n", index, value)
	}

	fmt.Println()
	fmt.Println("== slice ==")

	scores := []int{80, 90, 100}
	fmt.Printf("scores = %v (%T)\n", scores, scores)
	fmt.Printf("len(scores) = %d\n", len(scores))
	fmt.Printf("cap(scores) = %d\n", cap(scores))

	fmt.Println()
	fmt.Println("== slice に追加する ==")

	scores = append(scores, 110)
	fmt.Printf("append 後 scores = %v\n", scores)
	fmt.Printf("len(scores) = %d\n", len(scores))
	fmt.Printf("cap(scores) = %d\n", cap(scores))

	fmt.Println()
	fmt.Println("== 部分 slice を作る ==")

	part := scores[1:3]
	fmt.Printf("part = %v (%T)\n", part, part)
	fmt.Printf("len(part) = %d\n", len(part))
	fmt.Printf("cap(part) = %d\n", cap(part))

	fmt.Println()
	fmt.Println("== slice を range で回す ==")

	for index, value := range scores {
		fmt.Printf("index = %d, value = %d\n", index, value)
	}
}
