{{ define "posts" }}
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
{{ end }}