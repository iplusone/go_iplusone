package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "root:secret@tcp(localhost:3306)/tx_app?parseTime=true"
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("DB 接続失敗:", err)
	}
	fmt.Println("DB 接続成功")

	showAccounts(db)

	fmt.Println("\n--- 送金処理（Alice → Bob, 200円）---")
	if err := transfer(db, 1, 2, 200); err != nil {
		log.Println("送金失敗:", err)
	} else {
		fmt.Println("送金成功")
	}

	showAccounts(db)

	fmt.Println("\n--- 送金処理（Bob → Alice, 1000円）残高不足のケース ---")
	if err := transfer(db, 2, 1, 1000); err != nil {
		fmt.Println("送金失敗（期待通り）:", err)
	}

	showAccounts(db)
}

func showAccounts(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, balance FROM accounts ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("口座一覧:")
	for rows.Next() {
		var id, balance int
		var name string
		if err := rows.Scan(&id, &name, &balance); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  id=%d name=%s balance=%d\n", id, name, balance)
	}
}

func transfer(db *sql.DB, fromID, toID, amount int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	var fromBalance int
	if err := tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", fromID).Scan(&fromBalance); err != nil {
		tx.Rollback()
		return err
	}
	if fromBalance < amount {
		tx.Rollback()
		return fmt.Errorf("残高不足: 残高 %d, 送金額 %d", fromBalance, amount)
	}

	if _, err := tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromID); err != nil {
		tx.Rollback()
		return err
	}
	if _, err := tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
