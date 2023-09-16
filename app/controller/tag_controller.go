package controller

import (
	"bytes"
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
		buf := new(bytes.Buffer)
		code := http.StatusOK
		page, _, err := tc.Client.GetPageAndLimit(r)
		if err != nil {
			tc.Logger.Error(err.Error())
			buf, err := tc.Presenter.ExecuteError(buf, http.StatusInternalServerError)
			if err != nil {
				tc.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			bufWriteTo(buf, w, code)
			return
		}

		resp, err := tc.Client.GetTags(page, 100)
		if err != nil {
			tc.Logger.Error(err.Error())
			buf, err := tc.Presenter.ExecuteError(buf, http.StatusInternalServerError)
			if err != nil {
				tc.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			bufWriteTo(buf, w, code)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			tc.Logger.Error(err.Error())
			buf, err := tc.Presenter.ExecuteError(buf, http.StatusInternalServerError)
			if err != nil {
				tc.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			bufWriteTo(buf, w, code)
			return
		}

		var tags model.Tags

		if err := json.Unmarshal(body, &tags); err != nil {
			tc.Logger.Error(err.Error())
			buf, err := tc.Presenter.ExecuteError(buf, http.StatusInternalServerError)
			if err != nil {
				tc.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			bufWriteTo(buf, w, code)
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			tc.Logger.Error(err.Error())
			buf, err := tc.Presenter.ExecuteError(buf, http.StatusInternalServerError)
			if err != nil {
				tc.Logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			bufWriteTo(buf, w, code)
			return
		}

		buf, err = tc.Presenter.ExecuteTagIndex(buf, r, &presenter.TagIndex{
			Tags: &tags,
			Pagination: &presenter.Pagination{
				Pager:       &pagination,
				QueryParams: "",
			},
		})
		if err != nil {
			tc.Logger.Error(err.Error())
			code = http.StatusInternalServerError
			buf, err = tc.Presenter.ExecuteError(buf, code)
			if err != nil {
				tc.Logger.Error(err.Error())
			}
		}
		bufWriteTo(buf, w, code)
	})
}
