# 正本カリキュラムと既存ハンズオンの対応表

このドキュメントは、[docs/go-master-curriculum.md](go-master-curriculum.md) を正本としたときに、現在の `examples/handson` がどこまで対応しているかを整理したものです。

結論を先に書くと、現在の `examples/handson` は `Go 基礎文法ハンズオン` としてはよく整理されていますが、`12週間の正本カリキュラムに対応する実践ハンズオン` としてはまだ不足があります。

## 判定基準

対応状況は次の3段階で整理します。

- `対応`
  その週の中心テーマを、現在の handson だけで概ね学べる
- `部分対応`
  関連する言語機能はあるが、正本カリキュラムが求める実践課題や計測・環境前提までは満たしていない
- `未対応`
  現在の handson だけでは、その週の中心テーマをカバーできない

## 全体評価

| 週 | 正本カリキュラムの大項目 | 対応状況 | コメント |
| --- | --- | --- | --- |
| 1 | Docker ベースの開発基盤 | 未対応 | `examples/handson` は Go 単体実行が中心で、Docker Compose、Web、DB、Front の E2E 疎通は未整備 |
| 2 | メモリ管理と動的データ | 部分対応 | ポインタ、スライス、文字列の基本はあるが、ベンチマークやアドレス観測などの物理挙動確認が不足 |
| 3 | データ集約とカプセル化 | 部分対応 | 構造体、メソッド、可視性の入口はあるが、アライメントや package 境界設計は未整備 |
| 4 | 抽象化と疎結合設計 | 部分対応 | interface の基本はあるが、DIP、Mock、差し替え設計までは未整備 |
| 5 | I/O とリソース管理 | 未対応 | `io.Reader`/`Writer`、`bufio`、`context`、`defer Close` を主題にした handson がない |
| 6 | 品質保証と性能計測 | 未対応 | `testing`、ベンチマーク、エスケープ解析、`pprof` を扱う handson がない |
| 7 | 並行処理の基本原理 | 対応 | `goroutine`、`channel`、`select` の入口は `46`〜`49` でカバーできる |
| 8 | 高度な並行設計 | 部分対応 | `Mutex` はあるが、`WaitGroup` 単独テーマ、Worker Pool、`-race`、`atomic` は不足 |
| 9 | Web プロトコルと通信 | 未対応 | `net/http`、ミドルウェア、JSON API の handson がない |
| 10 | 永続化と整合性 | 未対応 | `database/sql`、接続管理、トランザクションの handson がない |
| 11 | システムアーキテクチャ | 未対応 | package 構成、レイヤ分離、構造化ログ、エラー戦略の handson がない |
| 12 | 配布最適化と運用 | 未対応 | `scratch`、静的リンク、ヘルスチェック、グレースフルシャットダウンの handson がない |

## 週ごとの詳細対応

### 第1週: Docker ベースの開発基盤

- 対応状況: `未対応`
- 正本で求める内容:
  `Dockerfile`、`docker-compose.yml`、コンテナ間通信、Web + DB + Front の E2E 疎通
- 現在の handson:
  なし
- コメント:
  現在の handson は `go run ./examples/handson/...` を前提にした基礎文法編であり、第1週の Docker/Linux 基準には乗っていません

### 第2週: メモリ管理と動的データ

- 対応状況: `部分対応`
- 関連 handson:
  - [17_arrays_and_slices](../examples/handson/17_arrays_and_slices/README.md)
  - [18_slice_append_and_operations](../examples/handson/18_slice_append_and_operations/README.md)
  - [28_pointers_basic](../examples/handson/28_pointers_basic/README.md)
  - [29_pointer_arguments](../examples/handson/29_pointer_arguments/README.md)
  - [30_pointer_and_slice](../examples/handson/30_pointer_and_slice/README.md)
- 足りない要素:
  - 値渡しと参照渡しのベンチマーク比較
  - スライス再確保時のアドレス観測
  - UTF-8 のバイト表現を確認する演習

### 第3週: データ集約とカプセル化

- 対応状況: `部分対応`
- 関連 handson:
  - [31_structs_basic](../examples/handson/31_structs_basic/README.md)
  - [33_structs_and_pointers](../examples/handson/33_structs_and_pointers/README.md)
  - [35_methods_basic](../examples/handson/35_methods_basic/README.md)
  - [37_pointer_receivers](../examples/handson/37_pointer_receivers/README.md)
- 足りない要素:
  - `unsafe.Sizeof` を使ったアライメント観測
  - package を跨いだ可視性設計
  - `internal` を使った境界設計

### 第4週: 抽象化と疎結合設計

- 対応状況: `部分対応`
- 関連 handson:
  - [38_interfaces_basic](../examples/handson/38_interfaces_basic/README.md)
  - [39_interfaces_with_structs](../examples/handson/39_interfaces_with_structs/README.md)
  - [40_interface_collections](../examples/handson/40_interface_collections/README.md)
  - [42_empty_interface](../examples/handson/42_empty_interface/README.md)
  - [43_type_assertions](../examples/handson/43_type_assertions/README.md)
