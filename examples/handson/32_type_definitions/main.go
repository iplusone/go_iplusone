package main

import "fmt"

type UserID int
type UserName string

type Profile struct {
	ID   UserID
	Name UserName
}

func main() {
	fmt.Println("== type で新しい型を定義する ==")

	var id UserID = 1001
	var name UserName = "Ishii"

	fmt.Printf("id = %v (%T)\n", id, id)
	fmt.Printf("name = %v (%T)\n", name, name)

	fmt.Println()
	fmt.Println("== struct の中で使う ==")

	profile := Profile{
		ID:   id,
		Name: name,
	}
	fmt.Printf("profile = %+v (%T)\n", profile, profile)

	fmt.Println()
	fmt.Println("== 元の型とは別の型として扱う ==")

	rawID := 2002
	fmt.Printf("rawID = %v (%T)\n", rawID, rawID)
	fmt.Printf("UserID(rawID) = %v (%T)\n", UserID(rawID), UserID(rawID))

	fmt.Println()
	fmt.Println("== 型変換して使う ==")

	nextID := UserID(3003)
	nextName := UserName("Guest")
	fmt.Printf("nextID = %v (%T)\n", nextID, nextID)
	fmt.Printf("nextName = %v (%T)\n", nextName, nextName)
}
