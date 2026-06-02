# 02 Literals: コードに直接書く値にも種類がある

変数の前に、まず値そのものの書き方があります。ここを意識して見ることは意外と少ないのですが、Go のコードを読むときにはかなり大事です。

このハンズオンでは、変数に入れる前の「値の書き方」そのものを見ます。Go では、コードに直接書く `10` や `"hello"` も、それぞれ意味のあるリテラルです。

前半では整数リテラルとして、10進数、8進数、16進数を並べています。同じ数を表していても、書き方が違うと「どういう前提の値か」が分かります。ここは実務で頻出というより、「Go では整数も複数の表現を持てる」と知るための部分です。

次の実数リテラルでは `3.14` と指数表記を出しています。数値の書き方が変わっても、最終的には値として扱えることを確認できます。

文字列では、解釈付き文字列と生文字列の違いが重要です。`\n` が改行として解釈されるか、そのまま文字として入るか。この差は後で SQL、JSON、複数行テキストを書くときにも効いてきます。

最後の真偽値リテラルは短いですが、`true` と `false` がただの文字列ではなく `bool` だと確認する意味があります。

この回は、Go の値が「どう存在するか」ではなく「どう書かれるか」を見る回です。地味ですが、文法の土台になります。

## 実行コード

```go
package main

import "fmt"

func main() {
	fmt.Println("== 整数リテラル ==")

	decimal := 10
	octal := 0o12
	hexadecimal := 0xA

	fmt.Printf("10進数 10 = %v (%T)\n", decimal, decimal)
	fmt.Printf("8進数 0o12 = %v (%T)\n", octal, octal)
	fmt.Printf("16進数 0xA = %v (%T)\n", hexadecimal, hexadecimal)

	fmt.Println()
	fmt.Println("== 実数リテラル ==")

	floatNumber := 3.14
	exponentNumber := 1.2e3

	fmt.Printf("実数 3.14 = %v (%T)\n", floatNumber, floatNumber)
	fmt.Printf("指数表記 1.2e3 = %v (%T)\n", exponentNumber, exponentNumber)

	fmt.Println()
	fmt.Println("== テキストリテラル ==")

	interpreted := "Go\nHands-on"
	raw := `Go
Hands-on`

	fmt.Printf("解釈付き文字列 = %q\n", interpreted)
	fmt.Printf("生文字列 = %q\n", raw)

	fmt.Println()
	fmt.Println("== 真偽値リテラル ==")

	isGoSimple := true
	isJavaScriptTypeSafe := false

	fmt.Printf("true = %v (%T)\n", isGoSimple, isGoSimple)
	fmt.Printf("false = %v (%T)\n", isJavaScriptTypeSafe, isJavaScriptTypeSafe)
}
```

## 実行方法

```bash
go run ./examples/handson/02_literals
```

## 実行結果の見どころ

- `10進数 10` `8進数 0o12` `16進数 0xA` が、最終的には同じ値として見えるところ
- `Go\nHands-on` と生文字列の表示差で、改行の扱いが目で分かるところ
- `true` と `false` にもちゃんと `bool` という型が付いているところ

## 試してほしい変更

- `0xA` を `0xFF` に変えて、表示結果を確認する
- 解釈付き文字列の `\n` を `\t` に変えてみる
- 生文字列の中に `"` を入れてもそのまま書けるか試す
