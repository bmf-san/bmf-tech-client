package presenter

import (
	"html/template"
	"net/http"
)

// ExecuteHomeIndex responses a index template.
func (pt *Presenter) ExecuteHomeIndex(w http.ResponseWriter, p *PostIndex) error {
	fm := template.FuncMap{
		"year": pt.year,
	}
	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(pt.templates, "templates/layout/base.tpl", "templates/home/index.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", p); err != nil {
		return err
	}
	return nil
}
