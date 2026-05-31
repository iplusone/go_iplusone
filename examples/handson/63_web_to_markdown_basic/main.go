package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

const sampleHTML = `<!DOCTYPE html>
<html>
<head><title>Go 学習ガイド</title></head>
<body>
  <h1>Go 学習ガイド</h1>
  <p>Go は <strong>シンプル</strong> で <em>高速</em> なプログラミング言語です。</p>
  <h2>基礎文法</h2>
  <ul>
    <li>変数と型</li>
    <li>制御構文</li>
    <li>関数</li>
  </ul>
  <h2>実践トピック</h2>
  <ol>
    <li>並行処理</li>
    <li>HTTP サーバ</li>
    <li>データベース接続</li>
  </ol>
  <p>詳細は <a href="https://go.dev/doc/">公式ドキュメント</a> を参照してください。</p>
</body>
</html>`

func main() {
	// ローカルサーバを起動（実際の Web サイトへの接続を避けるため）
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, sampleHTML)
	}))
	defer ts.Close()

	fmt.Println("=== HTML を取得 ===")
	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("取得バイト数: %d\n", len(htmlBytes))

	fmt.Println("\n=== HTML → Markdown に変換 ===")
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(string(htmlBytes))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(markdown)

	fmt.Println("=== Markdown をファイルに保存 ===")
	outPath := "output.md"
	if err := os.WriteFile(outPath, []byte(markdown), 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("保存先: %s\n", outPath)
}
