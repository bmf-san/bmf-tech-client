package presenter

import (
	"embed"
	"html/template"
	"net/url"
	"regexp"
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

func (p *Presenter) Unescape(text string) template.HTML {
	return template.HTML(text)
}

func (p *Presenter) year() int {
	return time.Now().Year()
}

func (p *Presenter) CurrentURLWithoutQuery(u *url.URL) string {
	u, err := url.Parse(u.String())
	if err != nil {
		return ""
	}
	u.RawQuery = ""
	return u.String()
}

func (p *Presenter) BaseURL(u *url.URL) string {
	u, err := url.Parse(u.String())
	if err != nil {
		return ""
	}
	u.RawPath = ""
	u.RawQuery = ""
	return u.String()
}

const regex = `<.*?>`

func (p *Presenter) StripTags(s string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(s, "")
}
