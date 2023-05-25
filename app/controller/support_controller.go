package controller

import (
	"net/http"

	"github.com/bmf-san/bmf-tech-client/app/logger"
	"github.com/bmf-san/bmf-tech-client/app/presenter"
)

// A SupportController is a controller for a support.
type SupportController struct {
	Logger    *logger.Logger
	Presenter *presenter.Presenter
}

// NewSupportController creates a SupportController.
func NewSupportController(logger *logger.Logger, presenter *presenter.Presenter) *SupportController {
	return &SupportController{
		Logger:    logger,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (hc *SupportController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := hc.Presenter.ExecuteSupportIndex(w, r); err != nil {
			hc.Logger.Error(err.Error())
			hc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
	})
}
