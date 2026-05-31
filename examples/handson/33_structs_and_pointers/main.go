package main

import "fmt"

type Member struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("== 構造体を値渡しする ==")

	member := Member{Name: "Ishii", Age: 20}
	fmt.Printf("呼び出し前 member = %+v\n", member)
	updateByValue(member)
	fmt.Printf("呼び出し後 member = %+v\n", member)

	fmt.Println()
	fmt.Println("== 構造体をポインタで渡す ==")

	fmt.Printf("呼び出し前 member = %+v\n", member)
	updateByPointer(&member)
	fmt.Printf("呼び出し後 member = %+v\n", member)

	fmt.Println()
	fmt.Println("== 構造体ポインタからフィールドへアクセス ==")

	ptr := &member
	ptr.Age = 30
	fmt.Printf("ptr.Age を更新後 member = %+v\n", member)
}

func updateByValue(member Member) {
	member.Age = 21
	fmt.Printf("updateByValue 内 member = %+v\n", member)
}

func updateByPointer(member *Member) {
	member.Age = 25
	fmt.Printf("updateByPointer 内 member = %+v\n", *member)
}
