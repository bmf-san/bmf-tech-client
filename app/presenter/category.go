package presenter

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/bmf-san/bmf-tech-client/app/model"
)

// CategoryIndex is a data for index template.
type CategoryIndex struct {
	Categories *model.Categories
	Pagination *Pagination
}

// ExecuteCategoryIndex responses a index template.
func (p *Presenter) ExecuteCategoryIndex(w http.ResponseWriter, r *http.Request, c *CategoryIndex) error {
	fm := template.FuncMap{
		"year": p.year,
		"isAd": p.IsAd,
	}
	u := os.Getenv("BASE_URL") + "/categories"
	m := &model.Meta{
		Title:         "bmf-tech.com - カテゴリ一覧",
		Canonical:     u,
		Description:   "bmf-tech.comのカテゴリー一覧ページです。",
		OGTitle:       "bmf-tech.com - カテゴリ一覧",
		OGDescription: "bmf-tech.comのカテゴリー一覧ページです。",
		OGURL:         u,
		OGType:        "article",
		OGImage:       "",
		OGSiteName:    "bmf-tech",
		OGLocale:      "ja_JP",
		TwitterCard:   "summary",
		TwitterSite:   "@bmf_san",
		NoIndex:       false,
	}

	var buf bytes.Buffer

	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(p.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/category/index.tpl", "templates/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", map[string]interface{}{"Meta": m, "Categories": c}); err != nil {
		if err := p.ExecuteError(w, http.StatusInternalServerError); err != nil {
			return err
		}
		return err
	}

	if _, err := buf.WriteTo(w); err != nil {
		return err
	}

	return nil
}
