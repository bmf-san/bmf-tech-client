package presenter

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/bmf-san/bmf-tech-client/app/model"
)

// PostIndex is a data for index template.
type PostIndex struct {
	Posts      *model.Posts
	Pagination *Pagination
}

// PostIndexBySearch is a data for index template by search.
type PostIndexBySearch struct {
	Keyword    string
	Posts      *model.Posts
	Pagination *Pagination
}

// PostIndexByCategory is a data for index template by category.
type PostIndexByCategory struct {
	CategoryName string
	Posts        *model.Posts
	Pagination   *Pagination
}

// PostIndexByTag is a data for index template by tag.
type PostIndexByTag struct {
	TagName    string
	Posts      *model.Posts
	Pagination *Pagination
}

// PostShow is a data for show template.
type PostShow struct {
	Post             *model.Post
	TwitterShareURL  string
	FacebookShareURL string
	HatenaShareURL   string
	SupportURL       string
}

// ExecutePostIndex responses a index template.
func (pt *Presenter) ExecutePostIndex(buf *bytes.Buffer, r *http.Request, p *PostIndex) (*bytes.Buffer, error) {
	fm := template.FuncMap{
		"year":      pt.year,
		"striptags": pt.StripTags,
		"summary":   pt.Summary,
		"isAd":      pt.IsAd,
	}
	u := os.Getenv("BASE_URL") + "/posts"
	m := &model.Meta{
		Title:         "bmf-tech.com - 記事一覧",
		Canonical:     u,
		Description:   "bmf-tech.comの記事一覧ページです。",
		OGTitle:       "bmf-tech.com - 記事一覧",
		OGDescription: "bmf-tech.comの記事一覧ページです。",
		OGURL:         u,
		OGType:        "article",
		OGImage:       "",
		OGSiteName:    "bmf-tech",
		OGLocale:      "ja_JP",
		TwitterCard:   "summary",
		TwitterSite:   "@bmf_san",
		NoIndex:       false,
	}

	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(pt.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/post/index.tpl", "templates/partial/pagination.tpl", "templates/partial/posts.tpl", "templates/partial/search.tpl"))
	if err := tpl.ExecuteTemplate(buf, "base", map[string]interface{}{"Meta": m, "Posts": p}); err != nil {
		return nil, err
	}

	return buf, nil
}

// ExecutePostIndexByKeyword responses a index template by keyword.
func (pt *Presenter) ExecutePostIndexByKeyword(buf *bytes.Buffer, r *http.Request, p *PostIndexBySearch) (*bytes.Buffer, error) {
	fm := template.FuncMap{
		"year":      pt.year,
		"striptags": pt.StripTags,
		"summary":   pt.Summary,
		"isAd":      pt.IsAd,
	}
	u := os.Getenv("BASE_URL") + "/posts"
	m := &model.Meta{
		Title:         "bmf-tech.com - 検索記事一覧",
		Canonical:     u,
		Description:   "検索記事一覧：" + p.Keyword,
		OGTitle:       "検索記事一覧：" + p.Keyword,
		OGDescription: "検索記事一覧：" + p.Keyword,
		OGURL:         u,
		OGType:        "article",
		OGImage:       "",
		OGSiteName:    "bmf-tech",
		OGLocale:      "ja_JP",
		TwitterCard:   "summary",
		TwitterSite:   "@bmf_san",
		NoIndex:       false,
	}

	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(pt.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/post/search.tpl", "templates/partial/pagination.tpl", "templates/partial/posts.tpl", "templates/partial/search.tpl"))
	if err := tpl.ExecuteTemplate(buf, "base", map[string]interface{}{"Meta": m, "Posts": p}); err != nil {
		return nil, err
	}

	return buf, nil
}

// ExecutePostIndexByCategory responses a index template by category.
func (pt *Presenter) ExecutePostIndexByCategory(buf *bytes.Buffer, r *http.Request, p *PostIndexByCategory) (*bytes.Buffer, error) {
	fm := template.FuncMap{
		"year":      pt.year,
		"striptags": pt.StripTags,
		"summary":   pt.Summary,
		"isAd":      pt.IsAd,
	}
	u := os.Getenv("BASE_URL") + "/posts/categories/" + p.CategoryName
	m := &model.Meta{
		Title:         "bmf-tech.com - カテゴリ別記事一覧",
		Canonical:     u,
		Description:   "カテゴリ別記事一覧：" + p.CategoryName,
		OGTitle:       "カテゴリ別記事一覧：" + p.CategoryName,
		OGDescription: "カテゴリ別記事一覧：" + p.CategoryName,
		OGURL:         u,
		OGType:        "article",
		OGImage:       "",
		OGSiteName:    "bmf-tech",
		OGLocale:      "ja_JP",
		TwitterCard:   "summary",
		TwitterSite:   "@bmf_san",
		NoIndex:       false,
	}

	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(pt.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/post/category.tpl", "templates/partial/pagination.tpl", "templates/partial/posts.tpl"))
	if err := tpl.ExecuteTemplate(buf, "base", map[string]interface{}{"Meta": m, "Posts": p}); err != nil {
		return nil, err
	}

	return buf, nil
}

// ExecutePostIndexByTag responses a index template by tag.
func (pt *Presenter) ExecutePostIndexByTag(buf *bytes.Buffer, r *http.Request, p *PostIndexByTag) (*bytes.Buffer, error) {
	fm := template.FuncMap{
		"year":      pt.year,
		"striptags": pt.StripTags,
		"summary":   pt.Summary,
		"isAd":      pt.IsAd,
	}
	u := os.Getenv("BASE_URL") + "/posts/tags/" + p.TagName
	m := &model.Meta{
		Title:         "bmf-tech.com - タグ別記事一覧",
		Canonical:     u,
		Description:   "タグ別記事一覧：" + p.TagName,
		OGTitle:       "タグ別記事一覧：" + p.TagName,
		OGDescription: "タグ別記事一覧：" + p.TagName,
		OGURL:         u,
		OGType:        "article",
		OGImage:       "",
		OGSiteName:    "bmf-tech",
		OGLocale:      "ja_JP",
		TwitterCard:   "summary",
		TwitterSite:   "@bmf_san",
		NoIndex:       false,
	}

	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(pt.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/post/tag.tpl", "templates/partial/pagination.tpl", "templates/partial/posts.tpl"))
	if err := tpl.ExecuteTemplate(buf, "base", map[string]interface{}{"Meta": m, "Posts": p}); err != nil {
		return nil, err
	}

	return buf, nil
}

// ExecutePostShow responses a show template by tag.
func (pt *Presenter) ExecutePostShow(buf *bytes.Buffer, r *http.Request, p *PostShow) (*bytes.Buffer, error) {
	fm := template.FuncMap{
		"year":     pt.year,
		"unescape": pt.Unescape,
		"isAd":     pt.IsAd,
	}
	s := pt.Summary(pt.StripTags(p.Post.HTMLBody))
	u := os.Getenv("BASE_URL") + "/posts/" + p.Post.Title
	m := &model.Meta{
		Title:         "bmf-tech.com - " + p.Post.Title,
		Canonical:     u,
		Description:   s,
		OGTitle:       p.Post.Title,
		OGDescription: s,
		OGURL:         u,
		OGType:        "article",
		OGImage:       "",
		OGSiteName:    "bmf-tech",
		OGLocale:      "ja_JP",
		TwitterCard:   "summary",
		TwitterSite:   "@bmf_san",
		NoIndex:       false,
	}

	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(pt.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/post/show.tpl"))
	if err := tpl.ExecuteTemplate(buf, "base", map[string]interface{}{"Meta": m, "Post": p}); err != nil {
		return nil, err
	}

	return buf, nil
}
