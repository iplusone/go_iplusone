# 45 Reflect Slice Checks

このハンズオンでは、`reflect.SliceOf` を使って `[]int` や `[]string` のようなスライス型を判定する方法を確認します。

## 学べること

- `Kind()` だけでは「スライスかどうか」までは分かっても、中の要素型までは区別しにくい
- `reflect.SliceOf(reflect.TypeOf(...))` を使うと `[]int` や `[]string` を比較できる
- 空のインタフェースに入った複雑な値も、型を見て安全に扱える

## 実行方法

```bash
go run ./examples/handson/45_reflect_slice_checks
```
