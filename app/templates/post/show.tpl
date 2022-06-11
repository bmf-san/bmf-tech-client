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
                <div class="margin-top-2rem margin-bottom-2rem">
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