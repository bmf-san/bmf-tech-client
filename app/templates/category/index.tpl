{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-align-center color-text-reverse">カテゴリ</h1>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "body" }}
<div class="container-readable margin-top-2rem">
    <div class="row">
        <div class="col">
            <ul>
            {{ range $i, $v := .Categories.Categories }}
            <li>
                <a href="/posts/categories/{{ $v.Name }}">{{ $v.Name }}</a>
            </li>
            {{ end }}
            </ul>
        </div>
    </div>
    <div class="row">
        <div class="col">
            <!-- Google Adsense ディスプレイ -->
            <script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-5146230866088201" crossorigin="anonymous"></script>
            <ins class="adsbygoogle" style="display:block" data-ad-client="ca-pub-5146230866088201" data-ad-slot="3559376296" data-ad-format="auto" data-full-width-responsive="true"></ins>
            <script>(adsbygoogle = window.adsbygoogle || []).push({});</script>
        </div>
    </div>
</div>
{{ template "pagination" .Categories.Pagination }}
{{ end }}