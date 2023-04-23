package presenter

import (
	"html/template"

	"github.com/bmf-san/bmf-tech-client/app/model"
)

// Pagination is a data for pagination template.
type Pagination struct {
	Pager       *model.Pagination
	QueryParams template.URL
}
