# 11 Architecture Refactor

第11週の実践編です。レイヤ分離、エラー戦略、構造化ログでシステムアーキテクチャを整えます。

## 目的

変更が伝播しにくい構造を作り、問題発生時に追跡できる情報を残せるようにします。

## 何を確認するか

### 1. レイヤの分離

`cmd/`、`internal/`、ドメイン / ユースケース / インフラの層を分けることで、変更の影響範囲を限定できます。

- `cmd/`: エントリポイント。フレームワーク固有の初期化
- `internal/domain/`: ビジネスロジック。外部依存なし
- `internal/usecase/`: ユースケース。ドメインに依存
- `internal/infra/`: DB・HTTP などの実装

依存の方向は外側から内側へ一方向に保ちます。

### 2. エラー戦略

`fmt.Errorf("...: %w", err)` でエラーに文脈を付けてラップします。

`errors.Is(err, target)` でラップされたエラーを種類で判定できます。

### 3. 構造化ログ

`log/slog` を使うと、JSON 形式のログを出力できます。

```go
slog.Info("request received", "method", r.Method, "path", r.URL.Path)
```

ログを構造化することで、後から検索・集計しやすくなります。

## 構成

```
11_architecture_refactor/
  README.md
  main.go
  internal/
    domain/
      item.go       ドメインモデル
    usecase/
      item_usecase.go
    infra/
      memory_repo.go
```

## 起動方法

```bash
go run ./examples/master/11_architecture_refactor
```

## 到達基準

- なぜ依存方向を一方向に保つべきなのか説明できる
- `%w` でエラーをラップして `errors.Is` で判定できる
- `slog` で JSON ログを出力できる

## 補足観点

- `slog.SetDefault` でデフォルトロガーを JSON ハンドラに差し替えられます
- `errors.As` を使うと特定の型のエラーを取り出せます
