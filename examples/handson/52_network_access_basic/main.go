package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	fmt.Println("== ネットワークアクセスの基本 ==")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintln(w, "<!DOCTYPE html>")
		fmt.Fprintln(w, "<html><body><h1>Hello from test server</h1><p>Network access basic.</p></body></html>")
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		fmt.Println("http.Get error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body error:", err)
		return
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Content-Length: %d\n", len(body))
	fmt.Println()
	fmt.Println("== Body ==")
	fmt.Println(string(body))
	fmt.Println()
	fmt.Println("メモ: http.Get で Response を取得する")
	fmt.Println("メモ: Response.Body は io.ReadCloser なので defer で Close する")
	fmt.Println("メモ: 取得した本文は []byte で受け取り、必要に応じて string に変換する")
}
