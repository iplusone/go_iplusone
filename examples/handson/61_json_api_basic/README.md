# 61. JSON API Basic

`encoding/json` を使って JSON を扱う基本を学びます。

- `json.Marshal` で構造体を JSON バイト列に変換する
- `json.Unmarshal` で JSON バイト列を構造体に変換する
- 構造体タグ（`json:"field_name"`）でフィールド名を制御する

## 実行

```bash
go run ./examples/handson/61_json_api_basic
```

## 学習ポイント

1. 構造体タグ `json:"name,omitempty"` で空の場合にフィールドを省略できる
2. `json.NewEncoder(w).Encode(v)` は HTTP レスポンスに直接書ける（バッファ不要）
3. `json.NewDecoder(r.Body).Decode(&v)` はリクエストボディを直接デコードできる
