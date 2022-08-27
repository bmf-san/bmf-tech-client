{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-align-center color-text-reverse">タグ</h1>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "body" }}
<div class="container-readable margin-top-2rem">
    <div class="row">
        <div class="col">
        {{ range $i, $v := .Tags.Tags }}
        <a class="tag margin-1rem" href="/posts/tags/{{ $v.Name }}">{{ $v.Name }}</a>
        {{ end }}
        </div>
    </div>
    <div class="row">
        <div class="col">
            <!-- Google Adsense ディスプレイ -->
            <script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-5146230866088201"crossorigin="anonymous"></script>
            <ins class="adsbygoogle" style="display:block" data-ad-client="ca-pub-5146230866088201" data-ad-slot="3559376296"data-ad-format="auto" data-full-width-responsive="true"></ins>
            <script>(adsbygoogle = window.adsbygoogle || []).push({});</script>
        </div>
    </div>
</div>
{{ template "pagination" .Tags.Pagination }}
{{ end }}