# 正本カリキュラム整備の不足一覧

このドキュメントは、[docs/go-master-curriculum.md](/home/ishii/projects/go_iplusone/docs/go-master-curriculum.md) を正本としたときに、現在の全体構成とハンズオンメニューで足りていない部分を整理したタスクリストです。

目的は次の2つです。

- 何が未整備なのかを曖昧にしない
- 次にどこから着手するべきかを優先順位付きで共有する

## 結論

現時点では、次がすべて整備済みです。

- 正本カリキュラム
- 正本と既存 handson の対応表
- `examples/handson` による基礎文法編（01〜62）
- `examples/master` による実践編（01〜12、全週実装済み）
- 全体導線の一本化（`docs/start-here.md`）

以下の不足一覧は整備完了済みの記録として残します。

## 不足一覧

### 1. 全体構成で足りないもの

#### 1-1. 学習の入口がまだ分散している

現状:

- `README.md`
- `docs/go-master-curriculum.md`
- `docs/handson-learning-guide.md`
- `docs/go-master-curriculum-handson-mapping.md`

に学習導線が分散しています。

問題:

- 初学者がどこから読むべきか迷いやすい
- 基礎文法編と実践編の違いが、文書をまたいで理解しないと分かりにくい

必要な対応:

- `最初に読むページ` を1つに決める
- そのページから `基礎文法編` と `実践編` を分岐できるようにする

#### 1-2. 各週の詳細タスク文書がない

現状:

- 正本には `大項目 / 中項目 / 小項目 / 取得目的 / 関連知見 / 成果物` はある
- しかし `週ごとの具体タスク` はまだ別紙化されていない

問題:

- 実際に何を作れば週を完了とみなすかが曖昧
- 学習者が途中で「何をやればよいか」分からなくなる

必要な対応:

- 第1週から第12週まで、各週の詳細タスク文書を `docs/` に分割して作る

#### 1-3. 実践編の進捗一覧がない

現状:

- `examples/master/README.md` はある
- ただし、どの週が実装済みでどの週が未着手か分からない

問題:

- 現在の整備状況が一目で把握できない
- 次の優先着手対象を決めづらい

必要な対応:

- `examples/master/README.md` に進捗状態を追加する

例:

- `01_docker_e2e`
  実装済み
- `02_memory_layout`
  未着手

#### 1-4. 付録モジュールの置き場がない ✅

`examples/appendix/` を新設しました。

- `examples/appendix/01_fyne_gui/README.md` — Fyne GUI 開発ガイド
- `examples/appendix/02_web_server/README.md` — Web サーバ構築ガイド

### 2. ハンズオンメニューで足りないもの

#### 2-1. `examples/master` が第1週しかない

現状:

- `examples/master/01_docker_e2e` のみ存在

問題:

- 正本カリキュラムに対応する実践編としてはほぼ未整備

必要な対応:

- 第2週以降を順次追加する

#### 2-2. `examples/handson/README.md` が基礎文法編の立ち位置を十分に説明していない

現状:

- 基礎文法を順番に学ぶ目次としては十分

問題:

- `examples/master` との違いが README 単体では見えにくい
- 正本対応編と混同される可能性がある

必要な対応:

- `examples/handson` は基礎文法編であることを明記する
- `examples/master` への導線を加える

#### 2-3. `examples/handson/README.md` にテーマ別ガイドがない

現状:

- ルート README にはテーマ別導線を追加済み
- handson 目次本体には未反映

問題:

- handson 目次だけを見た人には、テーマ別の学び方が見えない

必要な対応:

- `examples/handson/README.md` にも README と同様のテーマ別導線を追加する

#### 2-4. `examples/master/README.md` に週対応表がない

現状:

- 目次はある
- ただし `正本カリキュラムの何週に対応するか` が弱い

問題:

- カリキュラムと実践編の関係がすぐに見えない

必要な対応:

- `Week 1 -> 01_docker_e2e` のような対応表を入れる

#### 2-5. 各 master handson の共通フォーマットが未定義

現状:

- `01_docker_e2e` は README を個別に整えている

問題:

- 今後の回で README 構成がぶれる可能性がある

必要な対応:

- 各回の README に共通の章立てを決める

推奨フォーマット:

- 目的
- 何を確認するか
- 構成
- 起動方法
- 到達基準
- 補足観点

### 3. 正本カリキュラムに対する未実装領域

