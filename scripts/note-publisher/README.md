# note-publisher

`note` へ記事下書きを自動投入するための Playwright ベースの小さな投稿ツールです。

前提:

- note の公式投稿 API は前提にしない
- ローカルの Markdown ファイルを正本にする
- 最初は `下書き保存まで自動 / 公開は手動` にする

## できること

- 手動ログインして storage state を保存
- Markdown ファイルからタイトルと本文を読む
- note の新規記事画面を開く
- タイトルを入力する
- 本文を貼り付ける
- 下書き保存を試みる
- スクリーンショットを保存する

## 制限

- note 側 UI 変更に弱いです
- この初版は「Markdown を HTML + plain text として貼り付ける」方式です
- note エディタのペースト挙動によっては、一部の装飾が期待通りに入らない場合があります
- 有料設定や公開処理は入れていません

## セットアップ

```bash
cd scripts/note-publisher
npm install
npx playwright install chromium
```

## 1. ログイン状態を保存

```bash
npm run login
```

ブラウザが開いたら note へ手動ログインしてください。ログイン後、ターミナルへ戻って Enter を押すと `storage/note-auth.json` にセッションが保存されます。

## 2. 下書きを作成

```bash
npm run draft -- --file ../../docs/note-handson-code-articles/01_values_and_types.md
```

オプション:

- `--file <path>`: 投稿する Markdown ファイル
- `--headless`: ヘッドレス実行
- `--publish`: 下書き保存ではなく公開導線も試すための予約フラグ

## 2.5. note貼り付け用に整形する

Markdown をそのまま note に貼ると崩れやすいので、手動投稿前提なら整形済みファイルを作るのがおすすめです。

```bash
npm run build-note-ready
```

1ファイルだけ変換する場合:

```bash
npm run build-note-ready -- --file ../../docs/note-handson-code-articles/01_values_and_types.md
```

出力先:

- `dist/note-ready/*.txt`
  - note に手動で貼りやすいプレーンテキスト版
- `dist/note-ready/*.fragment.html`
  - note に貼るための本文 HTML 断片
- `dist/note-ready/*.html`
  - ブラウザで開いてそのままコピーするための完成版

おすすめの使い方:

1. `dist/note-ready/*.html` をブラウザで開く
2. 表示された記事全体をそのままコピーする
3. note の新規記事画面へ貼り付ける
4. 見出し・コードブロック・箇条書きの入り方を確認する

補足:

- `.txt` は最終手段のプレーンテキスト版です
- 実際に note へ投稿するなら、まず `.html` か `.fragment.html` を使ってください
- `.fragment.html` は自動投稿や独自クリップボード処理向けです

## 記事ファイル形式

front matter は任意です。

```md
---
title: "01 Values and Types: Goの最初のコードで..."
tags:
  - Go
  - Go入門
---

# 01 Values and Types: Goの最初のコードで...

本文...
```

front matter に `title` がなければ、最初の `# 見出し` をタイトルとして使います。

## 環境変数

- `NOTE_BASE_URL`
  - 既定値: `https://note.com`
- `NOTE_NEW_POST_URL`
  - 既定値: `https://note.com/notes/new`
- `NOTE_TITLE_SELECTOR`
  - 既定値: `textarea, input[placeholder*="タイトル"], [contenteditable="true"]`
- `NOTE_EDITOR_SELECTOR`
  - 既定値: `[contenteditable="true"]`
- `NOTE_DRAFT_BUTTON_TEXT`
  - 既定値: `下書き`

UI が変わったときはまずセレクタを調整してください。

## 保存先

- `storage/note-auth.json`
  ログインセッション
- `artifacts/*.png`
  投稿後のスクリーンショット

## 運用のおすすめ

最初は次の流れが安全です。

1. ローカル Markdown を作る
2. ツールで note 下書きへ投入する
3. note 上で見た目を確認する
4. 必要なら微修正する
5. 公開は手動で行う
