package controller

import (
	"bytes"
	"net/http"

	"log/slog"

	"github.com/bmf-san/bmf-tech-client/app/presenter"
)

// A SupportController is a controller for a support.
type SupportController struct {
	Logger    *slog.Logger
	Presenter *presenter.Presenter
}

// NewSupportController creates a SupportController.
func NewSupportController(logger *slog.Logger, presenter *presenter.Presenter) *SupportController {
	return &SupportController{
		Logger:    logger,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (sc *SupportController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		code := http.StatusOK
		buf, err := sc.Presenter.ExecuteSupportIndex(buf, r)
		if err != nil {
			sc.Logger.Error(err.Error())
			code = http.StatusInternalServerError
			buf, err = sc.Presenter.ExecuteError(buf, code)
			if err != nil {
				sc.Logger.Error(err.Error())
			}
		}

		bufWriteTo(buf, w, code)
	})
}
