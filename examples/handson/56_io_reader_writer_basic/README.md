# 56. I/O Reader Writer Basic

`io.Reader` と `io.Writer` の基本的な使い方を学びます。

- `io.Reader` はバイト列を読み込むインタフェース
- `io.Writer` はバイト列を書き出すインタフェース
- `io.Copy` でリーダーからライターへ転送できる

ファイル、HTTP ボディ、バッファなどが同じインタフェースを実装しているため、組み合わせて使えます。

## 実行

```bash
go run ./examples/handson/56_io_reader_writer_basic
```

## 学習ポイント

1. `strings.NewReader` は文字列を `io.Reader` として扱える
2. `bytes.Buffer` は `io.Writer` と `io.Reader` の両方を実装する
3. `io.Copy` で転送することで、全体をメモリに保持せずに処理できる
