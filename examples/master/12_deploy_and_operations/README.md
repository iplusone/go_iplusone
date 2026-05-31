# 12 Deploy And Operations

第12週の実践編です。scratch イメージ、グレースフルシャットダウン、ヘルスチェックで運用可能なアプリを作ります。

## 目的

「動く」だけでなく「止め方まで設計された」プロセスにします。

## 何を確認するか

### 1. 極小コンテナイメージ

`CGO_ENABLED=0` で静的リンクし、`scratch` ベースイメージを使うことで、実行バイナリだけを含む最小イメージを作れます。

- セキュリティリスクを減らす（シェルなし、パッケージなし）
- イメージサイズを数 MB 以下に抑えられる

### 2. グレースフルシャットダウン

`os/signal` で SIGTERM を受け取り、処理中のリクエストが終わるのを待ってから終了します。

Kubernetes や Docker はコンテナ停止時に SIGTERM を送ります。これを無視してすぐに終了すると、処理中のリクエストが切断されます。

### 3. ヘルスチェック API

`/healthz` エンドポイントを持つことで、オーケストレータ（Kubernetes 等）が生存確認できます。

## 構成

```
12_deploy_and_operations/
  README.md
  main.go
  Dockerfile
```

## 起動方法

ローカルで直接起動:

```bash
go run ./examples/master/12_deploy_and_operations
```

Docker でビルドして起動:

```bash
cd examples/master/12_deploy_and_operations
docker build -t deploy-ops .
docker run --rm -p 8012:8012 deploy-ops
```

ヘルスチェック確認:

```bash
curl http://localhost:8012/healthz
```

グレースフルシャットダウン確認:

```bash
# サーバ起動後、別ターミナルから
curl http://localhost:8012/slow  # 5秒かかるリクエストを送る
# すぐに Ctrl+C でシャットダウン → リクエストが完了してから終了することを確認
```

## 到達基準

- `scratch` ベースイメージと静的リンクの意味を説明できる
- SIGTERM を受けてグレースフルシャットダウンするコードを書ける
- なぜ SIGTERM を受けてすぐに kill される設計が危険なのか説明できる

## 補足観点

- `docker image inspect` でイメージサイズを確認できます
- Kubernetes の `readinessProbe` / `livenessProbe` はこの `/healthz` を叩きます
