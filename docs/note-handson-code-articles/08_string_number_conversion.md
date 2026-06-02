# 08 String Number Conversion: 文字列と数値は、そのままでは行き来しない

文字列と数値の変換は、実務でも本当によく出ます。フォーム入力、設定値、API のパラメータ、CSV など、文字列として受け取ったものを数値にしたい場面はかなり多いです。

このハンズオンでは、`strconv` を使って文字列と数値を相互変換します。実務でもかなりよく使う基本操作です。

最初は `strconv.Atoi` で文字列を整数にしています。ここで大事なのは、変換には失敗しうるので `err` を返すことです。Go らしいエラーハンドリングの入口でもあります。

次の `ParseFloat` は、文字列から実数への変換です。整数だけでなく、小数も文字列として受け取る場面はよくあるので、この形はかなり実用的です。

後半では `Itoa` と `FormatFloat` を使って数値から文字列へ戻しています。ここで分かるのは、「表示のために文字列へ変える」処理も明示的に書く必要があることです。

この回は、Go が型を曖昧に混ぜない言語だとよく分かる回です。数値と文字列は便利に自動変換されないからこそ、どの段階で何を扱っているかがはっきりします。

## 実行コード

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("== 文字列から整数へ ==")

	textNumber := "123"
	intValue, err := strconv.Atoi(textNumber)
	if err != nil {
		fmt.Println("Atoi error:", err)
		return
	}

	fmt.Printf("textNumber = %q (%T)\n", textNumber, textNumber)
	fmt.Printf("strconv.Atoi(textNumber) = %v (%T)\n", intValue, intValue)

	fmt.Println()
	fmt.Println("== 文字列から実数へ ==")

	textFloat := "3.14"
	floatValue, err := strconv.ParseFloat(textFloat, 64)
	if err != nil {
		fmt.Println("ParseFloat error:", err)
		return
	}

	fmt.Printf("textFloat = %q (%T)\n", textFloat, textFloat)
	fmt.Printf("strconv.ParseFloat(textFloat, 64) = %v (%T)\n", floatValue, floatValue)

	fmt.Println()
	fmt.Println("== 整数から文字列へ ==")

	score := 95
	scoreText := strconv.Itoa(score)

	fmt.Printf("score = %v (%T)\n", score, score)
	fmt.Printf("strconv.Itoa(score) = %q (%T)\n", scoreText, scoreText)

	fmt.Println()
	fmt.Println("== 実数から文字列へ ==")

	price := 980.5
	priceText := strconv.FormatFloat(price, 'f', 1, 64)

	fmt.Printf("price = %v (%T)\n", price, price)
	fmt.Printf("strconv.FormatFloat(price, 'f', 1, 64) = %q (%T)\n", priceText, priceText)

	fmt.Println()
	fmt.Println("メモ: Go では文字列と数値を直接足し算できない")
}
```

## 実行方法

```bash
go run ./examples/handson/08_string_number_conversion
```

## 実行結果の見どころ

- `"123"` が `int` へ変わるところ
- `"3.14"` が `float64` へ変わるところ
- 数値を文字列へ戻したとき、表示用のデータとして扱えるようになるところ
- 変換のたびに型がきちんと変わっているところ

## 試してほしい変更

- `"123"` を `"abc"` に変えて、`Atoi` のエラーを見る
- `FormatFloat` の桁数を変えて表示の違いを比べる
- 文字列と数値をそのまま `+` しようとしてみて、なぜできないか確認する
