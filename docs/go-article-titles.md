# Go言語コンテンツ 記事一覧タイトル案

## Go言語仕様

- Go言語仕様書（GoSpec）入門: まず何を読むべきか
- EBNFで読むGoSpec: 構文ルールの見方
- トークンとセミコロン自動挿入: Goの字句解析を理解する
- Goの型システム入門: 強い型付けの基本
- 基底型（Underlying Type）とは何か
- 代入可能性（Assignability）を仕様から理解する
- 型定義と型エイリアスの違い
- 型なし定数（Untyped Constants）の仕組み
- Goのゼロ値を正しく理解する
- スコープとシャドーイングの落とし穴
- 構造体と埋め込みの仕様を整理する
- インターフェースの仕組みと動的型
- メソッド集合をGoSpecベースで理解する
- ポインタとアドレス可能性（Addressability）の基礎
- スライスの内部構造と仕様上の振る舞い
- マップの仕様と注意点
- 文字列とruneの違いを仕様から学ぶ
- 演算子の優先順位と評価順序
- オーバーフローと整数演算の仕様
- defer / panic / recover の正しい理解
- goroutineの仕様上の意味とは
- channelの送受信とcloseの挙動を整理する
- select文の仕組みと使いどころ
- パッケージ初期化順序とinit関数の役割

## Go開発基礎

- Go開発環境の基本セットアップ
- go mod入門: モジュール管理の基本
- go workの使いどころ
- Goの標準的なディレクトリ構成
- internalパッケージの活用方法
- cmdディレクトリ構成の設計パターン
- import循環を避ける設計の考え方
- Goにおけるエラーハンドリングの基本
- errors.Is / errors.As / %w の使い分け
- context.Contextを正しく使う
- table-driven testの書き方
- Goのベンチマーク測定入門
- fuzz testを実務でどう使うか
- 標準ライブラリで始めるHTTPサーバー実装
- encoding/jsonの基本と注意点
- database/sqlの基本パターン
- syncパッケージの主要機能まとめ
- timeパッケージで失敗しないための基礎知識

## 実践設計・実装

- Goでレイヤードアーキテクチャを組む
- Goとクリーンアーキテクチャの相性
- interfaceを切りすぎない設計
- Web APIのエラーレスポンス設計
- バリデーション設計の基本方針
- RepositoryパターンはGoでどう使うか
- トランザクション管理の実装パターン
- worker pool実装の基本
- fan-out / fan-inパターン入門
- pipeline処理をGoで組む方法
- errgroupで並行処理を安全に扱う
- graceful shutdownの実装手順
- retry / backoffの設計ポイント
- pprofで性能を計測する基本
- appendとメモリアロケーションの関係
- GCを意識した実装の考え方

## Goとクラウド・周辺技術

- GoでGoogle Cloudを扱うときの設計原則
- Go Client Librariesの選び方
- GoでCloud Runアプリを設計する
- GoでCloud Storageを扱う基本
- GoでPub/Subを使う実装パターン
- GoでBigQueryを扱うときのポイント
- GoでSecret Managerを安全に使う
- GoでVertex AIを扱うときの考え方
- GoアプリのDockerfile最適化
- マルチステージビルドの基本
- GitHub ActionsでGoのCIを組む
- GoとgRPCの基本実装
- protobufとGoコード生成の基礎
- OpenAPIとGo API実装のつなぎ方

## 学習・資料系

- Go初学者が最初に読むべき公式ドキュメント
- Effective Goで学ぶべきポイント
- GoSpecを読む習慣をどう作るか
- Goコードレビューでよく見る指摘集
- Goでよくあるバグとその防ぎ方
- Go用語集: 仕様書を読むためのキーワード
- Go学習ロードマップ: 初級から中級へ
