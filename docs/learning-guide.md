# Go 実装技術マスター 学習ガイド

---

## はじめに

このガイドは、`go_iplusone` リポジトリを使って学習を進めるための手引きです。

このリポジトリの目的は、**Go の文法を覚えること**ではなく、**プログラミング全般の知見と計算機の仕組みを理解したうえで、Go で実装できるようになること**です。

学習は次の3段階で構成されています。

```
examples/handson  →  examples/master  →  examples/appendix
  基礎文法編            実践編（12週）         付録
```

---

## 学習の進め方

### ルールは1つ

**順番通りに進める。**

`examples/handson` の `01` から始め、`63` まで進めます。その後 `examples/master` の Week 1 から Week 12 へ進みます。付録は本編完了後に自由に選びます。

既知の内容であっても、成果物の接続と再現性のために順番通りに通過します。

### 迷ったときは

`docs/start-here.md` に戻ります。ここが唯一の入口です。

---

## 第1部: 基礎文法編（examples/handson）

Go の文法・型・言語機能を順番に固めます。全 **63 件**。

### 値・変数・制御構文

| # | タイトル | 学ぶこと |
|---|---|---|
| 01 | Values and Types | 値と変数、各種型（int, float, bool, string） |
| 02 | Literals | 整数・実数・テキスト・真偽値リテラル |
| 03 | Operations | 数値・文字列の演算 |
| 04 | Variables | 変数宣言・代入・スコープ |
| 05 | Short Variable Declaration | 暗黙的型宣言 |
| 06 | Short Declaration Rules | 短縮変数宣言のルール |
| 07 | Type Conversion | 型変換 |
| 08 | String Number Conversion | 文字列と数値の変換 |
| 09 | Constants | 定数 |
| 10 | Imports | `import` の基本 |
| 11 | Control Flow | `if`、`else`、`switch` |
| 12 | Switch Default Fallthrough | `switch` の `default` と `fallthrough` |
| 13 | For Statement | `for` の基本 |
| 14 | For Init And Post | `for` の初期化・条件・後処理 |
| 15 | Infinite Loop Break Continue | 無限ループ、`break`、`continue` |
| 16 | Labels And Goto | ラベル、`goto`、ラベル付き `break` |

### コレクション

| # | タイトル | 学ぶこと |
|---|---|---|
| 17 | Arrays And Slices | 配列、`range`、スライス |
| 18 | Slice Append And Operations | `append` とスライス操作 |
| 19 | Maps | `map` の基本操作 |

### 関数・高階関数・クロージャ

| # | タイトル | 学ぶこと |
|---|---|---|
| 20 | Functions Basic | 関数定義・シグネチャ |
| 21 | Multiple Return Values | 複数の戻り値 |
| 22 | Named Return Values | 名前付き戻り値 |
| 23 | Variadic Functions | 可変長引数 |
| 24 | Functions Are Values | 関数は値である |
| 25 | Higher Order Functions | 高階関数 |
| 26 | Functions And Closures | クロージャ |
| 27 | Printf Formatting | `Printf` フォーマット出力 |

### ポインタ

| # | タイトル | 学ぶこと |
|---|---|---|
| 28 | Pointers Basic | ポインタの基本 |
| 29 | Pointer Arguments | ポインタを引数に使う |
| 30 | Pointer And Slice | ポインタでスライスを操作する |

### 構造体・メソッド・型設計

| # | タイトル | 学ぶこと |
|---|---|---|
| 31 | Structs Basic | 構造体の基本 |
| 32 | Type Definitions | `type` で型を定義する |
| 33 | Structs And Pointers | 構造体と参照渡し |
| 34 | New And Make | `new` と `make` の違い |
| 35 | Methods Basic | メソッドの基本 |
| 36 | Extending Types | 既存型の拡張 |
| 37 | Pointer Receivers | ポインタレシーバ |

### インタフェース・型の見分け方

