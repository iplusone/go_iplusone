# Goハンズオン学習ガイド

このドキュメントは、`examples/handson` 配下に追加してきた Go 基礎ハンズオンを、学習順、目的、到達点の観点で整理した全体ガイドです。

各ハンズオンの `README.md` は個別回の説明、ここでは `どの順で何を学ぶか` と `今後どう増やしていくか` を扱います。

## 対象

- Go をこれから学び始める人
- Go の基本文法を手を動かして確認したい人
- ハンズオンを今後追加する時に、どこへ何を置くべきか確認したい人

## 使い方

1. まず [examples/handson/README.md](/home/ishii/projects/go_iplusone/examples/handson/README.md) で全体の目次を見る
2. 先頭の章から順番に `go run ./examples/handson/<directory>` で実行する
3. 出力を見ながら `main.go` を読み、値や条件を自分で変更して挙動を確かめる
4. つまずいたら各回の `README.md` と、このガイドの「学習の流れ」を見直す

## ディレクトリ構成

- `examples/handson/README.md`
  ハンズオン全体の目次
- `examples/handson/01_values_and_types` から `51_file_access_basic`
  各回の最小サンプル
- 各ディレクトリ内の `main.go`
  実行対象のコード
- 各ディレクトリ内の `README.md`
  その回で扱う内容と実行方法

## 学習の流れ

### 1. 値・型・リテラルを押さえる

最初の段階では、Go のプログラムが `どんな値を扱うか` を理解します。

- `01_values_and_types`
  値、変数、基本データ型の入口
- `02_literals`
  数値や文字列をコード上でどう表すか
- `03_operations`
  値に対してどんな演算ができるか

この段階の到達目標:

- `int` `float64` `bool` `string` の違いが分かる
- リテラルと変数の違いが説明できる
- 数値演算と文字列連結の基本が分かる

### 2. 変数と定数を整理する

次に、名前を付けて値を扱う方法を整理します。

- `04_variables`
  変数宣言、代入、スコープ、命名
- `05_short_variable_declaration`
  `:=` による暗黙的型宣言
- `06_short_declaration_rules`
  短縮変数宣言の使いどころとルール
- `09_constants`
  `const` と変更しない値

この段階の到達目標:

- `var` と `:=` の違いが分かる
- スコープの基本が説明できる
- 変数と定数を使い分けられる

### 3. 型変換を理解する

Go は暗黙の型変換が少ないため、ここは早めに押さえておくと後が楽になります。

- `07_type_conversion`
  数値型どうし、独自型、文字コードの変換
- `08_string_number_conversion`
  `strconv` を使った文字列と数値の変換

この段階の到達目標:

- Go では型変換を明示的に書く場面が多いと理解する
- `strconv.Atoi` `strconv.Itoa` などの基本が使える

### 4. パッケージと標準ライブラリに触れる

Go のコードは 1 ファイルだけで閉じず、パッケージを読み込んで機能を使います。

- `10_imports`
  `import` の基本と標準ライブラリ利用

この段階の到達目標:

- `import` が何のためにあるか説明できる
- `fmt` `strings` `strconv` の基本利用ができる

### 5. 制御構文で分岐を学ぶ

最後に、処理の流れを変える構文を確認します。

- `11_control_flow`
  `if` `else` 比較演算子 ショートステートメント付き `if` `switch`
- `12_switch_default_fallthrough`
  `default` と `fallthrough`

この段階の到達目標:

- 条件分岐を `if` と `switch` で書き分けられる
- 条件のない `switch` の意味が分かる
- `fallthrough` の挙動を誤解せず使える

### 6. 繰り返し処理を学ぶ

Go の繰り返し構文は `for` に集約されています。基本形と、初期化・後処理の見方をここで押さえます。

- `13_for_statement`
  基本の `for`、条件だけの `for`、`range`
- `14_for_init_and_post`
  `for` の初期化、条件、後処理の見方
- `15_infinite_loop_break_continue`
  無限ループ、`break`、`continue`
- `16_labels_and_goto`
  ラベル、`goto`、ラベル付き `break`

この段階の到達目標:

