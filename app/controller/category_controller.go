package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"log/slog"

	"github.com/bmf-san/bmf-tech-client/app/api"
	"github.com/bmf-san/bmf-tech-client/app/model"
	"github.com/bmf-san/bmf-tech-client/app/presenter"
)

// A CategoryController is a controller for category.
type CategoryController struct {
	Logger    *slog.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewCategoryController creates a CategoryController.
func NewCategoryController(logger *slog.Logger, client *api.Client, presenter *presenter.Presenter) *CategoryController {
	return &CategoryController{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (cc *CategoryController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, _, err := cc.Client.GetPageAndLimit(r)
		if err != nil {
			cc.Logger.Error(err.Error())
			if err := cc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				cc.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		resp, err := cc.Client.GetCategories(page, 100)
		if err != nil {
			cc.Logger.Error(err.Error())
			if err := cc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				cc.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			cc.Logger.Error(err.Error())
			if err := cc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				cc.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		var categories model.Categories

		if err := json.Unmarshal(body, &categories); err != nil {
			cc.Logger.Error(err.Error())
			if err := cc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				cc.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			cc.Logger.Error(err.Error())
			if err := cc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				cc.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if err = cc.Presenter.ExecuteCategoryIndex(w, r, &presenter.CategoryIndex{
			Categories: &categories,
			Pagination: &presenter.Pagination{
				Pager:       &pagination,
				QueryParams: "",
			},
		}); err != nil {
			cc.Logger.Error(err.Error())
			if err := cc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				cc.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	})
}
