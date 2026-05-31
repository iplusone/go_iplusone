# 01 Fyne GUI

Go によるクロスプラットフォーム GUI アプリケーション開発の入口です。

## 付録扱いの理由

本編カリキュラムは Web / API / DB を通じた現代的なシステム構築を主軸にしています。

Fyne はその後に学んだ知識を、ブラウザではなくネイティブ GUI という別の舞台へ応用するモジュールとして位置づけています。

## Fyne とは

[Fyne](https://fyne.io/) は Go 製のクロスプラットフォーム GUI ツールキットです。

- Windows、macOS、Linux で同じコードが動く
- CGO を通じてネイティブの描画 API を呼び出す
- ウィジェット、レイアウト、イベントハンドリングを標準で備える

## 前提知識

Fyne を始める前に次が必要です。

- Go の基礎文法（`examples/handson` 01〜52 相当）
- 構造体とインタフェース（`31_structs_basic` 〜 `43_type_assertions`）
- クロージャ（`26_functions_and_closures`）

## 開発環境の準備

Fyne は CGO を使うため、C コンパイラが必要です。

### Linux（Ubuntu / Debian）

```bash
sudo apt install gcc libgl1-mesa-dev xorg-dev
```

### macOS

Xcode Command Line Tools がインストールされていれば追加不要です。

```bash
xcode-select --install
```

### Windows

[TDM-GCC](https://jmeubank.github.io/tdm-gcc/) または MSYS2 の GCC が必要です。

## プロジェクトの作成

```bash
mkdir fyne-hello && cd fyne-hello
go mod init example.com/fyne-hello
go get fyne.io/fyne/v2
```

## 最小サンプル

```go
package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
)

func main() {
    a := app.New()
    w := a.NewWindow("Hello Fyne")
    w.SetContent(widget.NewLabel("Hello, World!"))
    w.ShowAndRun()
}
```

```bash
go run .
```

## 主要な概念

### App と Window

```
App（アプリケーション全体）
  └── Window（ウィンドウ）
        └── Content（ウィジェットのツリー）
```

`app.New()` でアプリを作り、`NewWindow()` でウィンドウを作り、`SetContent()` でウィジェットをセットします。

`ShowAndRun()` はウィンドウを表示してイベントループに入ります。これ以降の処理はイベント駆動になります。

### ウィジェット

よく使うウィジェットの一覧です。

| ウィジェット | 用途 |
| --- | --- |
| `widget.NewLabel(text)` | テキスト表示 |
| `widget.NewButton(text, func)` | ボタン |
| `widget.NewEntry()` | テキスト入力 |
| `widget.NewSelect(options, func)` | ドロップダウン |
| `widget.NewCheck(text, func)` | チェックボックス |
| `widget.NewProgressBar()` | プログレスバー |

### レイアウト

ウィジェットを並べる方法を `container` パッケージで指定します。

| レイアウト | 説明 |
| --- | --- |
| `container.NewVBox(...)` | 縦並び |
| `container.NewHBox(...)` | 横並び |
| `container.NewGridWithColumns(n, ...)` | グリッド（列数指定） |
| `container.NewBorder(top, bottom, left, right, center)` | 境界レイアウト |
| `container.NewStack(...)` | 重ね合わせ |

### イベントハンドリング

ボタンクリックなどのイベントはコールバック関数で受け取ります。

```go
btn := widget.NewButton("クリック", func() {
    fmt.Println("ボタンが押された")
})
```

UI の更新はメインスレッドで行う必要があります。goroutine から UI を更新するときは `canvas.Refresh(widget)` または `widget.SetText()` などのメソッドを使います。

### データバインディング

`binding` パッケージを使うと、データの変更が自動的に UI に反映されます。

```go
import "fyne.io/fyne/v2/data/binding"

count := binding.NewInt()
count.Set(0)

label := widget.NewLabelWithData(binding.IntToString(count))
btn := widget.NewButton("+1", func() {
    n, _ := count.Get()
    count.Set(n + 1)
})
```

## カウンタアプリのサンプル

```go
package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/data/binding"
    "fyne.io/fyne/v2/widget"
)

func main() {
    a := app.New()
    w := a.NewWindow("カウンタ")

    count := binding.NewInt()
    count.Set(0)

    label := widget.NewLabelWithData(binding.IntToString(count))

    inc := widget.NewButton("+1", func() {
        n, _ := count.Get()
        count.Set(n + 1)
    })
    dec := widget.NewButton("-1", func() {
        n, _ := count.Get()
        count.Set(n - 1)
    })

    w.SetContent(container.NewVBox(label, inc, dec))
    w.ShowAndRun()
}
```

## マルチプラットフォームビルド

Fyne はクロスコンパイルに対応しています。[fyne-cross](https://github.com/fyne-io/fyne-cross) を使うと Docker ベースで他 OS 向けにビルドできます。

```bash
go install github.com/fyne-io/fyne-cross@latest
fyne-cross windows -arch=amd64 .   # Windows 向け .exe を生成
fyne-cross darwin  -arch=amd64 .   # macOS 向けアプリを生成
```

## 発展トピック

- **カスタムウィジェット**: `fyne.Widget` interface を実装することで独自ウィジェットを作れます
- **テーマ**: `app.Settings().SetTheme()` でテーマを切り替えられます
- **通知**: `app.SendNotification()` でシステム通知を送れます
- **ファイルダイアログ**: `dialog.ShowFileOpen()` でファイル選択ダイアログを使えます

## 参考リンク

- [Fyne 公式ドキュメント](https://docs.fyne.io/)
- [Fyne API リファレンス](https://pkg.go.dev/fyne.io/fyne/v2)
- [fyne-io/examples（公式サンプル集）](https://github.com/fyne-io/examples)
