package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type messageResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func main() {
	port := env("PORT", "18081")
	origin := env("ALLOW_ORIGIN", "http://localhost:18080")

	db, err := openDB()
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
		setCORS(w, origin)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		msg, err := loadMessage(ctx, db)
		if err != nil {
			http.Error(w, "failed to load message", http.StatusInternalServerError)
			log.Printf("load message: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(msg); err != nil {
			log.Printf("encode response: %v", err)
		}
	})

	log.Printf("starting app on :%s", port)
	log.Printf("db host=%s db name=%s", env("DB_HOST", "db"), env("DB_NAME", "e2e_app"))
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func openDB() (*sql.DB, error) {
	dsn := env("DB_USER", "root") + ":" + env("DB_PASSWORD", "secret") +
		"@tcp(" + env("DB_HOST", "db") + ":" + env("DB_PORT", "3306") + ")/" +
		env("DB_NAME", "e2e_app") + "?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	pingCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func loadMessage(ctx context.Context, db *sql.DB) (messageResponse, error) {
	const q = `
SELECT id, title, message
FROM messages
ORDER BY id
LIMIT 1
`

	var msg messageResponse
	err := db.QueryRowContext(ctx, q).Scan(&msg.ID, &msg.Title, &msg.Message)
	return msg, err
}

func setCORS(w http.ResponseWriter, origin string) {
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
