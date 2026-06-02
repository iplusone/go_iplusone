# 10 Imports: Goのコードは標準ライブラリを読み込んで育つ

Go を少し触っていくと、言語そのものよりも「標準ライブラリがかなり使いやすいな」と感じる瞬間が増えてきます。この回は、その入口です。

このハンズオンは、`import` の基本を見る回です。Go は標準ライブラリがかなり強いので、ここから一気に「言語だけ」から「使える道具」へ広がっていきます。

コードでは `fmt` `strings` `strconv` を読み込んでいます。`fmt` は出力、`strings` は文字列操作、`strconv` は型変換です。前のハンズオンで見てきたことが、パッケージを通して自然につながってきます。

大事なのは、`import` はただの宣言ではなく「このコードは外の機能に依存しています」と示すことです。何を使っているかが先頭で分かるので、Go のファイルはかなり読みやすいです。

最後の「使っていない import はコンパイルエラーになる」というメモも重要です。Go は不要な依存をそのままにしない設計なので、コードが散らかりにくくなっています。

この回は短いですが、Go らしい整理の美しさがよく出ています。

## 実行コード

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("== import の基本 ==")
	fmt.Println("import は別のパッケージの機能を使うために書く")

	fmt.Println()
	fmt.Println("== fmt パッケージ ==")

	name := "Go"
	fmt.Printf("name = %q\n", name)

	fmt.Println()
	fmt.Println("== strings パッケージ ==")

	message := "go hands-on"
	upper := strings.ToUpper(message)
	contains := strings.Contains(message, "hand")

	fmt.Printf("message = %q\n", message)
	fmt.Printf("strings.ToUpper(message) = %q\n", upper)
	fmt.Printf("strings.Contains(message, %q) = %v\n", "hand", contains)

	fmt.Println()
	fmt.Println("== strconv パッケージ ==")

	text := "42"
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Atoi error:", err)
		return
	}

	fmt.Printf("text = %q\n", text)
	fmt.Printf("strconv.Atoi(text) = %d\n", number)

	fmt.Println()
	fmt.Println("メモ: 使っていない import はコンパイルエラーになる")
}
```

## 実行方法

```bash
go run ./examples/handson/10_imports
```

## 実行結果の見どころ

- `fmt` `strings` `strconv` が、それぞれ役割の違う道具として自然に並んでいるところ
- 文字列の大文字変換や部分一致判定が、標準ライブラリだけで書けるところ
- `strconv.Atoi` によって、前の型変換の話がパッケージ利用へつながっているところ

## 試してほしい変更

- `strings.ToUpper` を `strings.ToLower` に変えてみる
- `strings.Contains` の検索文字列を変えて結果を比べる
- 使っていない import を1つ追加して、コンパイル時に何が起きるか確かめる
