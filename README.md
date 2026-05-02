# go_iplusone

Go コンテナ化環境の雛形です。

## 起動方法

```bash
docker compose up --build
```

ブラウザで `http://localhost:8080` にアクセスすると、プレースホルダー画面が確認できます。

## サービス

- `web` - Go アプリケーション
- `db` - MySQL 8.0

## 追加作業

- `main.go` を必要なルーティング / ハンドラに拡張
- `internal/` や `pkg/` ディレクトリを追加してアプリ構造を整理
- `docker-compose.yml` に Redis やキャッシュ、CI 用の追加サービスを追加
