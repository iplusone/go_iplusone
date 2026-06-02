# 07 Type Conversion: Goでは型変換をはっきり書く

Go を触っていて最初に少し厳しく感じるのが、型変換をかなり明示的に書かせるところかもしれません。でも、個人的にはこの厳しさのおかげで、コードの意図が見えやすくなっていると思っています。

このハンズオンは、Go の「型変換は明示的に書く」という文化を確認する回です。ここを早めに押さえておくと、後で「なんで勝手に変換してくれないのか」と悩みにくくなります。

最初は `int` と `float64` の行き来です。整数から実数へ、実数から整数へと変換していますが、どちらも明示的に `float64(...)` や `int(...)` と書いています。Go ではこのはっきりした書き方が基本です。

次の計算では、型をそろえないと計算できない例が出ます。これがあるので、変換はただの文法ではなく「計算できる状態を作るための操作」だと分かります。

文字コードの例では、数値から `rune` を経由して文字列にする流れが出ます。ここは後で文字列や Unicode を学ぶときの入口になります。

独自型 `Score` への変換も重要です。同じ元型が `int` でも、別の型として扱うなら変換が必要です。Go の型システムがかなりはっきりしていることが分かる回です。

## 実行コード

```go
package main

import "fmt"

type Score int

func main() {
	fmt.Println("== 整数と実数の型変換 ==")

	count := 12
	rate := 2.5

	floatCount := float64(count)
	intRate := int(rate)

	fmt.Printf("count = %v (%T)\n", count, count)
	fmt.Printf("float64(count) = %v (%T)\n", floatCount, floatCount)
	fmt.Printf("rate = %v (%T)\n", rate, rate)
	fmt.Printf("int(rate) = %v (%T)\n", intRate, intRate)

	fmt.Println()
	fmt.Println("== 計算で型をそろえる ==")

	total := float64(count) * rate
	fmt.Printf("float64(count) * rate = %v (%T)\n", total, total)

	fmt.Println()
	fmt.Println("== 文字コードと文字列 ==")

	codePoint := 65
	character := string(rune(codePoint))

	fmt.Printf("codePoint = %v (%T)\n", codePoint, codePoint)
	fmt.Printf("string(rune(codePoint)) = %q (%T)\n", character, character)

	fmt.Println()
	fmt.Println("== 独自型への変換 ==")

	rawScore := 95
	playerScore := Score(rawScore)

	fmt.Printf("rawScore = %v (%T)\n", rawScore, rawScore)
	fmt.Printf("Score(rawScore) = %v (%T)\n", playerScore, playerScore)

	fmt.Println()
	fmt.Println("メモ: Go では多くの場合、型変換は明示的に書く")
}
```

## 実行方法

```bash
go run ./examples/handson/07_type_conversion
```

## 実行結果の見どころ

- `float64(count)` や `int(rate)` のように、変換後の型が明確に変わっているところ
- `float64(count) * rate` のように、型をそろえると計算できるところ
- `string(rune(codePoint))` で、数値が文字として見えるようになるところ
- `Score(rawScore)` によって、同じ `int` 系でも別の型として扱えるところ

## 試してほしい変更

- `float64(count)` を外して計算しようとしてみる
- `rate := 2.5` を `rate := 2` に変えて、型と計算結果を比べる
- `codePoint := 65` を別の値に変えて、どんな文字になるか試す
