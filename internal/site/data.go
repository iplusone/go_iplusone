package site

func buildContentPages() map[string]SimplePage {
	return map[string]SimplePage{
		"blog": {
			Title: "Blog",
			Lead:  "技術メモ、プロダクトづくり、事業設計の気づきを蓄積するためのブログ一覧ページです。",
			Sections: []ContentSection{
				{Title: "編集方針", Content: "本家サイトの文脈を引き継ぎつつ、Go 実装や移植方針、運用ルールのナレッジを順次追加できる構成にしています。"},
				{Title: "今後の拡張", Content: "カテゴリ一覧、タグ、月別アーカイブ、個別記事ページへ広げやすい導線の起点として実装しています。"},
			},
			CTAs: []CTA{{Label: "TOPへ戻る", URL: "/"}},
		},
		"contact": {
			Title: "Contact",
			Lead:  "お問い合わせの入口です。現段階ではフォーム送信は未接続ですが、問い合わせ種別ごとの受け口を整理しています。",
			Sections: []ContentSection{
				{Title: "相談できる内容", Content: "Web 制作、業務システム、予約・入札プラットフォーム、既存サイト移行、実験的な AI 活用までを想定しています。"},
				{Title: "次の実装候補", Content: "問い合わせフォーム、バリデーション、送信完了画面、メール通知を順次追加しやすい構成です。"},
			},
			CTAs: []CTA{{Label: "NEWSを見る", URL: "/news"}},
		},
		"hp_templates": {
			Title: "HP Templates",
			Lead:  "テンプレートベースで立ち上げられるサイト構成を紹介するページです。",
			Sections: []ContentSection{
				{Title: "用途", Content: "スモールスタート向けコーポレートサイト、採用ページ、地域情報サイトなど、初期速度重視の構成を並べられます。"},
				{Title: "運用イメージ", Content: "将来的にはテンプレートカード一覧、特徴、価格感、デモ導線をここへ追加していきます。"},
			},
		},
		"project-works": {
			Title: "Project Works",
			Lead:  "公開可能な制作実績や取り組みをまとめる一覧ページです。",
			Sections: []ContentSection{
				{Title: "掲載単位", Content: "プロジェクトごとに課題、提供機能、技術構成、成果を並べられる前提で構成しています。"},
				{Title: "TOPとの接続", Content: "TOP から事例へ自然に誘導するため、将来的にはカード表示やカテゴリ絞り込みを追加します。"},
			},
		},
		"corporates": {
			Title: "Corporates",
			Lead:  "企業向けの導入カテゴリや情報導線を整理するページです。",
			Sections: []ContentSection{
				{Title: "想定読者", Content: "会社案内、導入検討、実績確認、問い合わせ前の比較検討をしたい訪問者を想定しています。"},
				{Title: "今後の拡張", Content: "業種別導線、課題別導線、CTA 最適化などを載せやすい構成です。"},
			},
		},
		"map": {
			Title: "National Map",
			Lead:  "地域軸の導線を持たせるためのマップページです。",
			Sections: []ContentSection{
				{Title: "役割", Content: "都道府県別、エリア別、ジャンル別の入口へ接続するためのハブページとして使えます。"},
				{Title: "今後の実装", Content: "地図 UI や都道府県リンク一覧、関連ページの集約を追加できるようにしています。"},
			},
		},
		"homepages": {
			Title: "Homepages",
			Lead:  "ホームページ制作サービスの紹介ページです。",
			Sections: []ContentSection{
				{Title: "訴求軸", Content: "スピード、運用性、拡張性のバランスを重視したサイト制作の説明を集約できます。"},
				{Title: "関連導線", Content: "テンプレート一覧、制作実績、問い合わせへの接続を想定しています。"},
			},
		},
		"railway_routes/front": {
			Title: "Railway Routes",
			Lead:  "鉄道沿線や駅単位の導線ページです。",
			Sections: []ContentSection{
				{Title: "利用イメージ", Content: "エリア集客型ページやローカル検索導線の入口として運用できる前提で作っています。"},
				{Title: "展開イメージ", Content: "沿線一覧、駅一覧、関連施設や事業ページへのリンクをここに展開できます。"},
			},
		},
		"suno-library": {
			Title: "SUNO Library",
			Lead:  "音楽生成や関連素材を蓄積していくためのライブラリページです。",
			Sections: []ContentSection{
				{Title: "用途", Content: "公開済み楽曲の整理、制作メモ、再利用できるフレーズやテーマのストックを置く想定です。"},
				{Title: "今後の拡張", Content: "フィルタ、タグ、ジャンル分類などを追加しやすい構成です。"},
			},
		},
		"platforms/bidding": {
			Title: "Bidding Platform",
			Lead:  "入札・見積もり比較を支えるプラットフォームの説明ページです。",
			Sections: []ContentSection{
				{Title: "想定機能", Content: "案件登録、条件管理、比較表示、参加企業導線を中心に設計できるようにしています。"},
				{Title: "TOPでの役割", Content: "7つの基盤PRや AI-DDD セクションから遷移する説明ページとして配置しています。"},
			},
			CTAs: []CTA{{Label: "Reservation Platform も見る", URL: "/platforms/reservation"}},
		},
		"platforms/reservation": {
			Title: "Reservation Platform",
			Lead:  "予約フローや空き枠管理を支えるプラットフォームの説明ページです。",
			Sections: []ContentSection{
				{Title: "想定機能", Content: "空き枠表示、予約受付、通知、管理画面への展開を見据えた説明ページです。"},
				{Title: "TOPでの役割", Content: "Bidding Platform と並ぶ基盤紹介ページとして構成しています。"},
			},
			CTAs: []CTA{{Label: "Bidding Platform も見る", URL: "/platforms/bidding"}},
		},
	}
}

