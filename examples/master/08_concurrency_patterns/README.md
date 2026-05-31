# 08 Concurrency Patterns

第8週の実践編です。Mutex、WaitGroup、Worker Pool で実戦的な並行設計を実装します。

## 目的

複数の goroutine が同じデータを触る場面で、安全に制御するパターンを理解します。

## 何を確認するか

### 1. Mutex による排他制御

複数の goroutine が同じ変数をインクリメントすると、更新が競合して値が壊れます（レースコンディション）。

`sync.Mutex` で `Lock/Unlock` を囲むことで、一度に1つの goroutine だけが実行されます。

### 2. WaitGroup による完了待ち

`sync.WaitGroup` を使うと、複数の goroutine の完了を確実に待てます。

- `wg.Add(n)`: n 件のタスクを追加
- `wg.Done()`: タスク完了を通知（defer で呼ぶ）
- `wg.Wait()`: 全タスクの完了を待つ

### 3. Worker Pool パターン

無制限に goroutine を起動すると、リソースを使い切る場合があります。

Worker Pool では、goroutine 数を固定し、ジョブを channel で配布します。

## 構成

```
08_concurrency_patterns/
  README.md
  main.go
```

## 起動方法

```bash
go run ./examples/master/08_concurrency_patterns
```

競合検出（race detector）を有効にして動かすと問題箇所が見つかります:

```bash
go run -race ./examples/master/08_concurrency_patterns
```

## 到達基準

- なぜ `count++` は安全ではないのか説明できる
- `sync.Mutex` と `sync.WaitGroup` の役割の違いを説明できる
- Worker Pool の基本構造を書ける

## 補足観点

- `sync/atomic` を使うと `Mutex` なしで単純なカウンタを安全に操作できます
- `sync.Once` を使うと初期化処理を一度だけ実行できます
