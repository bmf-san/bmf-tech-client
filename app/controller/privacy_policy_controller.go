package controller

import (
	"net/http"

	"log/slog"

	"github.com/bmf-san/bmf-tech-client/app/presenter"
)

// A PrivacyPolicyController is a controller for a privacy policy.
type PrivacyPolicyController struct {
	Logger    *slog.Logger
	Presenter *presenter.Presenter
}

// NewPrivacyPolicyController creates a PrivacyPolicyController.
func NewPrivacyPolicyController(logger *slog.Logger, presenter *presenter.Presenter) *PrivacyPolicyController {
	return &PrivacyPolicyController{
		Logger:    logger,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (pc *PrivacyPolicyController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := pc.Presenter.ExecutePrivacyPolicyIndex(w, r); err != nil {
			pc.Logger.Error(err.Error())
			if err := pc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				pc.Logger.Error(err.Error())
				w.Write([]byte(err.Error()))
			}
			return
		}
	})
}
