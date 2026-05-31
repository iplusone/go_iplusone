package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"` // 空の場合はフィールドを省略
}

func main() {
	fmt.Println("=== 構造体 → JSON ===")
	u := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Println(string(b))

	// omitempty の確認
	u2 := User{ID: 2, Name: "Bob"}
	b2, _ := json.Marshal(u2)
	fmt.Println(string(b2)) // email フィールドが省略される

	fmt.Println("\n=== JSON → 構造体 ===")
	data := `{"id":3,"name":"Charlie","email":"charlie@example.com"}`
	var u3 User
	if err := json.Unmarshal([]byte(data), &u3); err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Printf("ID=%d Name=%s Email=%s\n", u3.ID, u3.Name, u3.Email)

	fmt.Println("\n=== インデント付き出力 ===")
	pretty, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(pretty))
}
