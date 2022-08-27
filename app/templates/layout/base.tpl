{{ define "base" }}
<!DOCTYPE html>
<html lang="ja">
<head>
    <title>bmf-tech.com</title>
    {{ template "meta" .Meta }}
    <link rel="stylesheet" href="https://unpkg.com/sea.css/dist/sea.min.css">
    <link rel="stylesheet" href="/style.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.5.1/build/styles/monokai.min.css">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.1/highlight.min.js"></script>
    <script>hljs.initHighlightingOnLoad();</script>
    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-784B55NW88"></script>
    <script>
        window.dataLayer = window.dataLayer || [];
        function gtag() { dataLayer.push(arguments); }
        gtag('js', new Date());

        gtag('config', 'G-784B55NW88');
    </script>
</head>
<body>
    <header>
        <div>
            <nav class="nav sp-hidden">
                <div class="col-nav">
                    <div class="nav-left">
                        <a class="nav-link-logo color-text" href="/"><b>bmf-tech</b></a>
                    </div>
                    <div class="nav-right">
                        <a class="nav-link color-text" href="/">ホーム</a>
                        <a class="nav-link color-text" href="/posts">記事</a>
                        <a class="nav-link color-text" href="/categories">カテゴリ</a>
                        <a class="nav-link color-text" href="/tags">タグ</a>
                    </div>
                </div>
            </nav>
            <nav class="nav pc-hidden">
                <div class="col-nav">
                    <div class="nav-left">
                        <a class="nav-link text-decoration-none color-text">bmf-tech</a>
                    </div>
                    <div class="nav-right">
                        <div id="nav-sp-drawer">
                            <input id="nav-sp-input" type="checkbox" class="sp-hidden">
                            <label id="nav-sp-open" for="nav-sp-input"><span></span></label>
                            <label class="sp-hidden" id="nav-sp-close" for="nav-sp-input"><span></span></label>
                            <div id="nav-sp-content">
                                <h1>bmf-tech</h1>
                                <a class="nav-link color-text" href="/">ホーム</a>
                                <a class="nav-link color-text" href="/posts">記事</a>
                                <a class="nav-link color-text" href="/categories">カテゴリ</a>
                                <a class="nav-link color-text" href="/tags">タグ</a>
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
                    <a class="color-text-reverse" href="/sitemap">Sitemap</a>
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


