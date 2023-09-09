package presenter

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/bmf-san/bmf-tech-client/app/model"
)

// ErrorData is a data for template.
type ErrorData struct {
	Code    int
	Message string
}

// ExecuteError responses a error template.
func (p *Presenter) ExecuteError(w http.ResponseWriter, code int) error {
	e := &ErrorData{
		Code:    code,
		Message: handleErrorMessage(code),
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

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, "base", map[string]interface{}{"Meta": m, "ErrorData": e}); err != nil {
		return err
	}

	w.WriteHeader(e.Code)

	if _, err := buf.WriteTo(w); err != nil {
		return err
	}

	return nil
}

// handleErrorMessage handles an error message by code.
func handleErrorMessage(code int) string {
	switch code {
	case http.StatusInternalServerError:
		return "すみません！エラーです！"
	}
	return ""
}
