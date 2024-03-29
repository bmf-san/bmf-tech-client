package presenter

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/bmf-san/bmf-tech-client/app/model"
)

// ExecutePrivacyPolicyIndex responses a index template.
func (pt *Presenter) ExecutePrivacyPolicyIndex(buf *bytes.Buffer, r *http.Request) (*bytes.Buffer, error) {
	fm := template.FuncMap{
		"year": pt.year,
		"isAd": pt.IsAd,
	}
	u := os.Getenv("BASE_URL")
	m := &model.Meta{
		Title:         "bmf-tech.com - プライバシーポリシー",
		Canonical:     u,
		Description:   "bmf-tech.comのプライバシーポリシーページです。",
		OGTitle:       "bmf-tech.com - プライバシーポリシー",
		OGDescription: "bmf-tech.comのプライバシーポリシーページです。",
		OGURL:         u,
		OGType:        "website",
		OGImage:       "",
		OGSiteName:    "bmf-tech",
		OGLocale:      "ja_JP",
		TwitterCard:   "summary",
		TwitterSite:   "@bmf_san",
		NoIndex:       false,
	}

	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(pt.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/privacy_policy/index.tpl"))
	if err := tpl.ExecuteTemplate(buf, "base", map[string]interface{}{"Meta": m}); err != nil {
		return nil, err
	}

	return buf, nil
}
