# 49 Select Basic

複数の `channel` を `select` で扱う基本を確認するためのハンズオンです。

扱う内容:

- 複数 goroutine から別々の channel に値を送る
- `select` で受信可能な channel を選ぶ
- 受信順が起動順とは限らないこと

実行方法:

```bash
go run ./examples/handson/49_select_basic
```
