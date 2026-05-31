# 59. Benchmark And Alloc

`testing.B` を使ったベンチマークと、アロケーション数の計測を学びます。

- `BenchmarkXxx` 関数で `b.N` 回ループする
- `-benchmem` でメモリ確保量も計測できる
- ヒープ確保が多いと GC の頻度が上がり遅くなる

「遅い」「速い」を感覚ではなく数値で判断するための基本ツールです。

## 実行

```bash
go test -bench=. -benchmem ./examples/handson/59_benchmark_and_alloc/...
```

## 出力の読み方

```
BenchmarkConcat-8   100000   12345 ns/op   4096 B/op   10 allocs/op
```

- `ns/op`: 1回あたりのナノ秒
- `B/op`: 1回あたりのメモリ確保バイト数
- `allocs/op`: 1回あたりのヒープ確保回数

## 学習ポイント

1. ループ内の文字列連結（`+=`）は毎回新しいメモリを確保するため `allocs/op` が多くなる
2. `strings.Builder` を使うとアロケーションを削減できる
3. `b.ResetTimer()` で初期化処理をベンチマークから除外できる
