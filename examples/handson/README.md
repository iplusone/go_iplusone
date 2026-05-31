# Handson Index

Go の基礎文法と言語機能を順番に固めるためのハンズオン目次です。

全体ガイドは [docs/handson-learning-guide.md](/home/ishii/projects/go_iplusone/docs/handson-learning-guide.md) にまとめています。

このディレクトリだけで学習を完結させるのではなく、このあとに `examples/master` へ進んで Docker、Web、DB、計測、運用までつなげます。

つまり学習順は次で固定します。

1. `examples/handson`
   Go の文法、型、関数、ポインタ、構造体、interface、並行処理の基礎
2. `examples/master`
   その知識を実際のシステム実装へ接続する実践編

## 目次

1. [01 Values and Types](./01_values_and_types/README.md)
   値と変数、`int`、`uint`、`uintptr`、`float64`、`complex128`、`bool`、`string`
2. [02 Literals](./02_literals/README.md)
   整数リテラル、実数リテラル、テキストリテラル、真偽値リテラル
3. [03 Operations](./03_operations/README.md)
   数値の演算、文字列の演算
4. [04 Variables](./04_variables/README.md)
   変数宣言、値の代入、変数の利用範囲、変数名の付け方
5. [05 Short Variable Declaration](./05_short_variable_declaration/README.md)
   暗黙的型宣言
6. [06 Short Declaration Rules](./06_short_declaration_rules/README.md)
   短縮変数宣言のルール
7. [07 Type Conversion](./07_type_conversion/README.md)
   値のキャスト、型変換
8. [08 String Number Conversion](./08_string_number_conversion/README.md)
   文字列と数値の型変換
9. [09 Constants](./09_constants/README.md)
   定数
10. [10 Imports](./10_imports/README.md)
    `import` の基本
11. [11 Control Flow](./11_control_flow/README.md)
    `if`、`else`、比較演算子、ショートステートメント付き `if`、`switch`
12. [12 Switch Default Fallthrough](./12_switch_default_fallthrough/README.md)
    `switch` の `default` と `fallthrough`
13. [13 For Statement](./13_for_statement/README.md)
    `for` ステートメントの基本
14. [14 For Init And Post](./14_for_init_and_post/README.md)
    `for` の初期化、条件、後処理
15. [15 Infinite Loop Break Continue](./15_infinite_loop_break_continue/README.md)
    無限ループ、`break`、`continue`
16. [16 Labels And Goto](./16_labels_and_goto/README.md)
    ラベル、`goto`、ラベル付き `break`
17. [17 Arrays And Slices](./17_arrays_and_slices/README.md)
    配列、`range`、スライス、長さと容量
18. [18 Slice Append And Operations](./18_slice_append_and_operations/README.md)
    `append` とスライスの基本操作
19. [19 Maps](./19_maps/README.md)
    `map` の基本操作
20. [20 Functions Basic](./20_functions_basic/README.md)
    関数定義、基本形、シグネチャ
21. [21 Multiple Return Values](./21_multiple_return_values/README.md)
    複数の戻り値
22. [22 Named Return Values](./22_named_return_values/README.md)
    名前付き戻り値
23. [23 Variadic Functions](./23_variadic_functions/README.md)
    可変長引数
24. [24 Functions Are Values](./24_functions_are_values/README.md)
    関数は値である
25. [25 Higher Order Functions](./25_higher_order_functions/README.md)
    高階関数
26. [26 Functions And Closures](./26_functions_and_closures/README.md)
    関数とクロージャ
27. [27 Printf Formatting](./27_printf_formatting/README.md)
    フォーマット出力 `Printf`
28. [28 Pointers Basic](./28_pointers_basic/README.md)
    ポインタ
29. [29 Pointer Arguments](./29_pointer_arguments/README.md)
    ポインタを引数に使う
30. [30 Pointer And Slice](./30_pointer_and_slice/README.md)
    ポインタでスライスを操作する
31. [31 Structs Basic](./31_structs_basic/README.md)
    構造体
32. [32 Type Definitions](./32_type_definitions/README.md)
    `type` で型として定義する
33. [33 Structs And Pointers](./33_structs_and_pointers/README.md)
    構造体と参照渡し
34. [34 New And Make](./34_new_and_make/README.md)
    `new` と `make`
35. [35 Methods Basic](./35_methods_basic/README.md)
    メソッドをつかう
36. [36 Extending Types](./36_extending_types/README.md)
    既存の型を拡張する
37. [37 Pointer Receivers](./37_pointer_receivers/README.md)
    ポインタによるレシーバ
38. [38 Interfaces Basic](./38_interfaces_basic/README.md)
    インタフェースの定義、実装、利用
39. [39 Interfaces With Structs](./39_interfaces_with_structs/README.md)
    構造体とインタフェースを組み合わせる
40. [40 Interface Collections](./40_interface_collections/README.md)
    interface の配列で複数の構造体をまとめる
