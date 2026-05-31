# 04 Abstraction And DI

第4週の実践編です。interface による抽象化と依存性注入で、テスト容易で変更に強い設計を実践します。

## 目的

上位ロジックを具体的な実装から切り離す設計パターンを理解します。

- interface で振る舞いを定義し、実装を差し替え可能にする
- コンストラクタ注入で依存を外から渡す
- DB なしでユニットテストを書けるようにする

## 何を確認するか

### 1. interface による抽象化

Go の interface は暗黙的に実装されます。

型が interface のメソッドセットを満たしていれば、その interface として扱えます。呼び出し元は実装の型を知る必要がありません。

### 2. 依存性注入（DI）

上位ロジックが直接 `DB` や外部サービスを生成するのではなく、interface として受け取る形にします。

これにより:
- 本番では実 DB を渡す
- テストでは Fake を渡す

という切り替えが、ロジックを変えずにできます。

### 3. Fake によるテスト

外部依存を Fake（テスト用実装）に差し替えることで、DB なしでロジックをテストできます。

## 構成

```
04_abstraction_and_di/
  README.md
  main.go           実験コードのエントリポイント
  store/
    store.go        UserStore interface の定義
    memory.go       メモリ上の Fake 実装
  service/
    service.go      UserStore に依存するビジネスロジック
```

## 起動方法

```bash
go run ./examples/master/04_abstraction_and_di
```

## 到達基準

- なぜ interface を使うと実装差し替えやテストが容易になるのか説明できる
- コンストラクタ注入の形を書ける
- Fake を使ったテストの構造を説明できる

## 補足観点

- `errors.Is` / `errors.As` を使うと、エラーの種類を型で判定できます（第11週で詳しく扱います）
- interface はできるだけ小さくするのが Go の慣習です（1〜2メソッドが多い）
