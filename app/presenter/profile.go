package presenter

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/bmf-san/bmf-tech-client/app/model"
)

// ExecuteProfileIndex responses a index template.
func (pt *Presenter) ExecuteProfileIndex(w http.ResponseWriter, r *http.Request) error {
	fm := template.FuncMap{
		"year": pt.year,
		"isAd": pt.IsAd,
	}
	u := os.Getenv("BASE_URL")
	m := &model.Meta{
		Title:         "bmf-tech.com - プロフィール",
		Canonical:     u,
		Description:   "bmf-tech.comのプロフィールページです。",
		OGTitle:       "bmf-tech.com - プロフィール",
		OGDescription: "bmf-tech.comのプロフィールページです。",
		OGURL:         u,
		OGType:        "website",
		OGImage:       "",
		OGSiteName:    "bmf-tech",
		OGLocale:      "ja_JP",
		TwitterCard:   "summary",
		TwitterSite:   "@bmf_san",
		NoIndex:       false,
	}

	var buf bytes.Buffer

	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(pt.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/profile/index.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", map[string]interface{}{"Meta": m}); err != nil {
		if err := pt.ExecuteError(w, http.StatusInternalServerError); err != nil {
			return err
		}
		return err
	}

	if _, err := buf.WriteTo(w); err != nil {
		return err
	}

	return nil
}
