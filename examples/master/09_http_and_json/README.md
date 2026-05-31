# 09 HTTP And JSON

第9週の実践編です。`net/http`、ミドルウェア、`encoding/json` で JSON API を構築します。

## 目的

HTTP をブラックボックスにせず、リクエスト・レスポンスの構造を理解した上で JSON API を実装します。

## 何を確認するか

### 1. net/http の最小 HTTP サーバ

Go の標準ライブラリだけで HTTP サーバを立てられます。

`http.HandleFunc` または `http.Handler` interface を実装するだけです。

### 2. ミドルウェアパターン

`http.Handler` をラップする関数を作ることで、ログ出力や認証チェックをルーティングと分離できます。

### 3. JSON のシリアライズ・デシリアライズ

`encoding/json` の `Encode` / `Decode` を使います。構造体タグ（`json:"field_name"`）でフィールド名を制御できます。

巨大な JSON はストリームとして `json.NewDecoder` で処理するとメモリ効率が上がります。

## 構成

```
09_http_and_json/
  README.md
  main.go       HTTPサーバのエントリポイント
```

## 起動方法

```bash
go run ./examples/master/09_http_and_json
```

別ターミナルから確認:

```bash
curl http://localhost:8099/api/items
curl -X POST http://localhost:8099/api/items -H "Content-Type: application/json" -d '{"name":"book"}'
```

## 到達基準

- `net/http` で最小の JSON API サーバを書ける
- ミドルウェアでリクエストログを記録できる
- なぜ巨大 JSON はストリームとして扱う方がよいのか説明できる

## 補足観点

- `http.StatusMethodNotAllowed` など、適切なステータスコードを返すことが重要です
- `json.Decoder` は `io.Reader` から直接デコードするため、レスポンスボディを全部バッファに読まなくてよいです
