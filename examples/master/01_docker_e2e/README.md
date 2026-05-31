# 01 Docker E2E

第1週の実践編として、`front` `app` `db` の3層を Docker Compose で起動し、ブラウザから Go バックエンド経由で DB のデータを表示するためのサンプルです。

このサンプルの目的は、単に Go を動かすことではなく、次の流れを最小構成で確認することです。

1. Docker Compose で複数サービスを起動する
2. Go コンテナから DB コンテナへ接続する
3. ブラウザから Front コンテナへアクセスする
4. Front の JavaScript から Go API を呼ぶ
5. Go が DB から取得したデータをブラウザに表示する

## ディレクトリ構成

- `docker-compose.yml`
  3サービスを束ねる Compose 定義
- `app/`
  Go バックエンド
- `front/`
  静的 HTML と JavaScript
- `initdb/`
  DB 初期化 SQL

## 起動手順

このディレクトリで次を実行します。

```bash
docker compose up --build
```

起動後、ブラウザで次へアクセスします。

```text
http://localhost:18080
```

`Load message from DB` ボタンを押すと、DB の `messages` テーブルにあるデータが画面に表示されます。

## サービス構成

- `front`
  `http://localhost:18080`
- `app`
  `http://localhost:18081`
- `db`
  MySQL 8.0

## このサンプルで確認したいこと

### 1. `localhost` ではなくサービス名で繋ぐ

Go バックエンドは DB に対して `localhost` ではなく `db:3306` で接続します。

理由:

- `localhost` は `app` コンテナ自身を指す
- DB は別コンテナなので、Compose ネットワーク上のサービス名 `db` で名前解決する必要がある

### 2. ブラウザから見る `localhost` と、コンテナ内の `localhost` は別物

ブラウザから `http://localhost:18080` にアクセスできるのは、ホスト側にポート公開しているからです。

一方で、`app` コンテナの中で `localhost:18081` と書いても、それは `app` コンテナ自身を指します。

### 3. Front と App は HTTP で会話する

`front/index.html` では JavaScript の `fetch()` を使って、`app` の `/api/message` を呼びます。

つまりブラウザは:

- DB を直接知らない
- Go の API にだけ話しかける

という構造です。

### 4. DB の初期データは SQL で入る

`initdb/001_seed.sql` によって、起動時にテーブル作成と初期データ投入が行われます。

これにより、環境を捨てて再作成しても、同じ初期状態を再現できます。

## 到達基準

このサンプルを動かしたあと、最低限次の説明ができれば第1週の到達点として十分です。

- なぜ Go から DB に `db:3306` で繋ぐのか
- なぜブラウザから `localhost:18080` で Front が見えるのか
- なぜブラウザは DB ではなく Go API を叩くのか
- なぜ SQL 初期化ファイルがあると再現性が上がるのか

## 次に見るとよい観点

- `docker compose ps` でコンテナが分かれていることを確認する
- `docker compose logs app` で Go がどのホストへ繋いでいるかを見る
- `docker compose exec db mysql -uroot -psecret e2e_app -e "select * from messages"` で DB の中身を確認する
