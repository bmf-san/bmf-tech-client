{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-align-center color-text-reverse">{{ .TagName }}</h1>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "body" }}
<div class="container-readable margin-top-2rem">
    {{ range $i, $v := .Posts }}
    <div class="row">
        <div class="col">
            <article>
                <span class="article-date">{{ $v.CreatedAt.Format "2006 Jan 02" }}</span>
                <h2><a href="/posts/{{ $v.Title }}">{{ $v.Title }}</a></h2>
                <div class="text-align-right margin-top-1rem">
                    <a href="/posts/categories/{{ $v.Category.Name }}">{{ $v.Category.Name }}</a>
                </div>
                <div class="text-align-right margin-top-1rem">
                    {{ range $v.Tags }}
                    <a class="tag" href="/posts/tags/{{ .Name }}">{{ .Name }}</a>
                    {{ end }}
                </div>
            </article>
        </div>
    </div>
    {{ end }}
    {{ template "pagination" .Pagination }}
</div>
{{ end }}