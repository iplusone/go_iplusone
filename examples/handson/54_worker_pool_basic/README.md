# 54. Worker Pool Basic

固定数のワーカー goroutine でジョブを処理する Worker Pool パターンを学びます。

- `jobs` channel にジョブを送る
- ワーカーは `jobs` から受け取って処理する
- `close(jobs)` でワーカーの `range` ループを終了させる

無制限に goroutine を起動するのではなく、数を固定することでリソースを制御します。

## 実行

```bash
go run ./examples/handson/54_worker_pool_basic
```

## 学習ポイント

1. ワーカー数を固定することで並列度を制御できる
2. `close(jobs)` は送信側が閉じる。受信側が閉じると panic になる
3. `WaitGroup` でワーカー全員の終了を待つ
