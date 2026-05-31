# 43 Type Assertions

このハンズオンでは、空のインタフェースに入っている値を元の型として取り出す `型アサーション` を確認します。

## 学べること

- `value.(int)` や `value.(string)` の形で中身の型を指定できる
- `v, ok := value.(int)` のように書くと安全に確認できる
- 空のインタフェースは何でも入るが、使うときは元の型へ戻す必要がある

## 実行方法

```bash
go run ./examples/handson/43_type_assertions
```
