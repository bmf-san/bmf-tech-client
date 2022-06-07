{{ define "headline" }}
<div class="headline background-color-main color-text-reverse">
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="text-align-center color-text-reverse">{{ .Post.Post.Title }}</h1>
            </div>
        </div>
    </div>
</div>
{{ end }}
{{ define "body" }}
<section>
    <div class="container-readable margin-top-2rem">
        <div class="row">
            <div class="col text">
                <div class="text-align-right">
                    <span class="article-date"><span class="article-date">{{ .Post.Post.UpdatedAt.Format "2006年1月2日 01:02:03" }} 更新</span>
                </div>
                <div class="text-align-right margin-top-1rem">
                    <a href="/posts/categories/{{ .Post.Post.Category.Name }}">{{ .Post.Post.Category.Name }}</a>
                </div>
                <div class="text-align-right margin-top-1rem">
                {{ range .Post.Post.Tags }}
                    <a class="tag" href="/posts/tags/{{ .Name }}">{{ .Name }}</a>
                {{ end }}
                </div>
				<div>

                {{ .Post.Post.HTMLBody | unescape }}
				</div>
            </div>
        </div>
    </div>
</section>
{{ end }}