func seedTopSections() []TopSection {
	return []TopSection{
		{
			ID:          "about",
			Title:       "会社概要",
			Description: "数多くのコンピュータシステム開発、とくに WEB サイトサービスの構築を中心に、サーバ環境構築、維持管理、企画から実装・テストまで一貫して対応してきた会社です。",
		},
		{
			ID:          "service",
			Title:       "SERVICE",
			Description: "サーバ環境、プログラミング言語、WEBアプリケーションフレームワークを最大限活かした構築支援を提供します。",
		},
	}
}

func seedPlatforms() []FeatureCard {
	return []FeatureCard{
		{
			Eyebrow:     "EC Site Platform",
			Title:       "ECサイト基盤",
			Description: "売る仕組みを、すぐ始められる。出店から決済、売上管理までを一気通貫で支えるEC基盤です。",
			Meta:        "基盤: ec-plus1 / 案内先: 説明ポータル",
			LinkType:    "説明ポータル",
			Points:      []string{"ブランド別にすぐ出店", "決済と注文管理を標準装備", "在庫と売上をまとめて見える化"},
			MetricLabel: "強み",
			MetricValue: "最短で販売開始",
			URL:         "https://ec-plus1.iplusone.co.jp/",
			External:    true,
			Links:       []CTA{{Label: "説明ポータルを見る", URL: "https://ec-plus1.iplusone.co.jp/", External: true}},
		},
		{
			Eyebrow:     "Estimate Platform",
			Title:       "見積もり基盤",
			Description: "提案スピードを、仕組みで引き上げる。見積作成からPDF化までをスマートに回せる基盤です。",
			Meta:        "基盤: estimates / 案内先: 説明ポータル",
			LinkType:    "説明ポータル",
			Points:      []string{"見積作成をスピードアップ", "PDF出力までワンストップ", "履歴管理で提案を再利用しやすい"},
			MetricLabel: "強み",
			MetricValue: "提案業務を省力化",
			URL:         "https://estimates.iplusone.co.jp/",
			External:    true,
			Links:       []CTA{{Label: "説明ポータルを見る", URL: "https://estimates.iplusone.co.jp/", External: true}},
		},
		{
			Eyebrow:     "Bidding Platform",
			Title:       "入札サイト",
			Description: "見積もり基盤を使って、複数業者から同じフォーマットで見積を集め、比較検討しやすくした入札サイトです。",
			Meta:        "基盤: estimates / 案内先: 説明ページ",
			LinkType:    "説明ページ",
			Points:      []string{"各社に同一フォーマットで入力依頼", "見積内容を横並びで比較しやすい", "判断材料を整理して蓄積しやすい"},
			MetricLabel: "強み",
			MetricValue: "比較検討を進めやすい",
			URL:         "/platforms/bidding",
			Links:       []CTA{{Label: "入札システムを見る", URL: "/platforms/bidding"}},
		},
		{
			Eyebrow:     "Community Platform",
			Title:       "コミュニティ基盤",
			Description: "人が集まり、関係が育つ。発信・参加・交流を生み出し続けるコミュニティ基盤です。",
			Meta:        "基盤: plus1_community / 案内先: 実サイト",
			LinkType:    "実サイト",
			Points:      []string{"参加導線をつくりやすい", "投稿と交流の場を育てやすい", "継続運営しやすい拡張性"},
			MetricLabel: "強み",
			MetricValue: "つながりを価値に変える",
			URL:         "https://spottown.iplusone.co.jp/",
			External:    true,
			Links:       []CTA{{Label: "spottownを見る", URL: "https://spottown.iplusone.co.jp/", External: true}},
		},
		{
			Eyebrow:     "Matching Platform",
			Title:       "マッチング基盤",
			Description: "人と仕事、人と募集をなめらかにつなぐ。求人や応募導線を扱いやすくするマッチング基盤です。",
			Meta:        "基盤: communication / 案内先: 実サイト / 基盤説明",
			LinkType:    "実サイト",
			Points:      []string{"求人情報を整理して掲載", "応募フローを一元管理", "採用や募集業務に合わせて拡張しやすい"},
			MetricLabel: "強み",
			MetricValue: "募集と応募をつなぎやすい",
			URL:         "https://jobs.iplusone.co.jp/",
			External:    true,
			Links: []CTA{
				{Label: "求人マッチングサイトを見る", URL: "https://jobs.iplusone.co.jp/", External: true},
				{Label: "マッチング基盤の説明を見る", URL: "https://jobs.iplusone.co.jp/platform", External: true},
			},
		},
		{
			Eyebrow:     "Real Estate Platform",
			Title:       "不動産物件サイト構築基盤",
			Description: "物件情報を見やすく整理し、掲載から問い合わせ導線までを一体で設計できる不動産物件サイト構築基盤です。",
			Meta:        "基盤: real_estate_website / 案内先: 実サイト",
			LinkType:    "実サイト",
			Points:      []string{"物件一覧と詳細ページを整理して掲載", "検索導線や問い合わせ導線を設計しやすい", "地域や用途に合わせて拡張しやすい"},
			MetricLabel: "強み",
			MetricValue: "物件訴求と反響獲得を両立",
			URL:         "https://real-estate-website.iplusone.co.jp/",
			External:    true,
			Links:       []CTA{{Label: "不動産物件サイトを見る", URL: "https://real-estate-website.iplusone.co.jp/", External: true}},
		},
		{
			Eyebrow:     "Reservation Platform",
			Title:       "予約システム基盤",
			Description: "来店や面談、各種申込の予約受付をわかりやすく整理し、登録から完了導線までをスムーズにつなぐ予約システム基盤です。",
			Meta:        "基盤: SoloReserve / 案内先: 説明ページ",
			LinkType:    "説明ページ",
			Points:      []string{"予約フォームを用途に合わせて構築しやすい", "申込から完了までの導線を整理しやすい", "業種ごとの受付フローに合わせて拡張しやすい"},
			MetricLabel: "強み",
			MetricValue: "受付業務をオンライン化しやすい",
			URL:         "/platforms/reservation",
			Links:       []CTA{{Label: "予約システム基盤を見る", URL: "/platforms/reservation"}},
		},
	}
}

