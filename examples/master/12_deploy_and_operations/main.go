package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthzHandler)
	mux.HandleFunc("/slow", slowHandler)
	mux.HandleFunc("/", rootHandler)

	srv := &http.Server{
		Addr:    ":8012",
		Handler: mux,
	}

	// SIGTERM / SIGINT を受け取るチャネル
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Println("起動: http://localhost:8012")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// シグナルを待つ
	sig := <-quit
	log.Printf("シグナル受信: %v - シャットダウン開始", sig)

	// 最大 10 秒待ってグレースフルシャットダウン
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("シャットダウン失敗:", err)
	}
	log.Println("シャットダウン完了")
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"status":"ok"}`)
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("slow リクエスト開始 (5秒)")
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "slow 完了")
	log.Println("slow リクエスト完了")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "running")
}
