# 03 Operations: Goで値にどんな操作ができるかを一気に見る

値を置くだけだと、まだコードは静かです。そこに演算が入ると、ようやく「処理している感じ」が出てきます。

このハンズオンは、値を置くだけではなく、そこにどう演算をかけるかを見る回です。Go の演算はかなり素直なので、ここで「数値なら数値のルール、文字列なら文字列のルールがある」とつかめれば十分です。

最初は整数同士の加算、減算、乗算、除算、剰余が並びます。ここで見てほしいのは、整数同士の `/` は小数ではなく整数として処理されることです。あとで型変換が必要になる理由の入口でもあります。

次に実数同士の計算があります。こちらは小数を含んだ結果になり、`%.2f` のような出力フォーマットも合わせて見られます。

最後は文字列連結です。Go では文字列同士は `+` でつなげられますが、数値と文字列をそのまま足すことはできません。ここが、型がはっきり分かれている言語らしいところです。

この回では、演算結果そのものよりも、「値の種類によってできる操作が変わる」という感覚を持つのが大事です。

## 実行コード

```go
package main

import "fmt"

func main() {
	fmt.Println("== 数値の演算 ==")

	a := 20
	b := 6

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)

	fmt.Println()

	x := 7.5
	y := 2.0

	fmt.Printf("%.1f + %.1f = %.1f\n", x, y, x+y)
	fmt.Printf("%.1f - %.1f = %.1f\n", x, y, x-y)
	fmt.Printf("%.1f * %.1f = %.1f\n", x, y, x*y)
	fmt.Printf("%.1f / %.1f = %.2f\n", x, y, x/y)

	fmt.Println()
	fmt.Println("== 文字列の演算 ==")

	first := "Go"
	second := "Hands-on"
	space := " "

	full := first + space + second

	fmt.Printf("%q + %q + %q = %q\n", first, space, second, full)
	fmt.Printf("%q の文字数 = %d\n", full, len(full))
}
```

## 実行方法

```bash
go run ./examples/handson/03_operations
```

## 実行結果の見どころ

- `20 / 6 = 3` となって、小数が切り捨てられるところ
- 実数側では `7.5 / 2.0 = 3.75` のように小数を保ったまま計算できるところ
- 文字列は `+` で連結できるが、数値とは別のルールで動いていると分かるところ

## 試してほしい変更

- `a / b` を `float64(a) / float64(b)` に変えて結果を比べる
- `first` と `second` の間の `space` を消して表示の違いを見る
- `%d` や `%.2f` の書式を変えて、出力の見え方を試す
