# 02 Web Server

`net/http` を使った Go 製 Web サーバの構築ガイドです。

本編の `examples/master/09_http_and_json`（JSON API）と `examples/master/12_deploy_and_operations`（運用）を土台に、テンプレートレンダリング・静的ファイル配信・ルーティング・ミドルウェアまで含めた Web サーバの全体像を扱います。

## 前提知識

- `net/http` の基本（`examples/handson/60_http_server_basic`）
- 構造体・インタフェース
- `html/template` または `text/template`
- `embed` パッケージ（任意）

## Web サーバの全体構成

```
main.go             エントリポイント・サーバ起動
handler/
  page.go           ページハンドラ
  api.go            JSON API ハンドラ
middleware/
  logging.go        ログ記録ミドルウェア
  recover.go        パニックリカバリミドルウェア
templates/
  layout.html       共通レイアウト
  index.html        トップページ
static/
  style.css         スタイルシート
  app.js            クライアントスクリプト
```

## ルーティング

### 標準ライブラリ（`http.ServeMux`）

Go 1.22 以降、`http.ServeMux` がパスパターンのメソッド指定に対応しました。

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /", indexHandler)
mux.HandleFunc("GET /about", aboutHandler)
mux.HandleFunc("POST /api/items", createItemHandler)
mux.HandleFunc("GET /api/items/{id}", getItemHandler)  // パスパラメータ
```

パスパラメータの取得:

```go
func getItemHandler(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    // ...
}
```

### サードパーティルータ

より複雑なルーティングが必要な場合は `chi` や `gorilla/mux` が使われます。

```go
import "github.com/go-chi/chi/v5"

r := chi.NewRouter()
r.Get("/items/{id}", getItemHandler)
r.Post("/items", createItemHandler)
```

## テンプレートレンダリング

### html/template の基本

```go
import "html/template"

tmpl := template.Must(template.ParseFiles(
    "templates/layout.html",
    "templates/index.html",
))

func indexHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]any{
        "Title": "トップページ",
        "Items": []string{"Go", "Docker", "Linux"},
    }
    tmpl.ExecuteTemplate(w, "layout", data)
}
```

### テンプレートの例

```html
<!-- layout.html -->
<!DOCTYPE html>
<html>
<head><title>{{.Title}}</title></head>
<body>
  {{template "content" .}}
</body>
</html>

<!-- index.html -->
{{define "content"}}
<h1>{{.Title}}</h1>
<ul>
  {{range .Items}}<li>{{.}}</li>{{end}}
</ul>
{{end}}
```

### embed でバイナリに埋め込む

テンプレートと静的ファイルを実行バイナリに含めると、ファイルシステムへの依存がなくなります。

```go
import "embed"

//go:embed templates/*
var templateFiles embed.FS

//go:embed static/*
var staticFiles embed.FS

tmpl := template.Must(template.ParseFS(templateFiles, "templates/*.html"))
http.Handle("/static/", http.FileServerFS(staticFiles))
```

## 静的ファイルの配信

```go
// ファイルシステムから配信
mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

// embed.FS から配信（バイナリに同梱）
mux.Handle("/static/", http.FileServerFS(staticFiles))
```

## ミドルウェアの実装

`http.Handler` をラップする関数がミドルウェアの基本形です。

```go
// ロギングミドルウェア
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        slog.Info("request", "method", r.Method, "path", r.URL.Path, "dur", time.Since(start))
    })
}

// パニックリカバリミドルウェア
func recoverMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if r := recover(); r != nil {
                http.Error(w, "internal server error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}

// 複数のミドルウェアを重ねる
handler := loggingMiddleware(recoverMiddleware(mux))
```

## セッションと認証の入口

### Cookie による簡易セッション

```go
http.SetCookie(w, &http.Cookie{
    Name:     "session",
    Value:    sessionToken,
    HttpOnly: true,   // JavaScript からアクセスできない
    Secure:   true,   // HTTPS のみ
    SameSite: http.SameSiteLaxMode,
})

// 読み取り
cookie, err := r.Cookie("session")
```

### Basic 認証（開発・内部ツール向け）

```go
func basicAuth(next http.Handler, user, pass string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        u, p, ok := r.BasicAuth()
        if !ok || u != user || p != pass {
            w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

## context を使ったリクエストスコープの値の受け渡し

ミドルウェアからハンドラへ値を渡すには `context` を使います。

```go
type contextKey string
const userKey contextKey = "user"

// ミドルウェアで値をセット
ctx := context.WithValue(r.Context(), userKey, "Alice")
next.ServeHTTP(w, r.WithContext(ctx))

// ハンドラで取り出す
user, _ := r.Context().Value(userKey).(string)
```

## グレースフルシャットダウン

本番運用では SIGTERM を受け取ったときにリクエスト処理を完了させてから終了します。詳細は `examples/master/12_deploy_and_operations` を参照してください。

```go
srv := &http.Server{Addr: ":8080", Handler: handler}

go srv.ListenAndServe()

quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
<-quit

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
srv.Shutdown(ctx)
```

## このリポジトリ自体のサーバ実装

このリポジトリのサーバ実装（`internal/site/`）はこのガイドの実例です。

- [internal/site/app.go](../../../internal/site/app.go) — ルーティングとミドルウェアの登録
- [internal/site/data.go](../../../internal/site/data.go) — インメモリデータ
- [internal/site/templates/](../../../internal/site/templates/) — HTML テンプレート
- [internal/site/static/](../../../internal/site/static/) — CSS

## 発展トピック

- **HTTPS**: `srv.ListenAndServeTLS(certFile, keyFile, handler)` でTLS対応
- **HTTP/2**: TLS と合わせると標準で有効になる
- **WebSocket**: `golang.org/x/net/websocket` または `nhooyr.io/websocket`
- **Server-Sent Events**: `text/event-stream` Content-Type で実装できる
- **Rate Limiting**: `golang.org/x/time/rate` パッケージで実装できる
