package main

import "fmt"

type Item struct {
	Name string
}

func main() {
	fmt.Println("== new ==")

	itemPtr := new(Item)
	fmt.Printf("itemPtr = %v (%T)\n", itemPtr, itemPtr)
	fmt.Printf("*itemPtr = %+v\n", *itemPtr)

	itemPtr.Name = "Go"
	fmt.Printf("更新後 *itemPtr = %+v\n", *itemPtr)

	fmt.Println()
	fmt.Println("== make で slice を作る ==")

	numbers := make([]int, 3, 5)
	fmt.Printf("numbers = %v (%T)\n", numbers, numbers)
	fmt.Printf("len(numbers) = %d, cap(numbers) = %d\n", len(numbers), cap(numbers))

	fmt.Println()
	fmt.Println("== make で map を作る ==")

	scores := make(map[string]int)
	scores["Go"] = 100
	fmt.Printf("scores = %v (%T)\n", scores, scores)

	fmt.Println()
	fmt.Println("== make で channel を作る ==")

	ch := make(chan string, 1)
	ch <- "hello"
	fmt.Printf("受信した値 = %q\n", <-ch)

	fmt.Println()
	fmt.Println("メモ: new はゼロ値へのポインタを返す")
	fmt.Println("メモ: make は slice / map / channel を初期化して使える状態で返す")
}
