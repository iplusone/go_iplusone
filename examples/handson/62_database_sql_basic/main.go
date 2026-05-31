package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// テーブル作成
	_, err = db.Exec(`CREATE TABLE users (
		id   INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// データ挿入（プリペアドステートメント）
	for _, name := range []string{"Alice", "Bob", "Charlie"} {
		_, err := db.Exec("INSERT INTO users (name) VALUES (?)", name)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("3件挿入しました")

	// 全件取得
	rows, err := db.Query("SELECT id, name FROM users ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("ユーザー一覧:")
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  id=%d name=%s\n", id, name)
	}

	// 1件取得
	var name string
	err = db.QueryRow("SELECT name FROM users WHERE id = ?", 2).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id=2 の名前: %s\n", name)
}