| # | タイトル | 学ぶこと |
|---|---|---|
| 38 | Interfaces Basic | インタフェースの定義・実装・利用 |
| 39 | Interfaces With Structs | 構造体とインタフェースの組み合わせ |
| 40 | Interface Collections | interface の配列で複数構造体をまとめる |
| 41 | Nil Receivers | nil レシーバ |
| 42 | Empty Interface | 空のインタフェース `interface{}` |
| 43 | Type Assertions | 型アサーション |
| 44 | Reflect Type Checks | `reflect.TypeOf` と `Kind` で型判定 |
| 45 | Reflect Slice Checks | `reflect` でスライス型を判定する |

### 並行処理

| # | タイトル | 学ぶこと |
|---|---|---|
| 46 | Goroutines Basic | goroutine の基本 |
| 47 | Channels Basic | channel の基本的な送受信 |
| 48 | Channels Bidirectional | 2 本の channel を使った双方向通信 |
| 49 | Select Basic | `select` による複数 channel の扱い |
| 50 | Mutex Basic | `sync.Mutex` による共有データ保護 |
| 53 | WaitGroup Basic | `sync.WaitGroup` による完了待ち |
| 54 | Worker Pool Basic | 固定数のワーカー goroutine による Worker Pool |
| 55 | Race Detector And Atomic | レースコンディションと `sync/atomic` |

### データアクセス

| # | タイトル | 学ぶこと |
|---|---|---|
| 51 | File Access Basic | `os.ReadFile`、`WriteFile`、`OpenFile` によるファイル操作 |
| 52 | Network Access Basic | `http.Get` による HTTP アクセス |
| 56 | I/O Reader Writer Basic | `io.Reader` / `io.Writer` とストリーム処理 |
| 57 | Context Timeout Basic | `context.WithTimeout` によるタイムアウト制御 |

### テスト・ベンチマーク

| # | タイトル | 学ぶこと |
|---|---|---|
| 58 | Testing Table Driven | テーブル駆動テストと `t.Run` |
| 59 | Benchmark And Alloc | ベンチマークとアロケーション計測 |

### Web・JSON・DB 入門

| # | タイトル | 学ぶこと |
|---|---|---|
| 60 | HTTP Server Basic | `net/http` で書く最小 HTTP サーバ |
| 61 | JSON API Basic | `encoding/json` による JSON 変換 |
| 62 | Database SQL Basic | `database/sql` と SQLite による CRUD |
| 63 | Web To Markdown Basic | Web ページを取得して Markdown に変換・保存 |

### 実行方法

各ディレクトリで次のように実行します（63 は除く）。

```bash
go run ./examples/handson/01_values_and_types
```

63 は独立したモジュールのため、ディレクトリ内で実行します。

```bash
cd examples/handson/63_web_to_markdown_basic
go run .
```

---

## 第2部: 実践編（examples/master）

`examples/handson` で固めた知識を、Docker・Web・DB・運用まで接続する **12 週間カリキュラム**です。

### 環境の前提

- Docker と Docker Compose がインストールされていること
- Linux 環境（または Docker 上の Linux）を基準にします

### 週ごとの内容

| 週 | テーマ | ゴール |
|---|---|---|
| **Week 1** | Docker ベースの開発基盤 | Docker Compose で front / app / db を起動し、ブラウザから DB データを表示する |
| **Week 2** | メモリ管理と動的データ | 値渡しと参照渡しの違いを計測し、スライスの再確保とUTF-8バイト表現を確認する |
| **Week 3** | データ集約とカプセル化 | 構造体のメモリ配置を計測し、レシーバとパッケージ可視性を実践する |
| **Week 4** | 抽象化と疎結合設計 | interface + DI + Fake でDB なしのユニットテストを書く |
| **Week 5** | I/O とリソース管理 | `io.Reader/Writer`、`defer`、`context` で外部リソースを安全に扱う |
| **Week 6** | 品質保証と性能計測 | テーブル駆動テストとベンチマークで正しさと性能を数値で証明する |
| **Week 7** | 並行処理の基本原理 | goroutine・channel・select の仕組みを実験で理解する |
| **Week 8** | 高度な並行設計 | Mutex・WaitGroup・Worker Pool を使って実戦的な並行設計を書く |
| **Week 9** | Web プロトコルと通信 | `net/http` と `encoding/json` でミドルウェア付き JSON API を構築する |
| **Week 10** | 永続化と整合性 | `database/sql`・プリペアドステートメント・トランザクションで安全な DB 操作を行う |
| **Week 11** | システムアーキテクチャ | レイヤ分離・エラーラップ・`slog` 構造化ログでアーキテクチャを整える |
| **Week 12** | 配布最適化と運用 | scratch イメージ・グレースフルシャットダウン・`/healthz` で運用可能なアプリを完成させる |

