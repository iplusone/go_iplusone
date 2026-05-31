# 05 I/O And Resource Lifecycle

第5週の実践編です。`io.Reader/Writer`、`defer`、`context` を使って外部リソースを安全に扱います。

## 目的

ファイルやネットワークなどの外部リソースは「借りたら必ず返す」設計が必要です。この週では、そのパターンを実践します。

## 何を確認するか

### 1. ストリーム処理

大きなデータを一括で読み込むと、それだけメモリを消費します。

`io.Reader` / `io.Writer` を使うと、一定サイズのバッファで処理し続けることができます。

### 2. defer による確実な後始末

`os.Open` で開いたファイルは必ず閉じる必要があります。

`defer f.Close()` を開いた直後に書くことで、関数を抜けるときに必ず閉じられます。

### 3. context によるタイムアウト

`context.WithTimeout` を使うと、指定時間を超えた処理をキャンセルできます。

これにより「終わらない処理」を作らない設計ができます。

## 構成

```
05_io_and_resource_lifecycle/
  README.md
  main.go
```

## 起動方法

```bash
go run ./examples/master/05_io_and_resource_lifecycle
```

## 到達基準

- なぜ巨大ファイルを一括読み込みしない方がよいのか説明できる
- `defer` を使った安全なリソース解放を書ける
- `context.WithTimeout` によるタイムアウト処理を書ける

## 補足観点

- `bufio.Scanner` はテキストを行単位で読むときに便利です
- `io.Copy` はリーダーからライターへデータを転送するシンプルな方法です