以下は、正本カリキュラムに対して `examples/master` 側で未実装のものです。

| 週 | 想定ディレクトリ | 状態 |
| --- | --- | --- |
| 1 | `01_docker_e2e` | 実装済み |
| 2 | `02_memory_layout` | 未着手 |
| 3 | `03_struct_and_capsule` | 未着手 |
| 4 | `04_abstraction_and_di` | 未着手 |
| 5 | `05_io_and_resource_lifecycle` | 未着手 |
| 6 | `06_testing_and_benchmark` | 未着手 |
| 7 | `07_concurrency_fundamentals` | 未着手 |
| 8 | `08_concurrency_patterns` | 未着手 |
| 9 | `09_http_and_json` | 未着手 |
| 10 | `10_db_and_transaction` | 未着手 |
| 11 | `11_architecture_refactor` | 未着手 |
| 12 | `12_deploy_and_operations` | 未着手 |

### 4. 部分対応しているが補強が必要な handson

既存 `examples/handson` のうち、正本の観点で補助回があると強くなるテーマは次です。

| 補強テーマ | 追加候補 |
| --- | --- |
| 並行処理の同期 | `51_waitgroup_basic` |
| 並行処理パターン | `52_worker_pool_basic` |
| レース検知と atomic | `53_race_detector_and_atomic` |
| I/O の基本 | `54_io_reader_writer_basic` |
| context の基本 | `55_context_timeout_basic` |
| テスト | `56_testing_table_driven` |
| ベンチマーク | `57_benchmark_and_alloc` |
| HTTP 入門 | `58_http_server_basic` |
| JSON API 入門 | `59_json_api_basic` |
| DB 入門 | `60_database_sql_basic` |

## 優先順位

### 最優先

1. `examples/master/02_memory_layout`
2. `examples/master/03_struct_and_capsule`
3. `examples/master/04_abstraction_and_di`

理由:

- 第1週の次に自然につながる
- 正本カリキュラムの序盤を固められる
- `基礎文法編` と `実践編` の違いを最も見せやすい

### 優先度高

1. `examples/master/05_io_and_resource_lifecycle`
2. `examples/master/06_testing_and_benchmark`
3. `examples/master/07_concurrency_fundamentals`
4. `examples/master/08_concurrency_patterns`

理由:

- 正本カリキュラムの中盤の核になる
- 今の `examples/handson` だけでは不足が大きい

### 優先度中

1. `examples/master/09_http_and_json`
2. `examples/master/10_db_and_transaction`
3. `examples/master/11_architecture_refactor`
4. `examples/master/12_deploy_and_operations`

理由:

- 実践的で価値は高い
- ただし序盤〜中盤の積み上げがある方が作りやすい

### ドキュメント整備

1. `examples/handson/README.md` にテーマ別導線を追加
2. `examples/master/README.md` に週対応表と進捗状態を追加
3. 各週の詳細タスク文書を `docs/` に追加

## 整備方針

今後の基本方針は次のとおりです。

1. `examples/handson` は基礎文法編として維持する
2. `examples/master` は正本カリキュラム対応の実践編として拡張する
3. `examples/master` は週単位の成果物として積み上げる
4. ドキュメントは `正本 -> 対応表 -> 不足一覧` の順で参照できるようにする

## 次の具体タスク

以下は整備完了済みです。

1. ✅ `examples/handson/README.md` にテーマ別ガイドを追加する
2. ✅ `examples/master/README.md` に週対応と進捗状態を追加する
3. ✅ `examples/master/02_memory_layout` を新設する
4. ✅ `examples/master/03_struct_and_capsule` を新設する
5. ✅ `examples/master/04_abstraction_and_di` を新設する
6. ✅ `examples/master/05_io_and_resource_lifecycle` を新設する
7. ✅ `examples/master/06_testing_and_benchmark` を新設する
8. ✅ `examples/master/07_concurrency_fundamentals` を新設する
9. ✅ `examples/master/08_concurrency_patterns` を新設する
10. ✅ `examples/master/09_http_and_json` を新設する
11. ✅ `examples/master/10_db_and_transaction` を新設する
12. ✅ `examples/master/11_architecture_refactor` を新設する
13. ✅ `examples/master/12_deploy_and_operations` を新設する
14. ✅ `examples/handson/53〜62` を新設する（WaitGroup、Worker Pool、atomic、I/O、context、テスト、ベンチマーク、HTTP、JSON、DB）
