package presenter

import (
	"html/template"
	"net/http"

	"github.com/bmf-san/bmf-tech-client/app/model"
)

// TagIndex is a data for index template.
type TagIndex struct {
	Tags       *model.Tags
	Pagination *model.Pagination
}

// ExecuteTagIndex responses a index template.
func (p *Presenter) ExecuteTagIndex(w http.ResponseWriter, r *http.Request, t *TagIndex) error {
	fm := template.FuncMap{
		"year": p.year,
	}
	m := &model.Meta{
		Canonical:     p.CurrentURLWithoutQuery(r.URL),
		Description:   "タグ一覧",
		OGTitle:       "タグ一覧",
		OGDescription: "タグ一覧",
		OGURL:         p.CurrentURLWithoutQuery(r.URL),
		OGType:        "article",
		OGImage:       "",
		OGSiteName:    "bmf-tech",
		OGLocale:      "ja_JP",
		TwitterCard:   "summary",
		TwitterSite:   "@bmf_san",
		NoIndex:       false,
	}
	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(p.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/tag/index.tpl", "templates/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", map[string]interface{}{"Meta": m, "Tags": t}); err != nil {
		return err
	}
	return nil
}
