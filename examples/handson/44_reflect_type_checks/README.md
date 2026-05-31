# 44 Reflect Type Checks

このハンズオンでは、`reflect.TypeOf` と `Kind` を使って、空のインタフェースに入ってきた値の型を確認する書き方を学びます。

## 学べること

- `reflect.TypeOf(value)` で値の型情報を取得できる
- `.Kind()` を使うと `int` や `string` のような基本的な種類を比較できる
- 型が合うときだけ型アサーションして代入する、という安全寄りの流れが分かる

## 実行方法

```bash
go run ./examples/handson/44_reflect_type_checks
```
