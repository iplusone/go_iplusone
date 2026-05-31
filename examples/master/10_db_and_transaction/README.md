# 10 DB And Transaction

第10週の実践編です。`database/sql`、プリペアドステートメント、トランザクションで永続化を扱います。

## 目的

DB 接続を有限リソースとして扱い、安全な CRUD とトランザクションによる整合性保護を実践します。

## 何を確認するか

### 1. database/sql とコネクションプール

`sql.Open` は接続を即時に確立するのではなく、プールを初期化します。

`db.SetMaxOpenConns`、`db.SetMaxIdleConns` でプールサイズを制御します。

### 2. プリペアドステートメント

SQL に値を直接埋め込むのではなく、`?` プレースホルダと引数バインドを使います。

これにより:
- SQL インジェクションを防げる
- 実行計画をキャッシュできる

### 3. トランザクション

複数の更新をひとまとまりにし、失敗したとき全体を巻き戻します。

```go
tx, err := db.Begin()
// ...更新処理...
if err != nil {
    tx.Rollback()
    return
}
tx.Commit()
```

## 構成

```
10_db_and_transaction/
  README.md
  app/
    main.go         DBアクセスのサンプルコード
    go.mod
    Dockerfile
  docker-compose.yml
  initdb/
    001_seed.sql    テーブル作成と初期データ
```

## 起動方法

```bash
cd examples/master/10_db_and_transaction
docker compose up --build
```

## 到達基準

- なぜプリペアドステートメントを使うのか説明できる
- `Begin` / `Commit` / `Rollback` の役割を説明できる
- なぜトランザクションなしでは整合性が壊れるのか説明できる

## 補足観点

- `context` を使った `QueryContext` / `ExecContext` でタイムアウトをつけられます
- `rows.Close()` を `defer` で呼ぶことを忘れないようにしましょう