- 足りない要素:
  - DIP を意識した依存方向の分離
  - Mock/Fake を使ったテスト容易性の体験
  - DB 実装とメモリ実装の差し替え演習

### 第5週: I/O とリソース管理

- 対応状況: `未対応`
- 正本で求める内容:
  `io.Reader`、`io.Writer`、`bufio`、`os.File`、`defer`、`context`
- 現在の handson:
  なし
- コメント:
  現状の handson は、標準ライブラリの入門や `Printf` までで止まっており、ストリーム処理や資源解放に踏み込めていません

### 第6週: 品質保証と性能計測

- 対応状況: `未対応`
- 正本で求める内容:
  `testing`、テーブル駆動テスト、ベンチマーク、`-benchmem`、`-gcflags=-m`、`pprof`
- 現在の handson:
  なし
- コメント:
  handson 自体はあるものの、テストコードやベンチマーク文化はまだカリキュラム化されていません

### 第7週: 並行処理の基本原理

- 対応状況: `対応`
- 関連 handson:
  - [46_goroutines_basic](../examples/handson/46_goroutines_basic/README.md)
  - [47_channels_basic](../examples/handson/47_channels_basic/README.md)
  - [48_channels_bidirectional](../examples/handson/48_channels_bidirectional/README.md)
  - [49_select_basic](../examples/handson/49_select_basic/README.md)
- コメント:
  正本の入口に相当する内容はすでに追加済みです

### 第8週: 高度な並行設計

- 対応状況: `部分対応`
- 関連 handson:
  - [50_mutex_basic](../examples/handson/50_mutex_basic/README.md)
  - [46_goroutines_basic](../examples/handson/46_goroutines_basic/README.md)
- 足りない要素:
  - `sync.WaitGroup` を明示的に主題化した回
  - Worker Pool、Pipeline、Fan-in/Fan-out
  - `go run -race` を使った競合検出
  - `sync/atomic` による比較

### 第9週: Web プロトコルと通信

- 対応状況: `未対応`
- 正本で求める内容:
  `net/http`、Handler、ミドルウェア、`encoding/json`
- 現在の handson:
  なし
- コメント:
  リポジトリ本体には Web アプリはあるものの、学習用に分解した handson はまだありません

### 第10週: 永続化と整合性

- 対応状況: `未対応`
- 正本で求める内容:
  `database/sql`、ドライバ、接続管理、プリペアドステートメント、トランザクション
- 現在の handson:
  なし
- コメント:
  Docker Compose に DB は存在しますが、handson としては未整備です

### 第11週: システムアーキテクチャ

- 対応状況: `未対応`
- 正本で求める内容:
  `cmd`、`internal`、レイヤ分離、エラー戦略、構造化ログ
- 現在の handson:
  なし
- コメント:
  学習対象としては重要ですが、単発サンプルより小さなアプリのリファクタリング課題として作る方が自然です

### 第12週: 配布最適化と運用

- 対応状況: `未対応`
- 正本で求める内容:
  `scratch`、静的リンク、グレースフルシャットダウン、ヘルスチェック
- 現在の handson:
  なし
- コメント:
  ここは単独の文法学習ではなく、実行可能な小規模アプリを対象にした仕上げ課題が必要です

## どこが合っていて、どこが違うのか

### 今の handson が強いところ

- Go の基本文法を段階的に学ぶ流れが明快
- 値、制御構文、関数、ポインタ、構造体、interface まで素直に辿れる
- 並行処理の入口も `46`〜`50` で追加できている

### 正本カリキュラムに対して足りないところ

- Docker/Linux/Web/DB を基準にした週構成になっていない
- `文法の確認` はあるが、`物理挙動の観測` や `実装上のトレードオフ` まで踏み込めていない
- テスト、ベンチマーク、HTTP、DB、運用を学ぶ実践編が未整備

## 再編方針

既存の `examples/handson` は、無理に壊さず `基礎文法編` として活かすのが最も自然です。

そのうえで、正本カリキュラムに対応する新しい系統を別に持つ方が整理しやすいです。

推奨する構成:

- `examples/handson`
  Go 基礎文法編
- `examples/master`
  正本カリキュラム対応の実践編

例:

- `examples/master/01_docker_e2e`
- `examples/master/02_memory_layout`
- `examples/master/03_struct_and_capsule`
- `examples/master/04_abstraction_and_di`
- `examples/master/05_io_and_resource_lifecycle`

## この対応表から導ける結論

現在の handson は `間違っている` のではなく、`役割が違う` という整理が適切です。

- 現在の handson
  Go 基礎文法を習得する教材
- 正本カリキュラム
  Go を使って実装技術と計算機科学の知見を習得する教材

したがって、今後の整備方針としては:

1. `examples/handson` を基礎文法編として維持する
2. `examples/master` を正本カリキュラム対応編として新設する
3. 各週の成果物は `examples/master` 側で積み上げる

の3点を基本方針にするのが最も整合的です。
