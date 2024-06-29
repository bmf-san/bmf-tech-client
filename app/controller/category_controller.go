package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/bmf-san/bmf-tech-client/app/api"
	"github.com/bmf-san/bmf-tech-client/app/logger"
	"github.com/bmf-san/bmf-tech-client/app/model"
	"github.com/bmf-san/bmf-tech-client/app/presenter"
)

// A CategoryController is a controller for category.
type CategoryController struct {
	Logger    *logger.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewCategoryController creates a CategoryController.
func NewCategoryController(logger *logger.Logger, client *api.Client, presenter *presenter.Presenter) *CategoryController {
	return &CategoryController{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (cc *CategoryController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		code := http.StatusOK
		page, _, err := cc.Client.GetPageAndLimit(r)
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := cc.Presenter.ExecuteError(buf, code)
			if err != nil {
				cc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		resp, err := cc.Client.GetCategories(r.Context(), page, 100)
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := cc.Presenter.ExecuteError(buf, code)
			if err != nil {
				cc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := cc.Presenter.ExecuteError(buf, code)
			if err != nil {
				cc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		var categories model.Categories

		if err := json.Unmarshal(body, &categories); err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := cc.Presenter.ExecuteError(buf, code)
			if err != nil {
				cc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := cc.Presenter.ExecuteError(buf, code)
			if err != nil {
				cc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		buf, err = cc.Presenter.ExecuteCategoryIndex(buf, r, &presenter.CategoryIndex{
			Categories: &categories,
			Pagination: &presenter.Pagination{
				Pager:       &pagination,
				QueryParams: "",
			},
		})
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err = cc.Presenter.ExecuteError(buf, code)
			if err != nil {
				cc.Logger.ErrorContext(r.Context(), err.Error())
			}
		}

		bufWriteTo(buf, w, code)
	})
}
