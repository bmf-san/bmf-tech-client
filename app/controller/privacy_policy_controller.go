package controller

import (
	"bytes"
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
func (pc *PrivacyPolicyController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		code := http.StatusOK
		buf, err := pc.Presenter.ExecutePrivacyPolicyIndex(buf, r)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err = pc.Presenter.ExecuteError(buf, http.StatusInternalServerError)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
		}

		bufWriteTo(buf, w, code)
	})
}