- `for init; condition; post` の3要素を説明できる
- 後処理を `i++` 以外にも変えられると分かる
- ループ変数のスコープを意識できる
- `break` と `continue` の違いを説明できる
- ラベル付き `break` の役割を理解できる

### 7. 複雑な値を学ぶ

Go では複数の値をまとめて扱うために、配列やスライスを使います。まずは固定長の配列と、柔軟に扱えるスライスの違いを押さえます。

- `17_arrays_and_slices`
  配列、`range`、スライス、`len`、`cap`
- `18_slice_append_and_operations`
  `append`、部分取得、更新、結合
- `19_maps`
  `map` の作成、参照、更新、削除

この段階の到達目標:

- 配列とスライスの見た目と型の違いが分かる
- `range` で要素を取り出せる
- `len` と `cap` の役割を説明できる
- `append` とスライス共有の基本挙動が分かる
- `map[key]` と `value, ok := map[key]` の違いが分かる

### 8. 関数を学ぶ

Go の処理は関数に切り出して再利用します。まずは定義、引数、戻り値、シグネチャの基本を押さえます。

- `20_functions_basic`
  関数定義、基本形、シグネチャ
- `21_multiple_return_values`
  複数の戻り値
- `22_named_return_values`
  名前付き戻り値
- `23_variadic_functions`
  可変長引数
- `24_functions_are_values`
  関数を値として扱う
- `25_higher_order_functions`
  高階関数
- `26_functions_and_closures`
  関数とクロージャ
- `27_printf_formatting`
  `Printf` の書式指定
- `28_pointers_basic`
  ポインタ
- `29_pointer_arguments`
  ポインタを引数に使う
- `30_pointer_and_slice`
  ポインタでスライスを操作する
- `31_structs_basic`
  構造体
- `32_type_definitions`
  `type` で独自型を定義する
- `33_structs_and_pointers`
  構造体とポインタ渡し
- `34_new_and_make`
  `new` と `make`
- `35_methods_basic`
  メソッドを定義して使う
- `36_extending_types`
  既存型を元に振る舞いを拡張する
- `37_pointer_receivers`
  ポインタレシーバ
- `38_interfaces_basic`
  インタフェースの定義、実装、利用
- `39_interfaces_with_structs`
  構造体を interface 型として扱う
- `40_interface_collections`
  interface の配列で複数の構造体をまとめる
- `41_nil_receivers`
  nil レシーバ
- `42_empty_interface`
  空のインタフェース
- `43_type_assertions`
  型アサーション
- `44_reflect_type_checks`
  `reflect.TypeOf` と `Kind` で型を判定する
- `45_reflect_slice_checks`
  `reflect` でスライス型を判定する
- `46_goroutines_basic`
  `goroutine` の基本
- `47_channels_basic`
  `channel` の基本的な送受信
- `48_channels_bidirectional`
  2 本の `channel` を使った双方向通信
- `49_select_basic`
  `select` による複数 channel の扱い
- `50_mutex_basic`
  `Mutex` による共有データ保護
- `51_file_access_basic`
  データアクセスの入口として、ファイルを読む、書く、追記する、閉じる

この段階の到達目標:

- `func` を使った関数定義が読める
- 引数と戻り値の基本形が分かる
- シグネチャが `引数型` と `戻り値型` の組み合わせだと説明できる
- 複数の戻り値を受け取り、必要に応じて `_` を使える
- 名前付き戻り値に値を入れて `return` できる
- `...` を使った可変長引数とスライス展開が分かる
- 関数を変数、引数、戻り値として扱える
- 高階関数で処理を差し替える考え方が分かる
- クロージャが外側の変数を保持することが分かる
- `Printf` の基本的な書式指定を使い分けられる
- ポインタの基本操作を説明できる
- 値渡しとポインタ渡しの違いを説明できる
- スライス要素更新とスライス本体更新の違いを説明できる
- 構造体で複数の値をひとまとめに扱える
- `type` で定義した型が元の型と区別されることを理解する
- 構造体の値渡しとポインタ渡しの違いを説明できる
- `new` と `make` の役割の違いを説明できる
- 値レシーバとポインタレシーバの基本を説明できる
- 既存型を元に新しい型を定義してメソッドを追加できる
- ポインタレシーバで元の値を更新できる
- インタフェースがメソッド集合で表されることを理解する
- interface 型では定義された機能だけ見えることを理解する
- 同じ interface を実装した複数の構造体をまとめて扱える
- nil レシーバを安全に扱うための基本が分かる
- 空のインタフェースが何でも保持できることを理解する
- 型アサーションで元の型を取り出す基本が分かる
- `reflect.TypeOf` と `Kind` で基本型を判定できる
- `reflect.SliceOf` を使ってスライス型を比較できる
- `goroutine` が並行に進むことを説明できる
- `channel` の送信と受信の基本を書ける
- `select` で複数 channel を扱える
- `Mutex` の基本的な役割を説明できる
- ファイル、ネットワーク、HTTP、JSON、DB というデータアクセスの基本的な入口を説明できる

