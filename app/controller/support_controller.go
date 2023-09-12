package controller

import (
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
		if err := sc.Presenter.ExecuteSupportIndex(w, r); err != nil {
			sc.Logger.Error(err.Error())
			if err := sc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				sc.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	})
}
