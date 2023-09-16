package presenter

import (
	"bytes"
	"html/template"

	"github.com/bmf-san/bmf-tech-client/app/model"
)

// ErrorData is a data for template.
type ErrorData struct {
	Code    int
	Message string
}

// ExecuteError responses a error template.
func (p *Presenter) ExecuteError(buf *bytes.Buffer, code int) (*bytes.Buffer, error) {
	e := &ErrorData{
		Code:    code,
		Message: "すみません！エラーです！",
	}
	fm := template.FuncMap{
		"year": p.year,
		"isAd": p.IsAd,
	}
	m := &model.Meta{
		Title:   "bmf-tech.com - エラー",
		NoIndex: true,
	}
	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(p.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/error/index.tpl"))
	if err := tpl.ExecuteTemplate(buf, "base", map[string]interface{}{"Meta": m, "ErrorData": e}); err != nil {
		return nil, err
	}

	return buf, nil
}
