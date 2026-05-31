# 63. Web To Markdown Basic

Web ページを取得し、HTML を Markdown に変換する方法を学びます。

`51_file_access_basic`（ファイルアクセス）、`52_network_access_basic`（ネットワークアクセス）、`62_database_sql_basic`（データベースアクセス）の次に位置するハンズオンです。

## 使用ライブラリ

[JohannesKaufmann/html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) — HTML を Markdown テキストに変換する Go ライブラリ

## 実行

```bash
cd examples/handson/63_web_to_markdown_basic
go run .
```

## セットアップ

```bash
cd examples/handson/63_web_to_markdown_basic
go mod tidy
```

## 学習ポイント

1. `net/http` で取得した `Response.Body` は `io.Reader` — そのままライブラリに渡せる
2. `html-to-markdown` は `ConvertString` と `ConvertReader` の両方を持つ
3. このリポジトリでは `httptest` のローカルサーバで動作確認する
4. 変換後の Markdown を `os.WriteFile` でファイルに保存する流れは `51` で学んだ知識の応用
