package controller

import (
	"net/http"

	"github.com/bmf-san/bmf-tech-client/app/logger"
	"github.com/bmf-san/bmf-tech-client/app/presenter"
)

// A PrivacyPolicyController is a controller for a privacy policy.
type PrivacyPolicyController struct {
	Logger    *logger.Logger
	Presenter *presenter.Presenter
}

// NewPrivacyPolicyController creates a PrivacyPolicyController.
func NewPrivacyPolicyController(logger *logger.Logger, presenter *presenter.Presenter) *PrivacyPolicyController {
	return &PrivacyPolicyController{
		Logger:    logger,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (hc *PrivacyPolicyController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := hc.Presenter.ExecutePrivacyPolicyIndex(w, r); err != nil {
			hc.Logger.Error(err.Error())
			hc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
	})
}