41. [41 Nil Receivers](./41_nil_receivers/README.md)
    nil レシーバ
42. [42 Empty Interface](./42_empty_interface/README.md)
    空のインタフェース `interface{}`
43. [43 Type Assertions](./43_type_assertions/README.md)
    型アサーション
44. [44 Reflect Type Checks](./44_reflect_type_checks/README.md)
    `reflect.TypeOf` と `Kind` で型を判定する
45. [45 Reflect Slice Checks](./45_reflect_slice_checks/README.md)
    `reflect` でスライス型を判定する
46. [46 Goroutines Basic](./46_goroutines_basic/README.md)
    `goroutine` の基本
47. [47 Channels Basic](./47_channels_basic/README.md)
    `channel` の基本的な送受信
48. [48 Channels Bidirectional](./48_channels_bidirectional/README.md)
    2 本の `channel` を使った双方向通信
49. [49 Select Basic](./49_select_basic/README.md)
    `select` による複数 `channel` の扱い
50. [50 Mutex Basic](./50_mutex_basic/README.md)
    `sync.Mutex` による共有データ保護
51. [51 Data Access File Access Basic](./51_file_access_basic/README.md)
    `os.ReadFile` `os.WriteFile` `os.OpenFile` によるデータアクセスの入口としてのファイルアクセス
52. [52 Network Access Basic](./52_network_access_basic/README.md)
    `http.Get` による HTTP アクセス、`Response.Body` の読み取りとクローズ
53. [53 WaitGroup Basic](./53_waitgroup_basic/README.md)
    `sync.WaitGroup` による複数 goroutine の完了待ち
54. [54 Worker Pool Basic](./54_worker_pool_basic/README.md)
    固定数のワーカー goroutine で channel からジョブを処理する Worker Pool パターン
55. [55 Race Detector And Atomic](./55_race_detector_and_atomic/README.md)
    レースコンディションの発生と `sync/atomic` による安全なカウンタ操作
56. [56 I/O Reader Writer Basic](./56_io_reader_writer_basic/README.md)
    `io.Reader` / `io.Writer` の基本と `io.Copy` によるストリーム転送
57. [57 Context Timeout Basic](./57_context_timeout_basic/README.md)
    `context.WithTimeout` と `context.WithCancel` によるタイムアウト・キャンセル制御
58. [58 Testing Table Driven](./58_testing_table_driven/README.md)
    テーブル駆動テストと `t.Run` によるサブテスト
59. [59 Benchmark And Alloc](./59_benchmark_and_alloc/README.md)
    `BenchmarkXxx` と `-benchmem` によるアロケーション計測
60. [60 HTTP Server Basic](./60_http_server_basic/README.md)
    `net/http` で書く最小構成の HTTP サーバ
61. [61 JSON API Basic](./61_json_api_basic/README.md)
    `encoding/json` による構造体と JSON の変換、構造体タグの使い方
62. [62 Database SQL Basic](./62_database_sql_basic/README.md)
    `database/sql` と SQLite を使った CRUD の基本
63. [63 Web To Markdown Basic](./63_web_to_markdown_basic/README.md)
    Web ページを取得して `html-to-markdown` で Markdown に変換し、ファイルに保存する

## 実行方法

各ディレクトリで次のように実行できます。

```bash
go run ./examples/handson/01_values_and_types
```

## テーマ別ガイド

`01` から `52` まで順番に進めるのが基本ルートです。

以下は通し学習の途中で復習しやすくするためのテーマ別整理です。学習順を選ぶための分岐ではなく、`01` から順番に進める中での補助ガイドとして使ってください。

### 値・変数・制御構文の基礎

1. `01_values_and_types`
2. `02_literals`
3. `03_operations`
4. `04_variables`
5. `05_short_variable_declaration`
6. `06_short_declaration_rules`
7. `09_constants`
8. `11_control_flow`
9. `12_switch_default_fallthrough`
10. `13_for_statement`
11. `14_for_init_and_post`
12. `15_infinite_loop_break_continue`
13. `16_labels_and_goto`

### 型変換と文字列・数値の扱い

1. `07_type_conversion`
2. `08_string_number_conversion`
3. `10_imports`
4. `27_printf_formatting`

### スライス・map・コレクション操作

1. `17_arrays_and_slices`
2. `18_slice_append_and_operations`
3. `19_maps`
4. `30_pointer_and_slice`
5. `34_new_and_make`

### 関数・高階関数・クロージャ

1. `20_functions_basic`
2. `21_multiple_return_values`
3. `22_named_return_values`
4. `23_variadic_functions`
5. `24_functions_are_values`
6. `25_higher_order_functions`
7. `26_functions_and_closures`

### ポインタ周辺

1. `28_pointers_basic`
2. `29_pointer_arguments`
3. `30_pointer_and_slice`
4. `33_structs_and_pointers`
5. `37_pointer_receivers`

