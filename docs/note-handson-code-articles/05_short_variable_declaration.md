# 05 Short Variable Declaration: `:=` を使うとGoらしい書き味になる

Go のコードを見ていると、かなり早い段階で `:=` が何度も出てきます。最初は少し記号っぽく見えますが、慣れると「Goらしさ」が一番よく出る書き方のひとつです。

このハンズオンは、Go でかなり頻繁に使う `:=` の感覚をつかむ回です。Go を読んでいると本当にたくさん出てくるので、早めに慣れておくとその後が楽になります。

前半では `name := "Go"` のように、値から型が決まる宣言をまとめて見せています。ここで大事なのは、「型を書いていない」のではなく、「コンパイラが値から型を決めている」ということです。

次の再代入パートでは、一度 `:=` で作った変数にも普通の `=` で新しい値を入れられることを確認しています。`:=` は宣言であって、更新ではない。この区別を早めに持てると混乱しにくいです。

複数同時宣言も Go ではよく出てきます。戻り値を2つ受けるときなどにも自然につながるので、この時点で見ておく意味があります。

この回では、「短く書くこと」が目的ではなく、「Go では値から型が自然に決まる場面が多い」と理解できれば十分です。

## 実行コード

```go
package main

import "fmt"

func main() {
	fmt.Println("== 暗黙的型宣言 ==")

	name := "Go"
	age := 20
	height := 168.5
	isLearning := true

	fmt.Printf("name = %q (%T)\n", name, name)
	fmt.Printf("age = %d (%T)\n", age, age)
	fmt.Printf("height = %.1f (%T)\n", height, height)
	fmt.Printf("isLearning = %v (%T)\n", isLearning, isLearning)

	fmt.Println()
	fmt.Println("== 値から型が決まる ==")

	count := 10
	price := 980.0
	message := "こんにちは"

	fmt.Printf("count = %v (%T)\n", count, count)
	fmt.Printf("price = %v (%T)\n", price, price)
	fmt.Printf("message = %v (%T)\n", message, message)

	fmt.Println()
	fmt.Println("== 再代入 ==")

	score := 80
	fmt.Printf("更新前 score = %d (%T)\n", score, score)

	score = 95
	fmt.Printf("更新後 score = %d (%T)\n", score, score)

	fmt.Println()
	fmt.Println("== 複数の暗黙的型宣言 ==")

	firstName, lastName := "Shinichi", "Ishii"
	x, y := 10, 20

	fmt.Printf("firstName = %q (%T)\n", firstName, firstName)
	fmt.Printf("lastName = %q (%T)\n", lastName, lastName)
	fmt.Printf("x = %d (%T), y = %d (%T)\n", x, x, y, y)
}
```

## 実行方法

```bash
go run ./examples/handson/05_short_variable_declaration
```

## 実行結果の見どころ

- 文字列、整数、小数、真偽値それぞれで、型が自動的に決まっているところ
- `score` を再代入しても型は変わらないところ
- 複数同時宣言でも、各変数にちゃんと型がついているところ

## 試してほしい変更

- `age := 20` を `age := 20.0` に変えて型の違いを見る
- `score = "95"` のような再代入を試してコンパイル結果を確認する
- 複数同時宣言の値を入れ替えて、型推論がどう変わるか見る