func seedExternalLinks() []ExternalLink {
	return nil
}

func seedNews() []NewsItem {
	return []NewsItem{
		{
			ID:       "2026-05-go-top-structure",
			Title:    "Go 学習トップを GoSpec 専用に再整理",
			Category: "Development",
			Date:     "2026-05-02",
			Summary:  "会社紹介や他サービス導線を外し、Go の学習導線を中心にしたトップ構成へ整理し始めました。",
			Body: []string{
				"トップは GoSpec カリキュラムと更新情報に絞り、見た瞬間に目的が伝わる構成へ寄せています。",
				"まずはナビゲーション、共通レイアウト、トップ文言の整理を優先しています。",
			},
		},
		{
			ID:       "2026-05-public-pages-ready",
			Title:    "GoSpec 各章ページの導線を確認しやすく改善",
			Category: "Release",
			Date:     "2026-05-03",
			Summary:  "トップから章詳細へ進み、前後の章へ移動しやすい流れを整えました。",
			Body: []string{
				"各章ページでは到達目標と試すことを追いやすい表示を維持しています。",
				"今後は章間の関係や補助資料の見せ方も整えていきます。",
			},
		},
		{
			ID:       "2026-05-gospec-content-seed",
			Title:    "GoSpec 学習コンテンツの初期データを追加",
			Category: "Content",
			Date:     "2026-05-04",
			Summary:  "章一覧、重要用語、最小の演習ポイントを含む初期データセットを揃えました。",
			Body: []string{
				"将来的には演習コードや仕様引用の補助表示を足せる構成です。",
				"まずは学習導線と詳細表現の型を揃えることを目的にしています。",
			},
		},
	}
}

