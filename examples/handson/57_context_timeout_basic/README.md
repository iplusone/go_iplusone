# 57. Context Timeout Basic

`context.WithTimeout` を使ってタイムアウト付きの処理を書く方法を学びます。

- `context.WithTimeout` で期限つきのコンテキストを作る
- `defer cancel()` で必ずリソースを解放する
- `ctx.Done()` チャネルで期限切れを検知する

処理が終わらない goroutine を作らないために、context によるタイムアウト制御は重要です。

## 実行

```bash
go run ./examples/handson/57_context_timeout_basic
```

## 学習ポイント

1. `cancel` は必ず `defer` で呼ぶ（呼ばないとリソースリークする）
2. `ctx.Done()` は期限切れ・キャンセルの両方で閉じられる
3. `ctx.Err()` でなぜ終了したかを確認できる（`context.DeadlineExceeded` / `context.Canceled`）