### 各週の到達基準

各週は「書けたか」ではなく「説明できるか」が完了の基準です。

- **Week 1**: なぜ Go から DB に `localhost` ではなくコンテナ名で繋ぐのか
- **Week 2**: 値渡しと参照渡しで、何がコピーされるのか
- **Week 3**: なぜポインタレシーバを選ぶのか
- **Week 4**: なぜ interface を使うと実装差し替えやテストが容易になるのか
- **Week 5**: なぜ巨大ファイルを一括読み込みしない方がよいのか
- **Week 6**: なぜヒープ確保が増えると遅くなりやすいのか
- **Week 7**: unbuffered channel が同期になるのはなぜか
- **Week 8**: なぜ `count++` は安全ではないのか
- **Week 9**: なぜ巨大 JSON はストリームとして扱う方がよいのか
- **Week 10**: なぜトランザクションなしでは整合性が壊れるのか
- **Week 11**: なぜ依存方向を一方向に保つべきなのか
- **Week 12**: なぜ SIGTERM を受けてすぐに kill される設計が危険なのか

### 起動方法（共通）

各ディレクトリに移動して `docker compose up --build` を実行します。

```bash
cd examples/master/01_docker_e2e
docker compose up --build
```

---

## 第3部: 付録（examples/appendix）

本編 12 週を完了した後に取り組む発展モジュールです。順序はありません。

### 01 Fyne GUI

Go によるクロスプラットフォーム GUI アプリケーション開発。

- **対象OS**: Windows、macOS、Linux
- **前提**: CGO と C コンパイラのセットアップが必要
- **内容**: ウィジェット・レイアウト・イベント処理・データバインディング・マルチプラットフォームビルド

### 02 Web Server

`net/http` を使った Go 製 Web サーバの全体像。

- **内容**: ルーティング（Go 1.22 以降のパスパラメータ）・テンプレートレンダリング・`embed` によるファイル同梱・ミドルウェア・Cookie・context によるリクエストスコープ管理・グレースフルシャットダウン
- **実例**: このリポジトリ自体のサーバ実装（`internal/site/`）が参照先

### 03 Crawl Data Manager

Web クロールデータの取得・保存・管理ツールの設計ガイド。

- **内容**: HTTP クライアントによるクロール・HTML パース・並行 Worker Pool・SQLite へのデータ保存・CSV/JSON エクスポート・CLI 設計・レート制限
- **組み合わせる技術**: `net/http`、`golang.org/x/net/html`、`database/sql`、`sync.WaitGroup`、`context`、`encoding/json`、`encoding/csv`、`flag`

---

## ドキュメント一覧

| ドキュメント | 役割 |
|---|---|
| `docs/start-here.md` | 学習の唯一の入口 |
| `docs/go-master-curriculum.md` | 12 週間カリキュラムの正本 |
| `docs/handson-learning-guide.md` | 基礎ハンズオンの学習順ガイド |
| `docs/go-master-curriculum-handson-mapping.md` | 正本と handson の対応表 |
| `docs/go-master-curriculum-gap-list.md` | 整備状況の記録 |
| `docs/go-advanced-syntax-roadmap.md` | 基礎の次に広げる上級トピック |
| `docs/go-concurrency-learning-guide.md` | 並行処理の補助ガイド |
| `docs/gospec-learning-guide.md` | Go 言語仕様書の学習ガイド |

---

## 原則

迷ったとき、常にここに戻ります。

1. 学習順は **正本カリキュラムに従う**
2. `examples/handson` → `examples/master` → `examples/appendix` の順で進める
3. 途中で分からなくなったら `docs/start-here.md` へ戻る
4. 完了の基準は「書けた」ではなく「説明できる」