func seedCompanyFacts() []ContentSection {
	return []ContentSection{
		{Title: "商号", Content: "株式会社 i PLUS ONE（アイプラスワン）"},
		{Title: "設立", Content: "2005年5月2日"},
		{Title: "代表者", Content: "代表取締役 石井 伸一"},
		{Title: "資本金", Content: "10,000,000円"},
		{Title: "事業内容", Content: "コンピュータソフトウェア設計、開発、販売、保守管理、情報処理および情報提供、ホームページ企画制作など。"},
	}
}

func seedServices() []ContentSection {
	return []ContentSection{
		{Title: "Linuxサーバ(CentOS, AlmaLinux, RockyLinux)環境構築", Content: "サーバ環境構築"},
		{Title: "DockerDesktop, GitHub, Visual Studio Code", Content: "開発基盤整備"},
		{Title: "Web Design、HTML/CSS Code Generate and Check、AI refactoring", Content: "デザインとコード改善"},
		{Title: "WEBサイト構築全般支援", Content: "企画から公開まで支援"},
		{Title: "Linux Webアプリケーションサーバ環境構築 Apache/Nginx", Content: "Webサーバ構築"},
		{Title: "WEBアプリケーション開発 PHP/Laravel+MySQL", Content: "業務アプリ開発"},
		{Title: "サーバサイド開発", Content: "バックエンド実装"},
		{Title: "ホームページ・ドメイン・サーバの月定額サービス", Content: "継続運用支援"},
	}
}

