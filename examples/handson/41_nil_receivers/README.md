# 41 Nil Receivers

このハンズオンでは、ポインタレシーバが `nil` のときにどう振る舞うかを確認します。

## 学べること

- `var item *MyData` のように宣言だけしたポインタは `nil`
- `nil` の状態でもメソッド呼び出し自体は書ける
- ただし、メソッド内で `nil` チェックをしないと安全に扱えない

## 実行方法

```bash
go run ./examples/handson/41_nil_receivers
```
