package presenter

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/bmf-san/bmf-tech-client/app/model"
)

// TagIndex is a data for index template.
type TagIndex struct {
	Tags       *model.Tags
	Pagination *Pagination
}

// ExecuteTagIndex responses a index template.
func (p *Presenter) ExecuteTagIndex(w http.ResponseWriter, r *http.Request, t *TagIndex) error {
	fm := template.FuncMap{
		"year": p.year,
		"isAd": p.IsAd,
	}
	u := os.Getenv("BASE_URL") + "/tags"
	m := &model.Meta{
		Title:         "bmf-tech.com - タグ一覧",
		Canonical:     u,
		Description:   "bmf-tech.comのタグ一覧ページです。",
		OGTitle:       "bmf-tech.com - タグ一覧",
		OGDescription: "bmf-tech.comのタグ一覧ページです。",
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

	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(p.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/tag/index.tpl", "templates/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", map[string]interface{}{"Meta": m, "Tags": t}); err != nil {
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
