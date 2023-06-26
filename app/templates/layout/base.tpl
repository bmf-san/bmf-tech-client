{{ define "base" }}
<!DOCTYPE html>
<html lang="ja">
<head>
    {{ template "meta" .Meta }}
    <link rel="icon" href="/favicon.ico" />
    <link rel="stylesheet" href="https://unpkg.com/sea.css/dist/sea.min.css">
    <link rel="stylesheet" href="/css/style.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.5.1/build/styles/monokai.min.css">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.1/highlight.min.js"></script>
    <script>hljs.initHighlightingOnLoad();</script>
    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-784B55NW88"></script>
    <script async>
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
                        <a class="nav-link color-text" href="/posts">記事</a>
                        <a class="nav-link color-text" href="/categories">カテゴリ</a>
                        <a class="nav-link color-text" href="/tags">タグ</a>
                        <a class="nav-link color-text" href="/profile">プロフィール</a>
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
                                <a class="nav-link color-text" href="/posts">記事</a>
                                <a class="nav-link color-text" href="/categories">カテゴリ</a>
                                <a class="nav-link color-text" href="/tags">タグ</a>
                                <a class="nav-link color-text" href="/profile">プロフィール</a>
                            </div>
                        </div>
                    </div>
                </div>
            </nav>
        </div>
    </header>
    {{ template "headline" . }}
    {{ template "body" . }}
    <div class="row text-align-center row-sponsored-link margin-top-2rem">
        <div class="col">
        {{ if isAd }}
        <!-- A8.net -->
        <a href="https://px.a8.net/svt/ejp?a8mat=3TAZOV+6AU69E+348+1BNYOX" rel="nofollow">
            <img border="0" width="468" height="60" alt=""
                src="https://www26.a8.net/svt/bgt?aid=230620207381&wid=001&eno=01&mid=s00000000404008006000&mc=1"></a>
        <img border="0" width="1" height="1" src="https://www13.a8.net/0.gif?a8mat=3TAZOV+6AU69E+348+1BNYOX" alt="">
        {{ else }}
        <!-- A8.net -->
        <a href="https://px.a8.net/svt/ejp?a8mat=3TAZOU+G3ASDU+50+4Z4H29" rel="nofollow">
            <img border="0" width="970" height="90" alt=""
                src="https://www20.a8.net/svt/bgt?aid=230620206973&wid=001&eno=01&mid=s00000000018030086000&mc=1"></a>
        <img border="0" width="1" height="1" src="https://www17.a8.net/0.gif?a8mat=3TAZOU+G3ASDU+50+4Z4H29" alt="">
        {{ end }}
        </div>
    </div>
    <footer class="sticky-footer margin-top-2rem">
        <div class="container">
            <div class="row text-align-center" style="margin: 0 auto; max-width: 800px;">
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
    {{ block "script" . }}{{ end }}
</body>
</html>
{{ end }}


