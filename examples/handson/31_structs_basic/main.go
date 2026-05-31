package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("== 構造体を定義して使う ==")

	user := User{
		Name: "Ishii",
		Age:  20,
	}

	fmt.Printf("user = %+v\n", user)

	fmt.Println()
	fmt.Println("== フィールドにアクセスする ==")

	fmt.Printf("user.Name = %q\n", user.Name)
	fmt.Printf("user.Age = %d\n", user.Age)

	fmt.Println()
	fmt.Println("== フィールドを書き換える ==")

	user.Age = 21
	fmt.Printf("更新後 user = %+v\n", user)

	fmt.Println()
	fmt.Println("== 別の初期化方法 ==")

	guest := User{}
	guest.Name = "Guest"
	guest.Age = 18
	fmt.Printf("guest = %+v\n", guest)
}