### 9. データアクセスと周辺技術を学ぶ

Chapter 5 に相当する学習として、ファイル、ネットワーク、I/O、HTTP、JSON、DB へ順に進みます。

- `51_file_access_basic`
  `os.ReadFile` `os.WriteFile` `os.OpenFile` `defer Close`
- `52_network_access_basic`
  `http.Get` `Response.Body` `Close`
- `56_io_reader_writer_basic`
  `io.Reader` `io.Writer`
- `57_context_timeout_basic`
  `context.WithTimeout`
- `60_http_server_basic`
  `net/http`
- `61_json_api_basic`
  `encoding/json`
- `62_database_sql_basic`
  `database/sql`

この段階の到達目標:

- ファイルを読む、書く、追記する流れが分かる
- `[]byte` と `string` の変換が分かる
- `Response.Body` やファイルのような外部リソースを `defer` で閉じる理由を説明できる
- `io.Reader` / `io.Writer` という抽象化の役割が分かる
- JSON を Go の値へ変換する基本を説明できる
- `database/sql` で接続し、クエリを投げる最初の形が分かる

## 現時点のカリキュラム一覧

| No | Directory | 主題 |
| --- | --- | --- |
| 01 | `01_values_and_types` | 値と変数、基本データ型 |
| 02 | `02_literals` | 整数・実数・文字列・真偽値リテラル |
| 03 | `03_operations` | 数値の演算、文字列の演算 |
| 04 | `04_variables` | 変数宣言、代入、スコープ、命名 |
| 05 | `05_short_variable_declaration` | 暗黙的型宣言 |
| 06 | `06_short_declaration_rules` | 短縮変数宣言のルール |
| 07 | `07_type_conversion` | 値のキャスト、型変換 |
| 08 | `08_string_number_conversion` | 文字列と数値の型変換 |
| 09 | `09_constants` | 定数 |
| 10 | `10_imports` | `import` の基本 |
| 11 | `11_control_flow` | `if` `else` `switch` |
| 12 | `12_switch_default_fallthrough` | `default` `fallthrough` |
| 13 | `13_for_statement` | `for` ステートメント |
| 14 | `14_for_init_and_post` | `for` の初期化と後処理 |
| 15 | `15_infinite_loop_break_continue` | 無限ループ、`break`、`continue` |
| 16 | `16_labels_and_goto` | ラベル、`goto`、ラベル付き `break` |
| 17 | `17_arrays_and_slices` | 配列、`range`、スライス、長さと容量 |
| 18 | `18_slice_append_and_operations` | `append` とスライスの基本操作 |
| 19 | `19_maps` | `map` の基本操作 |
| 20 | `20_functions_basic` | 関数定義、基本形、シグネチャ |
| 21 | `21_multiple_return_values` | 複数の戻り値 |
| 22 | `22_named_return_values` | 名前付き戻り値 |
| 23 | `23_variadic_functions` | 可変長引数 |
| 24 | `24_functions_are_values` | 関数は値である |
| 25 | `25_higher_order_functions` | 高階関数 |
| 26 | `26_functions_and_closures` | 関数とクロージャ |
| 27 | `27_printf_formatting` | `Printf` のフォーマット出力 |
| 28 | `28_pointers_basic` | ポインタ |
| 29 | `29_pointer_arguments` | ポインタを引数に使う |
| 30 | `30_pointer_and_slice` | ポインタでスライスを操作する |
| 31 | `31_structs_basic` | 構造体 |
| 32 | `32_type_definitions` | `type` で型として定義する |
| 33 | `33_structs_and_pointers` | 構造体と参照渡し |
| 34 | `34_new_and_make` | `new` と `make` |
| 35 | `35_methods_basic` | メソッドをつかう |
| 36 | `36_extending_types` | 既存の型を拡張する |
| 37 | `37_pointer_receivers` | ポインタによるレシーバ |
| 38 | `38_interfaces_basic` | インタフェースの定義、実装、利用 |
| 39 | `39_interfaces_with_structs` | 構造体とインタフェースを組み合わせる |
| 40 | `40_interface_collections` | interface の配列で複数の構造体をまとめる |
| 41 | `41_nil_receivers` | nil レシーバ |
| 42 | `42_empty_interface` | 空のインタフェース |
| 43 | `43_type_assertions` | 型アサーション |
| 44 | `44_reflect_type_checks` | `reflect.TypeOf` と `Kind` で型を判定する |
| 45 | `45_reflect_slice_checks` | `reflect` でスライス型を判定する |
| 46 | `46_goroutines_basic` | `goroutine` の基本 |
| 47 | `47_channels_basic` | `channel` の基本的な送受信 |
| 48 | `48_channels_bidirectional` | 2 本の `channel` を使った双方向通信 |
| 49 | `49_select_basic` | `select` による複数 channel の扱い |
| 50 | `50_mutex_basic` | `Mutex` による共有データ保護 |
| 51 | `51_file_access_basic` | `ReadFile` `WriteFile` `OpenFile` によるファイルアクセス |
| 52 | `52_network_access_basic` | `http.Get` と `Response.Body` を使ったネットワークアクセスの基本 |
| 53 | `53_waitgroup_basic` | `sync.WaitGroup` による待ち合わせ |
| 54 | `54_worker_pool_basic` | worker pool パターンの入口 |
| 55 | `55_race_detector_and_atomic` | race detector と `sync/atomic` の入口 |
| 56 | `56_io_reader_writer_basic` | `io.Reader` `io.Writer` の基本 |
| 57 | `57_context_timeout_basic` | `context.WithTimeout` とキャンセル |
| 58 | `58_testing_table_driven` | テーブル駆動テスト |
| 59 | `59_benchmark_and_alloc` | ベンチマークとアロケーションの観察 |
| 60 | `60_http_server_basic` | `net/http` サーバーの基本 |
| 61 | `61_json_api_basic` | JSON API の基礎 |
| 62 | `62_database_sql_basic` | `database/sql` によるデータベースアクセスの基礎 |

