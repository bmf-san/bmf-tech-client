package presenter

import (
	"html/template"
	"net/http"

	"github.com/bmf-san/bmf-tech-client/app/model"
)

// CategoryIndex is a data for index template.
type CategoryIndex struct {
	Categories *model.Categories
	Pagination *model.Pagination
}

// ExecuteCategoryIndex responses a index template.
func (p *Presenter) ExecuteCategoryIndex(w http.ResponseWriter, r *http.Request, c *CategoryIndex) error {
	fm := template.FuncMap{
		"year": p.year,
	}
	m := &model.Meta{
		Canonical:     p.CurrentURLWithoutQuery(r.URL),
		Description:   "カテゴリ一覧",
		OGTitle:       "カテゴリ一覧",
		OGDescription: "カテゴリ一覧",
		OGURL:         p.CurrentURLWithoutQuery(r.URL),
		OGType:        "article",
		OGImage:       "",
		OGSiteName:    "bmf-tech",
		OGLocale:      "ja_JP",
		TwitterCard:   "summary",
		TwitterSite:   "@bmf_san",
		NoIndex:       false,
	}
	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(p.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/category/index.tpl", "templates/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", map[string]interface{}{"Meta": m, "Categories": c}); err != nil {
		return err
	}
	return nil
}
