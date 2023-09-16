package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"log/slog"

	"github.com/bmf-san/bmf-tech-client/app/api"
	"github.com/bmf-san/bmf-tech-client/app/model"
	"github.com/bmf-san/bmf-tech-client/app/presenter"
)

// A HomeController is a controller for a home.
type HomeController struct {
	Logger    *slog.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewHomeController creates a HomeController.
func NewHomeController(logger *slog.Logger, client *api.Client, presenter *presenter.Presenter) *HomeController {
	return &HomeController{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (hc *HomeController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		code := http.StatusOK
		page, limit, err := hc.Client.GetPageAndLimit(r)
		if err != nil {
			hc.Logger.Error(err.Error())
			code = http.StatusInternalServerError
			buf, err := hc.Presenter.ExecuteError(buf, code)
			if err != nil {
				hc.Logger.Error(err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		resp, err := hc.Client.GetPosts(page, limit)
		if err != nil {
			hc.Logger.Error(err.Error())
			code = http.StatusInternalServerError
			buf, err := hc.Presenter.ExecuteError(buf, code)
			if err != nil {
				hc.Logger.Error(err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			hc.Logger.Error(err.Error())
			code = http.StatusInternalServerError
			buf, err := hc.Presenter.ExecuteError(buf, code)
			if err != nil {
				hc.Logger.Error(err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		var posts model.Posts

		if err := json.Unmarshal(body, &posts); err != nil {
			hc.Logger.Error(err.Error())
			code = http.StatusInternalServerError
			buf, err := hc.Presenter.ExecuteError(buf, code)
			if err != nil {
				hc.Logger.Error(err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		buf, err = hc.Presenter.ExecuteHomeIndex(buf, r, &presenter.PostIndex{
			Posts: &posts,
		})
		if err != nil {
			hc.Logger.Error(err.Error())
			code = http.StatusInternalServerError
			buf, err = hc.Presenter.ExecuteError(buf, code)
			if err != nil {
				hc.Logger.Error(err.Error())
			}
		}

		bufWriteTo(buf, w, code)
	})
}
