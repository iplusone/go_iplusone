package main

import "fmt"

func main() {
	fmt.Println("== 短縮変数宣言の基本 ==")

	name := "Go"
	version := 1.23

	fmt.Printf("name = %q (%T)\n", name, name)
	fmt.Printf("version = %.2f (%T)\n", version, version)

	fmt.Println()
	fmt.Println("== 複数の短縮変数宣言 ==")

	firstName, lastName := "Shinichi", "Ishii"
	fmt.Printf("firstName = %q, lastName = %q\n", firstName, lastName)

	fmt.Println()
	fmt.Println("== 既存変数と新規変数を一緒に宣言 ==")

	count := 10
	fmt.Printf("更新前 count = %d\n", count)

	count, message := 20, "count を更新しつつ message を新規宣言"
	fmt.Printf("更新後 count = %d\n", count)
	fmt.Printf("message = %q\n", message)

	fmt.Println()
	fmt.Println("== 関数の中で使う ==")
	showShortDeclaration()

	fmt.Println()
	fmt.Println("メモ: := は関数の外では使えない")
	fmt.Println("メモ: := では少なくとも1つは新しい変数が必要")
}

func showShortDeclaration() {
	status := "この関数の中で宣言した変数"
	fmt.Printf("status = %q\n", status)
}
