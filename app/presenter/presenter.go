package presenter

import (
	"embed"
	"html/template"
	"time"
)

// Presenter represents the singular of presenter.
type Presenter struct {
	templates embed.FS
}

// NewPresenter creates a Presenter.
func NewPresenter(templates embed.FS) *Presenter {
	return &Presenter{
		templates: templates,
	}
}

func (p *Presenter) unescape(text string) template.HTML {
	return template.HTML(text)
}

func (p *Presenter) year() int {
	return time.Now().Year()
}