## ハンズオン追加ルール

今後ハンズオンを増やす時は、次のルールにそろえると一覧性を保ちやすくなります。

- ディレクトリ名は `連番_テーマ名` 形式にする
- 各回に `main.go` と `README.md` を置く
- `main.go` は単体で `go run` できるようにする
- `README.md` には最低限 `扱う内容` と `実行方法` を書く
- 新しい回を追加したら `examples/handson/README.md` とこのガイドも更新する

## 次に追加しやすいテーマ

ここまでの流れの次に追加しやすい候補は次のとおりです。

- `goquery` を使った HTML 解析
- JSON を構造体へ `Unmarshal` する詳細版
- `QueryRow` と `sql.ErrNoRows`
- DB の CRUD を分割した発展編
- transaction の基本
- プレースホルダと安全な SQL 実行

## 関連ドキュメント

- ハンズオン目次: [examples/handson/README.md](/home/ishii/projects/go_iplusone/examples/handson/README.md)
- GoSpec 学習ガイド: [docs/gospec-learning-guide.md](/home/ishii/projects/go_iplusone/docs/gospec-learning-guide.md)
- 高度な文法ロードマップ: [docs/go-advanced-syntax-roadmap.md](/home/ishii/projects/go_iplusone/docs/go-advanced-syntax-roadmap.md)
