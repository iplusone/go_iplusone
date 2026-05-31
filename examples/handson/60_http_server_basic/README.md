# 60. HTTP Server Basic

`net/http` パッケージを使って最小構成の HTTP サーバを書く方法を学びます。

- `http.HandleFunc` でパスとハンドラを登録する
- `http.ListenAndServe` でサーバを起動する
- `r.Method` でメソッドを判定する

Go は標準ライブラリだけで HTTP サーバを書けます。

## 実行

```bash
go run ./examples/handson/60_http_server_basic
```

別ターミナルから確認:

```bash
curl http://localhost:8060/hello
curl http://localhost:8060/echo?msg=test
```

## 学習ポイント

1. `http.ResponseWriter` に書いたものがレスポンスになる
2. `r.URL.Query().Get("key")` でクエリパラメータを取れる
3. ヘッダーを書く前に `WriteHeader` を呼ぶ（書いた後では変えられない）
