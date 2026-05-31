# 06 Testing And Benchmark

第6週の実践編です。テーブル駆動テスト、ベンチマーク、エスケープ解析で正しさと性能を証明します。

## 目的

「動いているはず」ではなく「動くことを証明できる」状態にします。

## 何を確認するか

### 1. テーブル駆動テスト

入力と期待値の組をテーブルとして並べ、ループで一気にテストします。

- 境界値を網羅しやすい
- ケースを追加するときにコードが増えない

### 2. ベンチマーク

`BenchmarkXxx` 関数に `b.N` 回ループを書くことで、処理時間とアロケーション数を計測できます。

```
go test -bench=. -benchmem ./examples/master/06_testing_and_benchmark/...
```

### 3. エスケープ解析

```
go build -gcflags=-m ./examples/master/06_testing_and_benchmark/...
```

変数がスタックに置かれるかヒープに逃げるかを確認できます。ヒープ確保が多いと GC の負荷が上がります。

## 構成

```
06_testing_and_benchmark/
  README.md
  calc.go         テスト対象の関数
  calc_test.go    テーブル駆動テストとベンチマーク
```

## 起動方法

テスト実行:

```bash
go test ./examples/master/06_testing_and_benchmark/...
```

ベンチマーク実行:

```bash
go test -bench=. -benchmem ./examples/master/06_testing_and_benchmark/...
```

## 到達基準

- テーブル駆動テストを書ける
- `-benchmem` の出力（ns/op、B/op、allocs/op）の意味を説明できる
- なぜヒープ確保が増えると遅くなりやすいのか説明できる

## 補足観点

- `t.Run` を使うとサブテストに名前がついてエラー特定が楽になります
- `-gcflags=-m` は `go build` だけでなく `go run` にも使えます
