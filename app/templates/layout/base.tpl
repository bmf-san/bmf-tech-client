{{ define "base" }}
<!DOCTYPE html>
<html lang="ja">
<head>
    <title>bmf-tech.com</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="description" content="" />
    <meta name="og:title" content="" />
    <meta name="og:description" content="" />
    <meta property="og:url" content="" />
    <meta property="og:type" content="">
    <meta property="og:image" content="" />
    <meta property="og:site_name" content="" />
    <meta property="og:locale" content="ja_JP" />
    <meta name="twitter:card" content="" />
    <meta name="twitter:site" content="" />
	<!-- TODO: noindexは本番リリース後に外す -->
	<meta name="robots" content="noindex">
	<link rel="stylesheet" href="https://unpkg.com/sea.css/dist/sea.min.css">
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.5.1/build/styles/monokai.min.css">
	<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.1/highlight.min.js"></script>
	<script>hljs.initHighlightingOnLoad();</script>
</head>
<!-- TODO: 本番リリース前にコメントアウト外して有効化 -->
<!-- <script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-5146230866088201"
	crossorigin="anonymous"></script> -->
<!-- bmf-tech -->
<!-- TODO: 本番リリース前にコメントアウト外して有効化 -->
<!-- <ins class="adsbygoogle" style="display:block" data-ad-client="ca-pub-5146230866088201" data-ad-slot="9196954828"
	data-ad-format="auto" data-full-width-responsive="true"></ins>
<script>
	(adsbygoogle = window.adsbygoogle || []).push({});
</script> -->
<body>
    <header>
        <div>
            <nav class="nav sp-hidden">
                <div class="col-nav">
                    <div class="nav-left">
                        <a class="nav-link-logo"><b>bmf-tech</b></a>
                    </div>
                    <div class="nav-right">
                        <a class="nav-link" href="/">ホーム</a>
                        <a class="nav-link" href="/posts">記事</a>
                        <a class="nav-link" href="/categories">カテゴリ</a>
                        <a class="nav-link" href="/tags">タグ</a>
                    </div>
                </div>
            </nav>
            <nav class="nav pc-hidden">
                <div class="col-nav">
                    <div class="nav-left">
                        <a class="nav-link text-decoration-none">bmf-tech</a>
                    </div>
                    <div class="nav-right">
                        <div id="nav-sp-drawer">
                            <input id="nav-sp-input" type="checkbox" class="sp-hidden">
                            <label id="nav-sp-open" for="nav-sp-input"><span></span></label>
                            <label class="sp-hidden" id="nav-sp-close" for="nav-sp-input"><span></span></label>
                            <div id="nav-sp-content">
                                <h1>bmf-tech</h1>
                                <a class="nav-link" href="/">ホーム</a>
                                <a class="nav-link" href="/posts">記事</a>
                                <a class="nav-link" href="/categories">カテゴリ</a>
                                <a class="nav-link" href="/tags">タグ</a>
                            </div>
                        </div>
                    </div>
                </div>
            </nav>
        </div>
    </header>

    {{ template "headline" . }}

    {{ template "body" . }}

    <footer class="sticky-footer margin-top-2rem">
        <div class="container">
            <div class="row text-align-center">
                <div class="col">
                    <a class="color-text-reverse" href="https://twitter.com/bmf_san">Twitter</a>
                </div>
                <div class="col">
                    <a class="color-text-reverse" href="https://github.com/bmf-san">Github</a>
                </div>
                <div class="col">
                    <a class="color-text-reverse" href="/sitemap">RSS</a>
                </div>
                <div class="col">
                    <a class="color-text-reverse" href="/feed">Feed</a>
                </div>
            </div>
            <hr>
            <div class="row text-align-center">
                <div class="col">
					<small>©{{ year }} Kenta Takeuchi</small>
                </div>
            </div>
        </div>
    </footer>
</body>
</html>
{{ end }}


