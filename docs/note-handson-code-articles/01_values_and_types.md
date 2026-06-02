# 01 Values and Types: Goの最初のコードで「値には型がある」をつかむ

Go を触り始めたとき、最初はどうしても文法の形ばかり見てしまいがちです。`var` があって、文字列があって、数値があって、くらいの理解でも前には進めます。

でも、Go をちゃんと読みやすく感じられるようになる最初のポイントは、`値には必ず型がある` と実感することだと思っています。

このハンズオンは、Go を始めたときに最初に押さえたい `値` と `型` の感覚を確認する回です。コード自体はとても小さいですが、ここで見ていることはずっと後まで効きます。

実行すると、`appName` や `message` のような値が、値そのものと一緒に `%T` で型も表示されます。ここでまず見たいのは、「Go では値に必ず型がついている」という事実です。文字列なら `string`、整数なら `int`、真偽値なら `bool` というように、値の種類がはっきり分かれています。

後半では `int` `uint` `uintptr` `float64` `complex128` `bool` `string` を順番に出しています。ここは暗記よりも、「Go はかなり早い段階で型を意識させる言語なんだな」と感じることが大事です。

最後の `ちょっと試す` では、数値型ごとの計算や文字列連結をまとめて見せています。型が違ってもそれぞれ自然な演算がある、という入口として読むと分かりやすいです。

この回では、値と型を眺めることに集中してください。まずは `%v` と `%T` をセットで出すだけでも、Go の見え方がかなり変わります。

## 実行コード

```go
package main

import "fmt"

func main() {
	fmt.Println("== 値と変数 ==")

	const appName = "Go Hands-on"
	message := "Goでは値に型がある"

	var age int = 20
	score := 88

	fmt.Printf("定数 appName: %v (%T)\n", appName, appName)
	fmt.Printf("変数 message: %v (%T)\n", message, message)
	fmt.Printf("変数 age: %v (%T)\n", age, age)
	fmt.Printf("変数 score: %v (%T)\n", score, score)

	fmt.Println()
	fmt.Println("== 基本データ型 ==")

	var integer int = -42
	var unsignedInteger uint = 42
	var pointerNumber uintptr = 0
	var realNumber float64 = 3.14
	var complexNumber complex128 = 2 + 3i
	var isGoFun bool = true
	var text string = "Goはシンプルで読みやすい"

	fmt.Printf("整数 int: %v (%T)\n", integer, integer)
	fmt.Printf("符号なし整数 uint: %v (%T)\n", unsignedInteger, unsignedInteger)
	fmt.Printf("uintptr: %v (%T)\n", pointerNumber, pointerNumber)
	fmt.Printf("実数 float64: %v (%T)\n", realNumber, realNumber)
	fmt.Printf("複素数 complex128: %v (%T)\n", complexNumber, complexNumber)
	fmt.Printf("真偽値 bool: %v (%T)\n", isGoFun, isGoFun)
	fmt.Printf("文字列 string: %v (%T)\n", text, text)

	fmt.Println()
	fmt.Println("== ちょっと試す ==")
	fmt.Printf("integer + 10 = %d\n", integer+10)
	fmt.Printf("unsignedInteger * 2 = %d\n", unsignedInteger*2)
	fmt.Printf("realNumber / 2 = %.2f\n", realNumber/2)
	fmt.Printf("complexNumber * 2 = %v\n", complexNumber*2)
	fmt.Printf("!isGoFun = %v\n", !isGoFun)
	fmt.Printf("text + \"!\" = %s\n", text+"!")
}
```

## 実行方法

```bash
go run ./examples/handson/01_values_and_types
```

## 実行結果の見どころ

- `Go Hands-on (string)` のように、値と型がセットで表示されるところ
- `20 (int)` と `88 (int)` のように、同じ整数でも「変数として名前がついた値」として眺められるところ
- `complex128` や `uintptr` のように、普段あまり触れない型も最初から並んでいるところ
- 最後の `text + "!"` で、文字列も演算の対象になると分かるところ

## 試してほしい変更

- `age` を `20` から `20.5` に変えようとしてみて、何が起きるか確かめる
- `realNumber` を `3.14` から `3` に変えて、型表示がどうなるか見る
- `%T` を消して `%v` だけにすると、どれくらい情報が減るか比べる
