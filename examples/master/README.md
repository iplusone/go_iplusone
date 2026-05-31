# Master Hands-on Index

正本カリキュラム [docs/go-master-curriculum.md](/home/ishii/projects/go_iplusone/docs/go-master-curriculum.md) に対応する実践編ハンズオンの目次です。

このディレクトリは、`examples/handson` を終えたあとに別ルートで選ぶものではありません。

`examples/handson` で Go の基礎文法と言語機能を固め、そのうえで `examples/master` で Docker、Web、DB、性能計測、運用までを一続きで実践します。

つまり、このリポジトリの学習は次の一本道です。

1. 正本カリキュラムを読む
2. `examples/handson` で基礎を固める
3. `examples/master` で実践へつなぐ

## 週対応表

| 週 | ディレクトリ | テーマ | 状態 |
| --- | --- | --- | --- |
| 1 | [01_docker_e2e](./01_docker_e2e/README.md) | Docker ベースの開発基盤 | 実装済み |
| 2 | [02_memory_layout](./02_memory_layout/README.md) | メモリ管理と動的データ | 実装済み |
| 3 | [03_struct_and_capsule](./03_struct_and_capsule/README.md) | データ集約とカプセル化 | 実装済み |
| 4 | [04_abstraction_and_di](./04_abstraction_and_di/README.md) | 抽象化と疎結合設計 | 実装済み |
| 5 | [05_io_and_resource_lifecycle](./05_io_and_resource_lifecycle/README.md) | I/O とリソース管理 | 実装済み |
| 6 | [06_testing_and_benchmark](./06_testing_and_benchmark/README.md) | 品質保証と性能計測 | 実装済み |
| 7 | [07_concurrency_fundamentals](./07_concurrency_fundamentals/README.md) | 並行処理の基本原理 | 実装済み |
| 8 | [08_concurrency_patterns](./08_concurrency_patterns/README.md) | 高度な並行設計 | 実装済み |
| 9 | [09_http_and_json](./09_http_and_json/README.md) | Web プロトコルと通信 | 実装済み |
| 10 | [10_db_and_transaction](./10_db_and_transaction/README.md) | 永続化と整合性 | 実装済み |
| 11 | [11_architecture_refactor](./11_architecture_refactor/README.md) | システムアーキテクチャ | 実装済み |
| 12 | [12_deploy_and_operations](./12_deploy_and_operations/README.md) | 配布最適化と運用 | 実装済み |

## 目次

1. [01 Docker E2E](./01_docker_e2e/README.md)
   Docker Compose で `front` `app` `db` を立ち上げ、ブラウザから Go 経由で DB のデータを表示する
2. [02 Memory Layout](./02_memory_layout/README.md)
   ポインタ、値渡し・参照渡し、スライスの内部構造、文字列のバイト表現を実験する
3. [03 Struct And Capsule](./03_struct_and_capsule/README.md)
   構造体の物理配置、メソッド定義、パッケージによるカプセル化を実践する
4. [04 Abstraction And DI](./04_abstraction_and_di/README.md)
   interface による抽象化と依存性注入でテスト可能な設計を作る
5. [05 I/O And Resource Lifecycle](./05_io_and_resource_lifecycle/README.md)
   `io.Reader/Writer`、`defer`、`context` で外部リソースを安全に扱う
6. [06 Testing And Benchmark](./06_testing_and_benchmark/README.md)
   テーブル駆動テスト、ベンチマーク、エスケープ解析で正しさと性能を証明する
7. [07 Concurrency Fundamentals](./07_concurrency_fundamentals/README.md)
   goroutine、channel、select の基本原理を理解する
8. [08 Concurrency Patterns](./08_concurrency_patterns/README.md)
   Mutex、WaitGroup、Worker Pool で実戦的な並行設計を実装する
9. [09 HTTP And JSON](./09_http_and_json/README.md)
   `net/http`、ミドルウェア、`encoding/json` で JSON API を構築する
10. [10 DB And Transaction](./10_db_and_transaction/README.md)
    `database/sql`、プリペアドステートメント、トランザクションで永続化を扱う
11. [11 Architecture Refactor](./11_architecture_refactor/README.md)
    レイヤ分離、エラー戦略、構造化ログでシステムアーキテクチャを整える
12. [12 Deploy And Operations](./12_deploy_and_operations/README.md)
    scratch イメージ、グレースフルシャットダウン、ヘルスチェックで運用可能なアプリにする

## 付録

本編 12 週を終えたあとは付録へ進めます。

- [examples/appendix/README.md](../appendix/README.md)
  Fyne GUI アプリケーション開発、Web サーバ構築ガイド

## 位置づけ

- `examples/handson`
  Go 基礎文法と言語機能の土台
- `examples/master`
  その土台を Docker、Web、DB、運用まで接続する実践編
- `examples/appendix`
  本編の知識を別の舞台（GUI、Web）へ応用する付録

## 実行方法

各ディレクトリの `README.md` に従ってください。
