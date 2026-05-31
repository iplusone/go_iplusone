# 03 Crawl Data Manager

Web ページをクロールしてデータを取得・保存・管理するツールの構築ガイドです。

## 目的

ここまで習得した HTTP クライアント・並行処理・database/sql・context・CLI 設計を組み合わせ、実用的なツールを作ることを目指します。

## 前提知識

- `net/http`（`examples/handson/52_network_access_basic`、`60_http_server_basic`）
- `database/sql`（`examples/handson/62_database_sql_basic`、`examples/master/10_db_and_transaction`）
- goroutine / channel / WaitGroup（`examples/handson/46〜55`）
- `context` によるタイムアウト（`examples/handson/57_context_timeout_basic`）
- `encoding/json`（`examples/handson/61_json_api_basic`）

## ツールの全体像

```
crawl-manager/
  cmd/
    crawl/
      main.go         クロール実行コマンド
    export/
      main.go         データエクスポートコマンド
    serve/
      main.go         閲覧用 HTTP サーバ
  internal/
    crawler/
      crawler.go      HTTP 取得とパース
      politeness.go   レート制限と robots.txt の尊重
    store/
      sqlite.go       SQLite へのデータ保存
    exporter/
      json.go         JSON エクスポート
      csv.go          CSV エクスポート
  schema.sql          DB スキーマ定義
  go.mod
```

## DB スキーマ

```sql
CREATE TABLE pages (
  id          INTEGER  PRIMARY KEY AUTOINCREMENT,
  url         TEXT     NOT NULL UNIQUE,
  title       TEXT,
  body_text   TEXT,
  status_code INTEGER  NOT NULL,
  crawled_at  DATETIME NOT NULL,
  error_msg   TEXT
);

CREATE TABLE links (
  id       INTEGER PRIMARY KEY AUTOINCREMENT,
  from_url TEXT NOT NULL,
  to_url   TEXT NOT NULL,
  UNIQUE(from_url, to_url)
);
```

## クローラーの実装

### HTTP 取得の基本

```go
func Fetch(ctx context.Context, url string) (*http.Response, error) {
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("User-Agent", "crawl-manager/1.0")
    return http.DefaultClient.Do(req)
}
```

`context.WithTimeout` でタイムアウトをつけることで、応答が遅いサイトで詰まらなくなります。

### HTML からリンクを抽出する

`golang.org/x/net/html` パッケージを使って HTML をパースします。

```go
import "golang.org/x/net/html"

func ExtractLinks(body io.Reader, base *url.URL) []string {
    var links []string
    doc, err := html.Parse(body)
    if err != nil {
        return links
    }
    var walk func(*html.Node)
    walk = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    if u, err := base.Parse(attr.Val); err == nil {
                        links = append(links, u.String())
                    }
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            walk(c)
        }
    }
    walk(doc)
    return links
}
```

### title タグの取得

```go
func ExtractTitle(body io.Reader) string {
    doc, _ := html.Parse(body)
    var title string
    var walk func(*html.Node)
    walk = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "title" {
            if n.FirstChild != nil {
                title = n.FirstChild.Data
            }
            return
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            walk(c)
        }
    }
    walk(doc)
    return title
}
```

## 並行クローリング

goroutine と channel を使って複数 URL を並行して取得します。

```go
func CrawlAll(ctx context.Context, urls []string, concurrency int) []Result {
    jobs := make(chan string, len(urls))
    results := make(chan Result, len(urls))

    var wg sync.WaitGroup
    for w := 0; w < concurrency; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for url := range jobs {
                result := fetchOne(ctx, url)
                results <- result
            }
        }()
    }

    for _, u := range urls {
        jobs <- u
    }
    close(jobs)

    go func() {
        wg.Wait()
        close(results)
    }()

    var out []Result
    for r := range results {
        out = append(out, r)
    }
    return out
}
```

Worker Pool パターン（`54_worker_pool_basic`）と WaitGroup（`53_waitgroup_basic`）の実践的な応用です。

## レート制限（Politeness）

短時間に大量リクエストを送ると相手サーバに迷惑をかけます。`time.Ticker` でリクエスト間隔を制御します。

```go
type RateLimiter struct {
    ticker *time.Ticker
}

func NewRateLimiter(interval time.Duration) *RateLimiter {
    return &RateLimiter{ticker: time.NewTicker(interval)}
}

func (r *RateLimiter) Wait() {
    <-r.ticker.C
}
```

