package site

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"time"
)

//go:embed templates/*.html all:static
var assets embed.FS

type App struct {
	tmpl        *template.Template
	content     map[string]SimplePage
	news        []NewsItem
	songs       []Song
	platforms   []FeatureCard
	chapters    []StudyChapter
	external    []ExternalLink
	topSections []TopSection
}

type NavItem struct {
	Label    string
	URL      string
	External bool
}

type ExternalLink struct {
	Label       string
	URL         string
	Description string
}

type TopSection struct {
	ID          string
	Title       string
	Description string
}

type FeatureCard struct {
	Title       string
	Eyebrow     string
	Description string
	Meta        string
	LinkType    string
	Points      []string
	MetricLabel string
	MetricValue string
	Links       []CTA
	URL         string
	External    bool
}

type CTA struct {
	Label    string
	URL      string
	External bool
}

type SimplePage struct {
	Title    string
	Lead     string
	Sections []ContentSection
	CTAs     []CTA
}

type ContentSection struct {
	Title   string
	Content string
}

type NewsItem struct {
	ID       string
	Title    string
	Category string
	Date     string
	Summary  string
	Body     []string
}

type Song struct {
	ID      string
	Title   string
	Theme   string
	Mood    string
	Summary string
	Lyrics  []string
}

type HomePageData struct {
	Hero          ContentSection
	LearningIntro ContentSection
	Chapters      []StudyChapter
	News          []NewsItem
}

type StudyChapter struct {
	Number      string
	Slug        string
	Title       string
	Summary     string
	Goal        string
	KeyTerms    []string
	Practice    []string
	ReadingPath []string
	Sections    []ContentSection
	CTAs        []CTA
}

type ViewData struct {
	Title         string
	CurrentPath   string
	PrimaryNav    []NavItem
	FooterNav     []NavItem
	ExternalLinks []ExternalLink
	Page          any
	Year          int
}

func New() *App {
	return &App{
		tmpl:        template.Must(template.ParseFS(assets, "templates/*.html")),
		content:     buildContentPages(),
		news:        seedNews(),
		songs:       seedSongs(),
		platforms:   seedPlatforms(),
		chapters:    seedStudyChapters(),
		external:    seedExternalLinks(),
		topSections: seedTopSections(),
	}
}

func (a *App) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/static/styles.css", http.StripPrefix("/static/", http.FileServer(mustSub("static"))))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticFS())))
	mux.HandleFunc("/healthz", a.healthz)
	mux.HandleFunc("/news/detail/", a.newsDetail)
	mux.HandleFunc("/news", a.newsIndex)
	mux.HandleFunc("/gospec/", a.studyChapter)
	mux.HandleFunc("/", a.home)
	return mux
}

func mustSub(dir string) http.FileSystem {
	sub, err := fs.Sub(assets, dir)
	if err != nil {
		panic(err)
	}
	return http.FS(sub)
}

func staticFS() http.FileSystem {
	if _, err := os.Stat("/workspace/static"); err == nil {
		return http.Dir("/workspace/static")
	}
	return mustSub("static")
}

func (a *App) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := ViewData{
		Title:         "GoSpec Learning | go_iplusone",
		CurrentPath:   "/",
		PrimaryNav:    primaryNav(),
		FooterNav:     footerNav(),
		ExternalLinks: a.external,
		Year:          time.Now().Year(),
		Page: HomePageData{
			Hero: ContentSection{
				Title:   "GoSpec Learning Hub",
				Content: "Go の言語仕様を中心に、型・式・文・初期化・実行モデルまでを体系立てて学ぶための入口です。会社紹介や他サービスではなく、Go を読む・試す・理解することに絞ったトップページに整理しました。",
			},
			LearningIntro: ContentSection{
				Title:   "GoSpec Curriculum",
				Content: "Go の言語仕様を、字句・型・式・文・初期化という流れで学び直せる教材セクションです。章を読む順番だけでなく、到達目標と最小の試行ポイントまで見えるようにしています。",
			},
			Chapters: a.chapters,
			News:     a.news,
		},
	}

	a.render(w, http.StatusOK, "home", data)
}

func (a *App) healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func (a *App) page(key string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, ok := a.content[key]
		if !ok {
			http.NotFound(w, r)
			return
		}

		a.render(w, http.StatusOK, "page", ViewData{
			Title:         page.Title + " | GoSpec Learning",
			CurrentPath:   "/" + key,
			PrimaryNav:    primaryNav(),
			FooterNav:     footerNav(),
			ExternalLinks: a.external,
			Year:          time.Now().Year(),
			Page:          page,
		})
	}
}

func (a *App) newsIndex(w http.ResponseWriter, r *http.Request) {
	a.render(w, http.StatusOK, "news_list", ViewData{
		Title:         "News | GoSpec Learning",
		CurrentPath:   "/news",
		PrimaryNav:    primaryNav(),
		FooterNav:     footerNav(),
		ExternalLinks: a.external,
		Year:          time.Now().Year(),
		Page: struct {
			Title string
			Lead  string
			Items []NewsItem
		}{
			Title: "最新のお知らせ",
			Lead:  "Go 学習コンテンツの更新、構成整理、教材追加の進捗をまとめています。",
			Items: a.news,
		},
	})
}

func (a *App) newsDetail(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/news/detail/")
	item, ok := findNews(a.news, id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	a.render(w, http.StatusOK, "news_detail", ViewData{
		Title:         item.Title + " | News | GoSpec Learning",
		CurrentPath:   "/news",
		PrimaryNav:    primaryNav(),
		FooterNav:     footerNav(),
		ExternalLinks: a.external,
		Year:          time.Now().Year(),
		Page:          item,
	})
}

func (a *App) studyChapter(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/gospec/")
	if slug == "" || slug == r.URL.Path {
		http.NotFound(w, r)
		return
	}

	chapter, prev, next, ok := findStudyChapter(a.chapters, slug)
	if !ok {
		http.NotFound(w, r)
		return
	}

	page := chapter
	page.CTAs = append(page.CTAs, CTA{Label: "TOPの学習カリキュラムへ戻る", URL: "/#gospec-curriculum"})
	if prev != nil {
		page.CTAs = append([]CTA{{Label: prev.Number + " " + prev.Title, URL: "/gospec/" + prev.Slug}}, page.CTAs...)
	}
	if next != nil {
		page.CTAs = append(page.CTAs, CTA{Label: next.Number + " " + next.Title, URL: "/gospec/" + next.Slug})
	}

	a.render(w, http.StatusOK, "study_chapter", ViewData{
		Title:         page.Title + " | GoSpec Learning",
		CurrentPath:   "/gospec/" + slug,
		PrimaryNav:    primaryNav(),
		FooterNav:     footerNav(),
		ExternalLinks: a.external,
		Year:          time.Now().Year(),
		Page:          page,
	})
}

func (a *App) render(w http.ResponseWriter, status int, name string, data ViewData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if err := a.tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func primaryNav() []NavItem {
	return []NavItem{
		{Label: "TOP", URL: "/"},
		{Label: "GOSPEC", URL: "/#gospec-curriculum"},
		{Label: "NEWS", URL: "/news"},
		{Label: "HEALTH", URL: "/healthz"},
	}
}

func footerNav() []NavItem {
	return []NavItem{
		{Label: "GoSpec Curriculum", URL: "/#gospec-curriculum"},
		{Label: "Chapter Updates", URL: "/news"},
		{Label: "Health Check", URL: "/healthz"},
	}
}