### 構造体・メソッド・型設計

1. `31_structs_basic`
2. `32_type_definitions`
3. `33_structs_and_pointers`
4. `34_new_and_make`
5. `35_methods_basic`
6. `36_extending_types`
7. `37_pointer_receivers`

### インタフェースと型の見分け方

1. `38_interfaces_basic`
2. `39_interfaces_with_structs`
3. `40_interface_collections`
4. `41_nil_receivers`
5. `42_empty_interface`
6. `43_type_assertions`
7. `44_reflect_type_checks`
8. `45_reflect_slice_checks`

### 並行処理の入口

1. `46_goroutines_basic`
2. `47_channels_basic`
3. `48_channels_bidirectional`
4. `49_select_basic`
5. `50_mutex_basic`

### データアクセスとその周辺

1. `51_file_access_basic`
2. `52_network_access_basic`
3. `56_io_reader_writer_basic`
4. `57_context_timeout_basic`
5. `60_http_server_basic`
6. `61_json_api_basic`
7. `62_database_sql_basic`

### 並行処理の実践

1. `53_waitgroup_basic`
2. `54_worker_pool_basic`
3. `55_race_detector_and_atomic`

### I/O・context・テスト・計測

1. `56_io_reader_writer_basic`
2. `57_context_timeout_basic`
3. `58_testing_table_driven`
4. `59_benchmark_and_alloc`

### Web・JSON・DB 入門

1. `60_http_server_basic`
2. `61_json_api_basic`
3. `62_database_sql_basic`
4. `63_web_to_markdown_basic`

## 読み進め方

まずは `01` から `63` まで順番に進めてください。

以下は途中で振り返るための整理です。

- `01` から `03` で値、型、演算の基礎を確認する
- `04` から `06` で変数宣言とスコープを整理する
- `07` `08` で型変換を確認する
- `09` `10` で定数と import を押さえる
- `11` `12` で条件分岐を学ぶ
- `13` `14` で繰り返し処理を学ぶ
- `15` で無限ループとループ制御を学ぶ
- `16` でラベルと `goto` の基本を確認する
- `17` で配列とスライスの基礎を学ぶ
- `18` で `append` とスライス操作を学ぶ
- `19` で `map` の基本を学ぶ
- `20` で関数の基本を学ぶ
- `21` で複数の戻り値を学ぶ
- `22` で名前付き戻り値を学ぶ
- `23` で可変長引数を学ぶ
- `24` で関数を値として扱う感覚をつかむ
- `25` で高階関数を学ぶ
- `26` でクロージャの動きを学ぶ
- `27` で `Printf` の書式指定を学ぶ
- `28` でポインタの基本を学ぶ
- `29` でポインタ引数と値渡しの違いを学ぶ
- `30` でポインタとスライスの関係を学ぶ
- `31` で構造体の基本を学ぶ
- `32` で独自型の定義を学ぶ
- `33` で構造体とポインタ渡しの違いを学ぶ
- `34` で `new` と `make` の違いを学ぶ
- `35` でメソッドの基本を学ぶ
- `36` で既存型を元に独自の振る舞いを足す方法を学ぶ
- `37` でポインタレシーバを学ぶ
- `38` でインタフェースの基本を学ぶ
- `39` で interface 型として扱う感覚を深める
- `40` で interface の配列に複数の構造体をまとめる
- `41` で nil レシーバの扱いを確認する
- `42` で空のインタフェースを学ぶ
- `43` で型アサーションの基本を学ぶ
- `44` で `reflect.TypeOf` と `Kind` による型判定を学ぶ
- `45` でスライス型の判定を学ぶ
- `46` で `goroutine` の基本を学ぶ
- `47` で `channel` の送信と受信を学ぶ
- `48` で channel を役割ごとに分ける形を学ぶ
- `49` で `select` による複数 channel の扱いを学ぶ
- `50` で `Mutex` による共有データ保護を学ぶ
- `51` でデータアクセスの入口として、ファイルを開く、読む、書く、閉じる流れを学ぶ
- `52` でネットワークアクセスの基本として、HTTP 経由でデータを取得する流れを学ぶ
- `53` で `WaitGroup` による複数 goroutine の完了待ちを学ぶ
- `54` で Worker Pool パターンを学ぶ
- `55` でレースコンディションと `atomic` による安全な操作を学ぶ
- `56` で `io.Reader` / `io.Writer` によるストリーム処理を学ぶ
- `57` で `context` によるタイムアウト・キャンセル制御を学ぶ
- `58` でテーブル駆動テストを学ぶ
- `59` でベンチマークとアロケーション計測を学ぶ
- `60` で `net/http` による最小 HTTP サーバを学ぶ
- `61` で JSON のシリアライズ・デシリアライズを学ぶ
- `62` で `database/sql` による CRUD の基本を学ぶ
- `63` で Web ページを取得して Markdown に変換し、ファイルへ保存する流れを学ぶ
