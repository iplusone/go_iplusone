# 55. Race Detector And Atomic

データ競合（レースコンディション）の発生と、`sync/atomic` による安全な操作を学びます。

- `go run -race` でレース検出器を有効にする
- `count++` は「読む・足す・書く」の3ステップで分割可能なため競合する
- `atomic.AddInt64` は不可分操作なので競合しない

## 実行

通常実行（競合が起きていても表面上は見えにくい）:

```bash
go run ./examples/handson/55_race_detector_and_atomic
```

レース検出器を有効にして実行:

```bash
go run -race ./examples/handson/55_race_detector_and_atomic
```

## 学習ポイント

1. `count++` は安全ではない。複数 goroutine から同時に触ると値が壊れる
2. `-race` フラグでレース検出器が有効になり、競合を検出してくれる
3. `atomic.AddInt64` は1ステップで完了する不可分操作（原子的操作）
