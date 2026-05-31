# 53. WaitGroup Basic

`sync.WaitGroup` を使って複数の goroutine の完了を待つ方法を学びます。

- `wg.Add(n)` で待つ件数を登録する
- goroutine 内で `defer wg.Done()` を呼ぶ
- `wg.Wait()` で全完了を待つ

`go` キーワードで goroutine を起動しても、`main` 関数は終了待ちをしてくれません。WaitGroup がその同期を担います。

## 実行

```bash
go run ./examples/handson/53_waitgroup_basic
```

## 学習ポイント

1. `wg.Done()` を `defer` で呼ぶことで、goroutine がパニックしても確実に実行される
2. `wg.Add` は goroutine を起動する前に呼ぶ（起動後に呼ぶと競合する）
3. `WaitGroup` は goroutine 数が動的に決まる場合に特に有効