func seedStudyChapters() []StudyChapter {
	return []StudyChapter{
		{
			Number:  "Chapter 1",
			Slug:    "reading-and-lexical-basics",
			Title:   "仕様書の読み方と字句の基礎",
			Summary: "GoSpec を読むための入口です。EBNF、トークン、セミコロン自動挿入を押さえ、仕様書がどんな視点で書かれているかを掴みます。",
			Goal:    "仕様書の文法記述を怖がらずに読み、コードがまず字句へ分解される流れを理解する。",
			KeyTerms: []string{
				"EBNF",
				"token",
				"identifier",
				"literal",
				"semicolon insertion",
			},
			Practice: []string{
				"`a := 1 + 2` をトークン単位に分解してみる",
				"`if` 文の EBNF を見て `[]` と `|` を読み替える",
				"セミコロンを書かなくても成立する例と崩れる例を比べる",
			},
			ReadingPath: []string{
				"表記",
				"字句要素",
				"トークン",
				"セミコロン",
				"識別子",
			},
			Sections: []ContentSection{
				{Title: "この章の役割", Content: "GoSpec の最初の壁は、専門用語よりも書き方そのものです。この章では、仕様書が EBNF という記法を使って文法を定義していること、そして Go のソースコードが最初に字句要素へ分解されることを押さえます。"},
				{Title: "ここで掴むべき感覚", Content: "Go のコードは、いきなり意味として解釈されるのではなく、まず token の列として読まれます。`識別子` `演算子` `リテラル` を見分けられるようになると、コンパイラが何を読み取っているかを追いやすくなります。"},
				{Title: "演習の進め方", Content: "短いコードを一行ずつ token に分解し、セミコロン自動挿入がどこで起きるかを観察します。ここでは正しさよりも、仕様の言葉とコードの見た目を対応づけることが大切です。"},
			},
		},
		{
			Number:  "Chapter 2",
			Slug:    "type-system-foundation",
			Title:   "型システムの土台を作る",
			Summary: "Go の型を、単なる文法ではなくルールの集合として理解する章です。各種型、メソッド集合、interface の土台をここで固めます。",
			Goal:    "型が値の集合と操作を定めていること、named type と type literal の違いを説明できるようにする。",
			KeyTerms: []string{
				"type",
				"named type",
				"type literal",
				"method set",
				"interface",
			},
			Practice: []string{
				"`struct` `slice` `map` `interface` を 1 つずつ最小コードで宣言する",
				"`type MyInt int` と `[]int` の違いを言葉で説明する",
				"`any` に値を入れて動的型を確認する",
			},
			ReadingPath: []string{
				"型",
				"メソッド集合",
				"構造体型",
				"スライス型",
				"インターフェース型",
			},
			Sections: []ContentSection{
				{Title: "この章の役割", Content: "Go の型を、単なる書き方ではなく、値の範囲と操作の約束として捉え直す章です。ここで各種型の輪郭を押さえると、後続の代入可能性やメソッド集合が理解しやすくなります。"},
				{Title: "named type と type literal", Content: "仕様を読む時に重要なのは、名前を持つ型と構造そのものを表す型リテラルを分けて考えることです。`type MyInt int` と `[]int` は、見た目以上に役割が違います。"},
				{Title: "演習の進め方", Content: "構造体、スライス、マップ、interface をそれぞれ最小コードで宣言し、どんな値を入れられるか、どんな操作ができるかを観察します。できれば `%T` で型を表示しながら確認すると理解が定着します。"},
			},
		},
		{
			Number:  "Chapter 3",
			Slug:    "underlying-type-and-assignability",
			Title:   "基底型・代入可能性・定数を理解する",
			Summary: "GoSpec の中核です。基底型、型の一意性、代入可能性、型なし定数、表現可能性をつなげて理解します。",
			Goal:    "なぜその代入が通るのか、なぜコンパイルエラーになるのかを仕様ベースで説明できるようにする。",
			KeyTerms: []string{
				"underlying type",
				"assignability",
				"representability",
				"untyped constant",
				"type identity",
			},
			Practice: []string{
				"`int` と `MyInt` が直接代入できない例を試す",
				"`type MySlice []int` に `[]int{}` を代入して確認する",
				"`var b int8 = 1000` を書いて表現可能性の境界を確かめる",
			},
			ReadingPath: []string{
				"型と値の性質",
				"型の一意性",
				"代入可能性",
				"表現可能性",
				"定数",
			},
			Sections: []ContentSection{
				{Title: "この章の役割", Content: "GoSpec を読む価値が最も見えやすい章です。基底型、型の一意性、代入可能性、型なし定数は別々の知識ではなく、Go が型安全性と利便性をどう両立しているかを説明するひとつの体系です。"},
				{Title: "特に意識すること", Content: "教材では `defined type` で説明されることがありますが、仕様を正確に追う時は `named type` を意識する方が安全です。`int` と `MyInt`、`[]int` と `type MySlice []int` を並べて違いを見ると理解しやすくなります。"},
				{Title: "演習の進め方", Content: "通る代入と通らない代入を意図的に並べ、エラーメッセージと仕様の記述を照らし合わせます。定数は `型なし` の状態から文脈で型が決まるため、変数との違いも一緒に試すのがおすすめです。"},
			},
		},
		{
			Number:  "Chapter 4",
			Slug:    "expressions-and-evaluation",
			Title:   "式・呼び出し・評価順序を読む",
			Summary: "毎日の Go コードに直結する章です。アドレス演算子、呼び出し、variadic、型変換、型アサーションなどをまとめて扱います。",
			Goal:    "値がどう評価され、どの条件で操作できるかを把握し、日常コードの挙動を論理で追えるようにする。",
			KeyTerms: []string{
				"expression",
				"addressable",
				"conversion",
				"type assertion",
				"evaluation order",
			},
			Practice: []string{
				"`&s[0]` は通るが `&m[\"a\"]` は通らない例を比べる",
				"variadic 引数に何も渡さない時の値を確認する",
				"`x.(string)` と `string(x)` の違いを整理する",
			},
			ReadingPath: []string{
				"式",
				"呼び出し",
				"引数に ... パラメーターを渡す",
				"アドレス演算子",
				"変換",
			},
			Sections: []ContentSection{
				{Title: "この章の役割", Content: "日常の Go コードで最も頻繁に遭遇するのが `式` のルールです。関数呼び出し、引数展開、アドレス演算子、変換、型アサーションは、すべてこの章の中で一貫した評価規則として定義されています。"},
				{Title: "addressable をどう見るか", Content: "初心者がつまずきやすい `&m[key]` と `&s[i]` の差は、仕様上は `addressable` かどうかの違いです。実装上の事情を覚える前に、まず仕様が何を許しているかを押さえると整理しやすくなります。"},
				{Title: "演習の進め方", Content: "通る例と落ちる例を 1 セットにして比べるのが効果的です。variadic 引数、型変換、型アサーションは似た見た目でも意味が違うので、それぞれ別のルールとして切り分けて試します。"},
			},
		},
		{
			Number:  "Chapter 5",
			Slug:    "statements-and-concurrency",
			Title:   "文・制御構文・並行処理をつなげる",
			Summary: "if、for、switch の基本に加えて、go、select、defer までを仕様として読みます。制御の流れと並行処理の意味をつなぐ章です。",
			Goal:    "Go の実行フローを文単位で理解し、goroutine と channel を含む制御の意味を説明できるようにする。",
			KeyTerms: []string{
				"statement",
				"goroutine",
				"channel",
				"select",
				"defer",
			},
			Practice: []string{
				"`defer` が逆順実行されるコードを試す",
				"buffered / unbuffered channel の違いを小さく試す",
				"`select` の `default` がある時とない時の違いを見る",
			},
			ReadingPath: []string{
				"文",
				"if 文",
				"for 文",
				"go 文",
				"select 文",
				"defer 文",
			},
			Sections: []ContentSection{
				{Title: "この章の役割", Content: "ここでは Go コードがどのように進み、分岐し、待ち、後処理を行うかを文単位で理解します。制御構文の理解に `go` `channel` `select` をつなげることで、Go の並行処理を仕様の言葉で捉えられるようになります。"},
				{Title: "goroutine と channel の見方", Content: "goroutine は OS スレッドそのものではなく、Go ランタイムが扱う並行実行単位です。channel は通信手段であると同時に同期手段でもあり、`select` は複数通信の切り替え点として機能します。"},
				{Title: "演習の進め方", Content: "`defer` の逆順実行、buffered/unbuffered channel、`select` の default の有無など、挙動が目に見えやすい題材を小さく試します。ここは動かして初めて腑に落ちる章です。"},
			},
		},
		{
			Number:  "Chapter 6",
			Slug:    "initialization-and-program-lifecycle",
			Title:   "初期化・パッケージ・実行モデルを押さえる",
			Summary: "最後に、Go プログラム全体がどう起動するかを学びます。ゼロ値、パッケージ初期化、init、main の関係を扱います。",
			Goal:    "プログラムが静かに準備を終えてから main に到達する流れを理解し、初期化まわりのバグを避けられるようにする。",
			KeyTerms: []string{
				"zero value",
				"package initialization",
				"init",
				"main",
				"import dependency",
			},
			Practice: []string{
				"パッケージ変数と `init()` の実行順を表示してみる",
				"`nil` スライスと空スライスの違いを確かめる",
				"`close` 後の channel 受信を最小コードで確認する",
			},
			ReadingPath: []string{
				"ビルトイン関数",
				"閉じる",
				"プログラムの初期化と実行",
				"ゼロ値",
				"パッケージの初期化",
			},
			Sections: []ContentSection{
				{Title: "この章の役割", Content: "最後の章では、Go プログラム全体がどのように準備され、どの順で起動するかを整理します。日常のバグで見落としやすい `zero value`、`init()`、パッケージ変数初期化の順序を、仕様の側から確認します。"},
				{Title: "初期化順序を見る視点", Content: "Go の初期化は単なる宣言順ではなく、依存関係を満たしながら進みます。`import`、パッケージ変数、`init()`、`main()` の順をひとまとまりで理解すると、起動時の不思議な挙動がかなり説明しやすくなります。"},
				{Title: "演習の進め方", Content: "変数初期化時に出力を入れたり、`init()` でログを出したりして、実行順を目で確認します。`nil` スライスと空スライス、閉じた channel からの受信も一緒に試すと理解が深まります。"},
			},
		},
	}
}

