# 12 Switch Default Fallthrough

`switch` の `default` と `fallthrough` を確認するためのハンズオンです。

扱う内容:

- `default`
- 通常の `switch` の終了
- `fallthrough`
- `fallthrough` が次の条件を再評価しないこと

## fallthrough とは

Go の `switch` は、通常は一致した `case` の処理を実行したらそこで終了します。

`fallthrough` を書くと、現在の `case` の次にある `case` の条件を再判定せず、そのまま次の処理へ進みます。

```go
score := 85

switch {
case score >= 80:
    fmt.Println("80点以上です")
    fallthrough
case score >= 60:
    fmt.Println("60点以上です")
}
```

この例では `score >= 60` をもう一度判定しているのではなく、`fallthrough` によって次の `case` の本体へ進んでいます。

## 注意点

- `fallthrough` は次の `case` にしか進めません
- 次の `case` の条件式は評価されません
- 多用すると流れが読みにくくなるため、必要な場面だけで使うのが安全です

実行方法:

```bash
go run ./examples/handson/12_switch_default_fallthrough
```
