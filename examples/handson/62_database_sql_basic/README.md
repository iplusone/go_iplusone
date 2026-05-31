# 62. Database SQL Basic

`database/sql` パッケージを使ってデータベースにアクセスする基本を学びます。

- `sql.Open` で DB に接続する
- `db.QueryRow` / `db.Query` でデータを取得する
- `db.Exec` でデータを更新・挿入する
- プリペアドステートメント（`?` プレースホルダ）で SQL インジェクションを防ぐ

このハンズオンは SQLite を使ってローカルで動かします。

## 実行

```bash
go run ./examples/handson/62_database_sql_basic
```

## 学習ポイント

1. `rows.Close()` は必ず `defer` で呼ぶ（呼ばないと接続が解放されない）
2. `?` プレースホルダに値を渡すことで、SQL インジェクションを防げる
3. `db.Ping()` で接続確認ができる
