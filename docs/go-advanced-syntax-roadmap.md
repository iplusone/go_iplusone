# Goの高度な文法ロードマップ

このドキュメントは、基礎ハンズオンの次に学ぶ `Goの高度な文法` を整理したロードマップです。

ここでいう「高度な文法」は、単に難しい構文というより、Go の設計意図や実務での使いどころまで含めて理解したい領域を指します。

## このドキュメントの位置づけ

- [examples/handson/README.md](/home/ishii/projects/go_iplusone/examples/handson/README.md)
  基礎ハンズオンの目次
- [docs/handson-learning-guide.md](/home/ishii/projects/go_iplusone/docs/handson-learning-guide.md)
  基礎ハンズオン全体の学習順
- [docs/go-concurrency-learning-guide.md](/home/ishii/projects/go_iplusone/docs/go-concurrency-learning-guide.md)
  並行処理トピックをまとめた学習ガイド
- このドキュメント
  基礎の次に進む上級トピックの整理

## 前提

次の内容を一通り押さえたあとに読む想定です。

- 値、変数、定数
- 制御構文
- スライス、map
- 関数
- 高階関数、クロージャ
- `Printf` の基本
- ポインタ
- 構造体
- メソッド
- インタフェースの基本
- 空のインタフェース、型アサーション、`reflect` の基本

## 現在の到達点

基礎ハンズオンでは、すでに次の領域まで進んでいます。

- ポインタ
- 構造体
- メソッド
- インタフェース
- `interface{}`
- 型アサーション
- `reflect.TypeOf` と `Kind`
- `reflect` を使ったスライス型判定

つまり、このロードマップは「完全な未着手一覧」ではなく、`ここまで学んだうえで次にどう広げるか` を整理する文書として扱うのが実態に合っています。

## 上級トピックの章立て案

### 1. エラー処理

Go らしさが最も出やすい領域です。

扱いたい内容:

- `error` インターフェース
- `if err != nil`
- 複数戻り値とエラー
- `fmt.Errorf`
- `errors.Is` `errors.As`
- `defer` と後処理

学習の狙い:

- 例外ではなく値としてエラーを扱う感覚をつかむ
- 関数設計とエラー返却の関係を理解する

### 2. 構造体とメソッド

データと振る舞いを整理するための中心的な文法です。

扱いたい内容:

- `struct`
- フィールドアクセス
- 構造体リテラル
- メソッド定義
- 値レシーバとポインタレシーバ

学習の狙い:

- 関数だけではなく、型に振る舞いを持たせる設計を理解する

### 3. ポインタ

Go の実務コードを読むうえで避けて通れない領域です。

扱いたい内容:

- `&` と `*`
- ポインタ型
- 関数にポインタを渡す
- 構造体とポインタ
- `nil` ポインタ

学習の狙い:

- 値渡しと参照的な操作の違いを理解する

### 4. インターフェース

Go の抽象化を支える重要な仕組みです。

扱いたい内容:

- `interface`
- メソッド集合
- 暗黙的実装
- `any`
- 型アサーション
- 型 switch

学習の狙い:

- 継承ではなく振る舞いでつなぐ Go の設計感覚を理解する

基礎編でカバー済み:

- `interface` の定義
- 構造体による実装
- interface 型としての利用
- `interface{}` の利用
- 型アサーション
- `reflect.TypeOf` による型判定の入口

この章で次に広げたい内容:

- `type switch`
- `any` と `interface{}` の関係
- 標準ライブラリにある interface の読み方
- `io.Reader` や `fmt.Stringer` のような実用 interface

### 5. 埋め込みとコンポジション

Go の「継承しない設計」を理解するための要所です。

扱いたい内容:

- 構造体の埋め込み
- メソッド昇格
- コンポジション

学習の狙い:

- 再利用を inheritance ではなく composition で組み立てる発想をつかむ

### 6. defer / panic / recover

高度というより「落とし穴になりやすい」文法です。

扱いたい内容:

- `defer`
- 実行順
- `panic`
- `recover`

学習の狙い:

- 後処理の書き方と、通常のエラー処理との役割分担を理解する

### 7. goroutine と channel

Go の特徴が最も強く出る分野です。

扱いたい内容:

- `go` 文
- `channel`
- 送信と受信
- buffered / unbuffered
- `select`
- `close`

学習の狙い:

- 並行処理をスレッドの知識だけに頼らず Go の文法として理解する

### 8. Generics

現代の Go を使うなら避けられない新しめの文法です。

扱いたい内容:

- 型パラメータ
- 制約
- ジェネリック関数
- ジェネリック型

学習の狙い:

- 再利用性と型安全性を両立する書き方を理解する

## これから広げやすい順番

現在の基礎編を踏まえて、そのまま追加しやすい順番は次のとおりです。

1. `type switch` と interface の実践
2. エラー処理
3. `defer` / `panic` / `recover`
4. 埋め込みとコンポジション
5. goroutine / channel
6. Generics

## 次の追加候補

今の `examples/handson` に続けて足すなら、次の章が自然です。

- `46_type_switch`
- `47_interface_design_patterns`
- `48_error_handling_basic`
- `49_error_wrapping`
- `50_defer_panic_recover`

## 補足

Go の高度な文法は、文法だけを暗記しても身につきにくい分野です。

大事なのは次の3点です。

- どう書くか
- なぜその設計が選ばれているか
- 実務でどこに出てくるか

そのため、この先のハンズオンは基礎編より少しだけ説明を厚くして、`構文` と `設計意図` を一緒に扱うのが向いています。
