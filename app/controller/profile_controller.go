package controller

import (
	"net/http"

	"log/slog"

	"github.com/bmf-san/bmf-tech-client/app/presenter"
)

// A ProfileController is a controller for a profile.
type ProfileController struct {
	Logger    *slog.Logger
	Presenter *presenter.Presenter
}

// NewProfileController creates a ProfileController.
func NewProfileController(logger *slog.Logger, presenter *presenter.Presenter) *ProfileController {
	return &ProfileController{
		Logger:    logger,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (pc *ProfileController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := pc.Presenter.ExecuteProfileIndex(w, r); err != nil {
			pc.Logger.Error(err.Error())
			if err := pc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				pc.Logger.Error(err.Error())
				w.Write([]byte(err.Error()))
			}
			return
		}
	})
}
