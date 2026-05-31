# TOPページ起点の公開構成メモ

このドキュメントは、このリポジトリで実装している TOP ページ `/` と、そこから直接たどれる公開導線を整理した構成メモです。

本家 `iplusone_wordpress_laravel_admin` の `docs/top-page-structure.md` を参照しつつ、Go アプリ側の現行実装に合わせて記載しています。

## 対象範囲

- TOPページ `/`
- TOPページ内アンカーリンク
- TOPページのヘッダー、本文、フッターから直接遷移できる内部ページ
- TOPページ内カードから直接遷移できる詳細ページ種別
- TOPページから直接出ていく外部URL

対象外:

- 2クリック目以降の導線網羅
- 管理画面
- API
- 未実装のデータ永続化詳細

## 入口

- URL: `/`
- 起動エントリ: [main.go](/home/ishii/projects/go_iplusone/main.go:1)
- ルーティング / ハンドラ: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:1)
- テンプレート: [internal/site/templates/home.html](/home/ishii/projects/go_iplusone/internal/site/templates/home.html:1)

## TOPページの描画部品

- 共通 head / header / footer: [internal/site/templates/layout.html](/home/ishii/projects/go_iplusone/internal/site/templates/layout.html:1)
- TOP本体: [internal/site/templates/home.html](/home/ishii/projects/go_iplusone/internal/site/templates/home.html:1)
- スタイル: [internal/site/static/styles.css](/home/ishii/projects/go_iplusone/internal/site/static/styles.css:1)
- 初期データ: [internal/site/data.go](/home/ishii/projects/go_iplusone/internal/site/data.go:1)

## 設計上の共通部分

### ルーティングと画面生成

- 入口は [main.go](/home/ishii/projects/go_iplusone/main.go:1) から [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:1) の `site.New()` に集約
- URL 配線は `App.Handler()` に集約し、`http.ServeMux` で公開ルートを一元管理
- 固定ページは `page(key)` ハンドラで共通化し、`content` マップ上の定義差し替えで増やせる構成
- 一覧 / 詳細ページも `render()` を共通で通し、`ViewData` を介してテンプレートへ値を渡す

### テンプレート共通部品

- head / header / footer は [internal/site/templates/layout.html](/home/ishii/projects/go_iplusone/internal/site/templates/layout.html:1) に分離
- 各ページテンプレートは共通部品を `template` 呼び出しで取り込み、本文だけを持つ構成
- 見た目の共通ルールは [internal/site/static/styles.css](/home/ishii/projects/go_iplusone/internal/site/static/styles.css:1) に集約

### ナビゲーション

- ヘッダーナビは `primaryNav()`
- フッターナビは `footerNav()`
- 外部導線は `seedExternalLinks()`

いずれも [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:1) または [internal/site/data.go](/home/ishii/projects/go_iplusone/internal/site/data.go:1) 側で定義し、個別テンプレートにベタ書きしない方針

### データ構造

- TOP、固定ページ、NEWS、SUNO はそれぞれ `HomePageData`、`SimplePage`、`NewsItem`、`Song` の型で整理
- 画面表示用の初期データは [internal/site/data.go](/home/ishii/projects/go_iplusone/internal/site/data.go:1) に集約
- 現段階はインメモリだが、将来的に DB 化する際もハンドラ側の責務を大きく変えずに差し替える想定

### 静的ファイル配信

- CSS は `embed` でバイナリに同梱
- `/static/` 配下は `mustSub("static")` を通じて配信
- デプロイ時にテンプレートや CSS の外部配置を前提にしない最小構成

## TOPページ内セクション

### 同一ページ内アンカー

- `/#about`
  - 定義: [internal/site/data.go](/home/ishii/projects/go_iplusone/internal/site/data.go:111)
- `/#service`
  - 定義: [internal/site/data.go](/home/ishii/projects/go_iplusone/internal/site/data.go:116)

### 主な表示ブロック

- ヒーロー領域
- 7つの基盤PR
- GoSpec学習カリキュラム
- 最新のお知らせ
- AI-DDD説明ブロック
- SUNOカード
- 外部サービス導線

主な描画定義:

- [internal/site/templates/home.html](/home/ishii/projects/go_iplusone/internal/site/templates/home.html:1)

## TOPから直接遷移できる内部ページ

### 固定ページ / 一覧

- `/news`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:196)
  - テンプレート: [internal/site/templates/news_list.html](/home/ishii/projects/go_iplusone/internal/site/templates/news_list.html:1)

- `/gospec/{chapter-slug}`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:1)
  - テンプレート: [internal/site/templates/study_chapter.html](/home/ishii/projects/go_iplusone/internal/site/templates/study_chapter.html:1)

- `/blog`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:129)

- `/contact`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:130)

- `/hp_templates`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:131)

- `/project-works`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:132)

- `/corporates`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:133)

- `/map`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:134)

- `/homepages`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:135)

- `/railway_routes/front`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:136)

- `/suno-songs`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:228)
  - テンプレート: [internal/site/templates/song_list.html](/home/ishii/projects/go_iplusone/internal/site/templates/song_list.html:1)

- `/suno-library`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:128)

- `/platforms/bidding`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:137)

- `/platforms/reservation`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:138)

### 稼働確認

- `/healthz`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:172)

## TOPから直接遷移できる詳細ページ種別

### 動的詳細

- `/news/detail/{id}`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:216)
  - データ: [internal/site/data.go](/home/ishii/projects/go_iplusone/internal/site/data.go:142)

- `/suno-songs/{id}`
  - 実装: [internal/site/app.go](/home/ishii/projects/go_iplusone/internal/site/app.go:248)
  - データ: [internal/site/data.go](/home/ishii/projects/go_iplusone/internal/site/data.go:174)

## TOPから直接出ていく外部URL

- `https://ec-plus1.iplusone.co.jp/`
- `https://estimates.iplusone.co.jp/`
- `https://spottown.iplusone.co.jp/`
- `https://jobs.iplusone.co.jp/`
- `https://jobs.iplusone.co.jp/platform`
- `https://real-estate-website.iplusone.co.jp/`
- `https://iplusone.co.jp/match3/`

定義元:

- [internal/site/data.go](/home/ishii/projects/go_iplusone/internal/site/data.go:131)

## データ方針

- 現段階の NEWS / SUNO / TOP カードはインメモリ初期データで保持
- データ定義は [internal/site/data.go](/home/ishii/projects/go_iplusone/internal/site/data.go:1) に集約
- 将来的な DB 差し替えを見据えて、ハンドラからはデータ取得関数を通して参照

## 簡易サイトマップ

```text
/
|- /#about
|- /#service
|- /news
|  `- /news/detail/{id}
|- /blog
|- /contact
|- /hp_templates
|- /project-works
|- /corporates
|- /map
|- /homepages
|- /railway_routes/front
|- /suno-songs
|  `- /suno-songs/{id}
|- /suno-library
|- /platforms/bidding
|- /platforms/reservation
|- /healthz
|- https://ec-plus1.iplusone.co.jp/
|- https://estimates.iplusone.co.jp/
|- https://spottown.iplusone.co.jp/
|- https://jobs.iplusone.co.jp/
|- https://jobs.iplusone.co.jp/platform
|- https://real-estate-website.iplusone.co.jp/
`- https://iplusone.co.jp/match3/
```

## 補足

- この構成書は「TOPから1クリックで到達できる範囲」を中心に整理しています。
- 詳細ページは件数可変のため URL パターンとして記載しています。
- `ABOUT` と `SERVICE` は別ページではなく TOP 内アンカーです。