## データ保存

`database/sql` を使って結果を SQLite に保存します。

```go
func (s *SQLiteStore) SavePage(p *Page) error {
    _, err := s.db.Exec(`
        INSERT OR REPLACE INTO pages (url, title, body_text, status_code, crawled_at, error_msg)
        VALUES (?, ?, ?, ?, ?, ?)`,
        p.URL, p.Title, p.BodyText, p.StatusCode, p.CrawledAt, p.ErrorMsg,
    )
    return err
}
```

プリペアドステートメントを使うことで SQL インジェクションを防ぎます（`62_database_sql_basic` 参照）。

## CLI コマンドの設計

### crawl コマンド

```bash
# 1つの URL をクロール（深さ 2、並行 5）
go run ./cmd/crawl -url https://example.com -depth 2 -concurrency 5

# URL リストファイルからクロール
go run ./cmd/crawl -file urls.txt -concurrency 3
```

### export コマンド

```bash
# JSON エクスポート
go run ./cmd/export -format json -out pages.json

# CSV エクスポート
go run ./cmd/export -format csv -out pages.csv
```

### serve コマンド

```bash
# クロール済みデータを Web で閲覧
go run ./cmd/serve -port 8080
```

## コマンドライン引数のパース

`flag` パッケージで引数を受け取ります。

```go
import "flag"

var (
    targetURL   = flag.String("url", "", "クロール対象の URL")
    depth       = flag.Int("depth", 1, "クロール深さ")
    concurrency = flag.Int("concurrency", 3, "並行数")
)

func main() {
    flag.Parse()
    if *targetURL == "" {
        flag.Usage()
        os.Exit(1)
    }
    // ...
}
```

## エラー設計

ネットワークエラー、タイムアウト、パースエラーなど複数の失敗形態があります。

```go
type CrawlError struct {
    URL string
    Err error
}

func (e *CrawlError) Error() string {
    return fmt.Sprintf("crawl %s: %v", e.URL, e.Err)
}

func (e *CrawlError) Unwrap() error { return e.Err }
```

`errors.Is(err, context.DeadlineExceeded)` でタイムアウトを判定し、リトライ戦略を分けられます。

## JSON / CSV エクスポート

```go
// JSON エクスポート
func ExportJSON(pages []Page, w io.Writer) error {
    return json.NewEncoder(w).Encode(pages)
}

// CSV エクスポート
func ExportCSV(pages []Page, w io.Writer) error {
    cw := csv.NewWriter(w)
    cw.Write([]string{"url", "title", "status_code", "crawled_at"})
    for _, p := range pages {
        cw.Write([]string{p.URL, p.Title, fmt.Sprint(p.StatusCode), p.CrawledAt.Format(time.RFC3339)})
    }
    cw.Flush()
    return cw.Error()
}
```

## 使用するパッケージ

| パッケージ | 用途 |
| --- | --- |
| `net/http` | HTTP リクエスト |
| `golang.org/x/net/html` | HTML パース |
| `database/sql` + `modernc.org/sqlite` | データ保存 |
| `encoding/json` | JSON エクスポート |
| `encoding/csv` | CSV エクスポート |
| `flag` | CLI 引数パース |
| `sync` | WaitGroup、Mutex |
| `context` | タイムアウト・キャンセル |
| `time` | レート制限 |

## セットアップ

```bash
cd examples/appendix/03_crawl_data_manager
go mod init github.com/iplusone/go_iplusone/examples/appendix/03_crawl_data_manager
go get golang.org/x/net/html
go get modernc.org/sqlite
```

## 到達基準

このツールを実装したあと、次の問いに答えられれば十分です。

- なぜ並行クローリングに Worker Pool パターンを使うのか
- なぜ `context.WithTimeout` をリクエストごとにつけるのか
- レート制限がなぜ必要か、どう実装するか
- SQL インジェクションをどう防ぐか
- エラーが起きた URL のデータをどう保持・再試行するか

## 発展トピック

- **robots.txt の尊重**: `temoto/robotstxt` パッケージで解析できます
- **重複 URL の管理**: `Bloom Filter` や `sync.Map` でメモリ効率よく管理できます
- **増分クロール**: `If-Modified-Since` ヘッダで変更があったページだけ再取得できます
- **分散クロール**: ジョブキューを Redis や DB に持たせると複数ワーカーで分散できます