func seedSongs() []Song {
	return []Song{
		{
			ID:      "city-lights-demo",
			Title:   "City Lights Demo",
			Theme:   "Urban Pop",
			Mood:    "Bright",
			Summary: "公開フロントの雰囲気確認用に置くデモ楽曲です。",
			Lyrics: []string{
				"Signal を越えて、次の景色へ向かう。",
				"小さな改善を積み上げて、大きな体験に変えていく。",
			},
		},
		{
			ID:      "platform-echo",
			Title:   "Platform Echo",
			Theme:   "Synth Rock",
			Mood:    "Driven",
			Summary: "サービス導線やプロダクト説明の熱量をイメージしたデモ楽曲です。",
			Lyrics: []string{
				"ばらばらの要素を、ひとつの流れに束ねる。",
				"設計と体験が重なる場所を探していく。",
			},
		},
	}
}

func findNews(items []NewsItem, id string) (NewsItem, bool) {
	for _, item := range items {
		if item.ID == id {
			return item, true
		}
	}
	return NewsItem{}, false
}

func findSong(items []Song, id string) (Song, bool) {
	for _, item := range items {
		if item.ID == id {
			return item, true
		}
	}
	return Song{}, false
}

func findStudyChapter(items []StudyChapter, slug string) (StudyChapter, *StudyChapter, *StudyChapter, bool) {
	for i, item := range items {
		if item.Slug != slug {
			continue
		}

		var prev *StudyChapter
		var next *StudyChapter
		if i > 0 {
			prev = &items[i-1]
		}
		if i+1 < len(items) {
			next = &items[i+1]
		}
		return item, prev, next, true
	}
	return StudyChapter{}, nil, nil, false
}
