{{ define "posts" }}
{{ range $i, $v := .Posts }}
<div class="row">
    <div class="col">
        <article>
            <h1 class="font-size-large"><b><a href="/posts/{{ $v.Title }}">{{ $v.Title }}</a></b></h1>
            <p class="margin-0rem"><a href="#"><a href="/posts/categories/{{ $v.Category.Name }}">{{ $v.Category.Name }}</a></p>
            <span class="article-date">{{ $v.CreatedAt.Format "2006 Jan 02" }}</span>
            <p>{{ striptags $v.HTMLBody | summary }}</p>
            <div>
            {{ range $v.Tags }}
                <a class="tag" href="/posts/tags/{{ .Name }}">{{ .Name }}</a>
            {{ end }}
            </div>
        </article>
    </div>
</div>
{{ end }}
{{ end }}