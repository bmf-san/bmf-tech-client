package controller

import (
	"bytes"
	"net/http"

	"github.com/bmf-san/bmf-tech-client/app/logger"
	"github.com/bmf-san/bmf-tech-client/app/presenter"
)

// A ProfileController is a controller for a profile.
type ProfileController struct {
	Logger    *logger.Logger
	Presenter *presenter.Presenter
}

// NewProfileController creates a ProfileController.
func NewProfileController(logger *logger.Logger, presenter *presenter.Presenter) *ProfileController {
	return &ProfileController{
		Logger:    logger,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (pc *ProfileController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		code := http.StatusOK
		buf, err := pc.Presenter.ExecuteProfileIndex(buf, r)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err = pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
		}
		bufWriteTo(buf, w, code)
	})
}
