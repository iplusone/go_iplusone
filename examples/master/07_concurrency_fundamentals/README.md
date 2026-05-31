# 07 Concurrency Fundamentals

第7週の実践編です。goroutine、channel、select の基本原理を実験します。

## 目的

Go の並行処理を支える仕組みを理解し、データ受け渡しと同期を channel で表現できるようにします。

## 何を確認するか

### 1. goroutine の軽量さ

Go のスレッドは OS スレッドではなく goroutine です。

goroutine はスタックが小さく（初期 2KB）、数千〜数万を起動しても問題ありません。ランタイムが OS スレッドへのスケジューリングを担います。

### 2. channel による同期

unbuffered channel は送受信が同時に揃うまで双方がブロックします。これにより、データ受け渡しと同期がひとつの機構で実現されます。

### 3. select による多重化

複数の channel を同時に待つには `select` を使います。

どれか準備できた channel の case が実行されます。

## 構成

```
07_concurrency_fundamentals/
  README.md
  main.go
```

## 起動方法

```bash
go run ./examples/master/07_concurrency_fundamentals
```

## 到達基準

- unbuffered channel が同期になるのはなぜか説明できる
- `select` で複数 channel を待つコードを書ける
- goroutine が OS スレッドと何が違うかを説明できる

## 補足観点

- `go run -race` を付けるとデータ競合を検出できます
- `runtime.GOMAXPROCS` で並列度を変えて実験できます
