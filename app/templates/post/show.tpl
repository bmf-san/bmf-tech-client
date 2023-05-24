{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <div class="text-align-center">
                    <h1 class="text-align-center color-text-reverse">{{ .Post.Post.Title }}</h1>
                    <p><a class="color-text-reverse" href="/posts/categories/{{ .Post.Post.Category.Name }}">{{ .Post.Post.Category.Name }}</a></p>
                    <p><span class="article-date color-text-reverse">{{ .Post.Post.UpdatedAt.Format "2006年1月2日 01:02:03" }} 更新</span></p>
                </div>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "body" }}
<section>
    <div class="container-readable">
        <div class="row">
            <div class="col">
                {{ if isAd }}
                <div>
                    <!-- Google Adsense インフィード -->
                    <script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-5146230866088201"
                    crossorigin="anonymous"></script>
                    <ins class="adsbygoogle" style="display:block" data-ad-format="fluid" data-ad-layout-key="-gw-3+1f-3d+2z"
                    data-ad-client="ca-pub-5146230866088201" data-ad-slot="7900864416"></ins>
                    <script>(adsbygoogle = window.adsbygoogle || []).push({});</script>
                </div>
                {{ end }}
                <div class="margin-top-2rem margin-bottom-2rem">
                {{ range .Post.Post.Tags }}
                <a class="tag" href="/posts/tags/{{ .Name }}">{{ .Name }}</a>
                {{ end }}
                <div class="margin-top-4rem col-toc">
                    <ul id="toc"></ul>
                </div>
                </div>
                <div class="article">
                {{ .Post.Post.HTMLBody | unescape }}
                </div>
            </div>
        </div>
        <hr>
        <div class="row">
            <div class="col">
                <div class="margin-top-3rem">
                    <p><u>関連書籍</u></p>
                    <ul>
                        {{ range .Post.Post.Tags }}
                        <li>
                            <a target="_blank" href="https://www.amazon.co.jp/gp/search?ie=UTF8&tag=bmf035-22&linkCode=ur2&linkId=1c71f3af4b958fd5f0a319848e40b1df&camp=247&creative=1211&index=books&keywords={{ .Name }}">{{.Name }}</a>
                        </li>
                        {{ end }}
                    </ul>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <button style="width:100%" onclick="window.open('https://paypal.me/bmfsan?country.x=JP&locale.x=ja_JP')">bmf-tech.comにお布施する！</button>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <!-- Google Adsense 記事内 -->
                <script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-5146230866088201" crossorigin="anonymous"></script>
                <ins class="adsbygoogle" style="display:block; text-align:center;" data-ad-layout="in-article" data-ad-format="fluid" data-ad-client="ca-pub-5146230866088201" data-ad-slot="5419252877"></ins>
                <script>(adsbygoogle = window.adsbygoogle || []).push({});</script>
            </div>
        </div>
    </div>
</section>
{{ end }}
{{ define "script" }}
<script async src="../js/toc.js"></script>
{{ end }}