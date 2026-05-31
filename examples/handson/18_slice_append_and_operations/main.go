package main

import "fmt"

func main() {
	fmt.Println("== append で要素を追加する ==")

	items := []string{"Go", "Java"}
	fmt.Printf("追加前 items = %v\n", items)
	fmt.Printf("len = %d, cap = %d\n", len(items), cap(items))

	items = append(items, "Python")
	fmt.Printf("追加後 items = %v\n", items)
	fmt.Printf("len = %d, cap = %d\n", len(items), cap(items))

	fmt.Println()
	fmt.Println("== 複数の要素を追加する ==")

	items = append(items, "Ruby", "Rust")
	fmt.Printf("items = %v\n", items)
	fmt.Printf("len = %d, cap = %d\n", len(items), cap(items))

	fmt.Println()
	fmt.Println("== スライスの一部を取り出す ==")

	part := items[1:4]
	fmt.Printf("part = %v\n", part)
	fmt.Printf("len(part) = %d, cap(part) = %d\n", len(part), cap(part))

	fmt.Println()
	fmt.Println("== 要素を書き換える ==")

	items[0] = "Golang"
	fmt.Printf("items = %v\n", items)
	fmt.Printf("part = %v\n", part)

	fmt.Println()
	fmt.Println("== スライスを結合する ==")

	more := []string{"TypeScript", "Kotlin"}
	items = append(items, more...)
	fmt.Printf("items = %v\n", items)
	fmt.Printf("len = %d, cap = %d\n", len(items), cap(items))
}
