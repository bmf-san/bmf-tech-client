{{ define "pagination" }}
<div class="pagination">
    <ul class="pagination-list">
        {{ range $i, $_ := .PaginationPageCount }}
            {{ if gt $i 0 }}
            <li><a {{if eq $.PaginationPage $i}} class="pagination-active" {{end}} href="?page={{ $i }}&limit={{ $.PaginationLimit }}"><span>{{ $i }}</span></a></li>
            {{ end }}
        {{ end }}
    </ul>
</div>
{{ end }}