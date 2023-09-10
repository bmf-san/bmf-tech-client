package controller

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/bmf-san/bmf-tech-client/app/api"
	"github.com/bmf-san/bmf-tech-client/app/model"
	"github.com/bmf-san/bmf-tech-client/app/presenter"
)

// A TagController is a controller for a tag.
type TagController struct {
	Logger    *slog.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewTagController creates a TagController.
func NewTagController(logger *slog.Logger, client *api.Client, presenter *presenter.Presenter) *TagController {
	return &TagController{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (tc *TagController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, _, err := tc.Client.GetPageAndLimit(r)
		if err != nil {
			tc.Logger.Error(err.Error())
			if err := tc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				tc.Logger.Error(err.Error())
				w.Write([]byte(err.Error()))
			}
			return
		}

		resp, err := tc.Client.GetTags(page, 100)
		if err != nil {
			tc.Logger.Error(err.Error())
			if err := tc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				tc.Logger.Error(err.Error())
				w.Write([]byte(err.Error()))
			}
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			tc.Logger.Error(err.Error())
			if err := tc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				tc.Logger.Error(err.Error())
				w.Write([]byte(err.Error()))
			}
			return
		}

		var tags model.Tags

		if err := json.Unmarshal(body, &tags); err != nil {
			tc.Logger.Error(err.Error())
			if err := tc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				tc.Logger.Error(err.Error())
				w.Write([]byte(err.Error()))
			}
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			tc.Logger.Error(err.Error())
			if err := tc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				tc.Logger.Error(err.Error())
				w.Write([]byte(err.Error()))
			}
			return
		}

		if err = tc.Presenter.ExecuteTagIndex(w, r, &presenter.TagIndex{
			Tags: &tags,
			Pagination: &presenter.Pagination{
				Pager:       &pagination,
				QueryParams: "",
			},
		}); err != nil {
			tc.Logger.Error(err.Error())
			if err := tc.Presenter.ExecuteError(w, http.StatusInternalServerError); err != nil {
				tc.Logger.Error(err.Error())
				w.Write([]byte(err.Error()))
			}
			return
		}
	})
}
