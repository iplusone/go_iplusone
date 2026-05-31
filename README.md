# go_iplusone

Go 言語自体を習得するための言語仕様の整理から、実務で必要となる開発ナレッジの体系化までを一貫して扱い、それらを蓄積・再利用できるナレッジ基盤を Go で構築するプロジェクトです。

GoSpec をはじめとする言語仕様の理解を土台に、型システム、並行処理、初期化、標準ライブラリ、設計原則、クラウド実装までを対象にし、学習・設計・実装を支えるコンテンツを育てていきます。

現在の公開フロントは、GoSpec 学習導線に絞った Go 学習サイトとして整理しています。

## 起動方法

```bash
docker compose up --build
```

ブラウザで `http://localhost:8021` にアクセスしてください。

## 主な公開URL

- `/`
- `/healthz`
- `/news`
- `/news/detail/{id}`
- `/gospec/{chapter-slug}`

## 構成

- `main.go`
  アプリ起動エントリポイント
- `docs/start-here.md`
  最初に読むページ。全部やるための唯一の入口
- `examples/handson`
  Go の基礎文法と言語機能を順番に固めるハンズオン
  目次: `examples/handson/README.md`
- `examples/master`
  Docker、Web、DB、計測、運用までつなぐ実践ハンズオン（Week 1〜12）
  目次: `examples/master/README.md`
- `examples/appendix`
  本編の外にある付録（Fyne GUI・Web サーバ）
  目次: `examples/appendix/README.md`
- `internal/site`
  ルーティング、埋め込みテンプレート、初期データ
- `docs/gospec-learning-guide.md`
  Go 言語仕様書を学ぶための全体ガイド
- `docs/go-master-curriculum.md`
  Go を使って実装技術と計算機科学の知見を学ぶ 12 週間カリキュラムの正本
- `docs/go-master-curriculum-handson-mapping.md`
  正本カリキュラムと既存 handson の対応関係を整理した表
- `docs/go-master-curriculum-gap-list.md`
  全体構成とハンズオンメニューの不足一覧
- `docs/handson-learning-guide.md`
  Go 基礎ハンズオンの学習順と構成を整理したガイド
- `docs/go-advanced-syntax-roadmap.md`
  Go の高度な文法をどう学ぶかのロードマップ
- `docs/top-page-structure.md`
  Go 学習トップの構成メモ

## ハンズオン実行例

最初の入口は [docs/start-here.md](/home/ishii/projects/go_iplusone/docs/start-here.md) です。以後の学習導線は、そこから辿ってください。

このリポジトリでは、3段階のルートを順番に進めます。

- `examples/handson`
  Go の文法、型、関数、ポインタ、構造体、interface、並行処理・I/O・テスト・HTTP・JSON・DB の基礎を固める（01〜62）
- `examples/master`
  その知識を Docker、Web、DB、計測、運用へ接続する（Week 1〜12）
- `examples/appendix`
  本編を終えた後に取り組む付録（Fyne GUI・Web サーバ）

現在の `examples/handson` は `01_values_and_types` から `63_web_to_markdown_basic` まで追加しており、値・制御構文・コレクション・関数・ポインタ・構造体・メソッド・インタフェース・`reflect`・並行処理・データアクセス・I/O・`context`・テスト・ベンチマーク・HTTP・JSON・`database/sql`・Web ページの Markdown 変換まで学べる状態です。

```bash
go run ./examples/handson/01_values_and_types
```

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

### 並行処理の実践

1. `53_waitgroup_basic`
2. `54_worker_pool_basic`
3. `55_race_detector_and_atomic`

### データアクセスの入口

1. `51_file_access_basic`
2. `52_network_access_basic`
3. `56_io_reader_writer_basic`
4. `57_context_timeout_basic`

### テスト・ベンチマーク

1. `58_testing_table_driven`
2. `59_benchmark_and_alloc`

### Web・JSON・DB 入門

1. `60_http_server_basic`
2. `61_json_api_basic`
3. `62_database_sql_basic`
4. `63_web_to_markdown_basic`

## 付録

`examples/master` を終えたあとは付録へ進めます。

- `examples/appendix/01_fyne_gui`
  Go によるクロスプラットフォーム GUI アプリケーション開発（Fyne）
- `examples/appendix/02_web_server`
  `net/http` を使った Web サーバの全体像（テンプレート・ミドルウェア・静的配信）

## 実装メモ

- テンプレートと CSS は `embed` でバイナリに同梱しています
- NEWS と GoSpec 章データは初期段階としてインメモリデータで描画しています
- `docker-compose.yml` に MySQL はありますが、現段階の公開フロントでは未接続です
