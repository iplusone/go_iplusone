# 51 Data Access File Access Basic

データアクセスの入口として、ファイルアクセスの基本を確認するためのハンズオンです。

扱う内容:

- `os.ReadFile` でファイル全体を読む
- `os.WriteFile` でファイル全体を書き出す
- `[]byte` と `string` の変換
- `os.OpenFile` で追記モードのファイルを開く
- `defer` で `Close` をまとめる

この回で意識したいこと:

- ファイルアクセスは OS リソースを扱うデータアクセスである
- 簡単な入出力は `ReadFile` / `WriteFile` で始められる
- 柔軟な制御が必要な時は `OpenFile` と `*os.File` を使う
- 開いたファイルは必ず `Close` し、そのために `defer` を使う

実行方法:

```bash
go run ./examples/handson/51_file_access_basic
```
