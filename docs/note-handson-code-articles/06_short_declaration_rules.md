# 06 Short Declaration Rules: `:=` は便利だけれど、ルールを知ると読みやすくなる

`:=` は便利ですが、便利だからこそルールが分からないまま使うと、あとで「あれ、これは書けないのか」と引っかかります。

前回の `:=` を踏まえて、この回ではそのルールを見ます。Go の短縮宣言は便利ですが、どこでも何度でも自由に使えるわけではありません。

最初の例はごく普通の短縮宣言です。そのあとに大事なのが、既存変数と新規変数を一緒に扱う書き方です。`count, message := ...` のように書けるのは Go らしい特徴ですが、このとき少なくとも1つは新しい変数が必要です。

このルールを知らないと、「なぜここは `:=` で書けて、ここは書けないのか」が分からなくなります。地味ですが、コンパイルエラーを読み解く力にもつながります。

もうひとつ重要なのは、`:=` が関数の外では使えないことです。パッケージレベルでは `var` を使う、関数の中では `:=` も使える。この使い分けがこの回のポイントです。

コード量は少ないですが、Go を自然に読めるようになるためのルールが詰まっています。

## 実行コード

```go
package main

import "fmt"

func main() {
	fmt.Println("== 短縮変数宣言の基本 ==")

	name := "Go"
	version := 1.23

	fmt.Printf("name = %q (%T)\n", name, name)
	fmt.Printf("version = %.2f (%T)\n", version, version)

	fmt.Println()
	fmt.Println("== 複数の短縮変数宣言 ==")

	firstName, lastName := "Shinichi", "Ishii"
	fmt.Printf("firstName = %q, lastName = %q\n", firstName, lastName)

	fmt.Println()
	fmt.Println("== 既存変数と新規変数を一緒に宣言 ==")

	count := 10
	fmt.Printf("更新前 count = %d\n", count)

	count, message := 20, "count を更新しつつ message を新規宣言"
	fmt.Printf("更新後 count = %d\n", count)
	fmt.Printf("message = %q\n", message)

	fmt.Println()
	fmt.Println("== 関数の中で使う ==")
	showShortDeclaration()

	fmt.Println()
	fmt.Println("メモ: := は関数の外では使えない")
	fmt.Println("メモ: := では少なくとも1つは新しい変数が必要")
}

func showShortDeclaration() {
	status := "この関数の中で宣言した変数"
	fmt.Printf("status = %q\n", status)
}
```

## 実行方法

```bash
go run ./examples/handson/06_short_declaration_rules
```

## 実行結果の見どころ

- ごく普通の短縮宣言と、既存変数を含む短縮宣言が両方出てくるところ
- `count` は更新されつつ、`message` は新規に作られているところ
- `:=` が関数の中で使われているのに対して、関数の外では使っていないところ

## 試してほしい変更

- `count, message := ...` をもう一度同じ変数だけで書こうとしてみる
- 関数の外に `name := "Go"` を置こうとしてコンパイル結果を見る
- 既存変数だけで `:=` を使えない理由を、自分の言葉で説明してみる
