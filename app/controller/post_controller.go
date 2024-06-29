package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/bmf-san/bmf-tech-client/app/api"
	"github.com/bmf-san/bmf-tech-client/app/logger"
	"github.com/bmf-san/bmf-tech-client/app/model"
	"github.com/bmf-san/bmf-tech-client/app/presenter"
	"github.com/bmf-san/goblin"
)

// A PostController is a controller for a post.
type PostController struct {
	Logger    *logger.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewPostController creates a NewPostController.
func NewPostController(logger *logger.Logger, client *api.Client, presenter *presenter.Presenter) *PostController {
	return &PostController{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (pc *PostController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		code := http.StatusOK
		page, limit, err := pc.Client.GetPageAndLimit(r)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		resp, err := pc.Client.GetPosts(r.Context(), page, limit)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		var posts model.Posts

		if err := json.Unmarshal(body, &posts); err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		buf, err = pc.Presenter.ExecutePostIndex(buf, r, &presenter.PostIndex{
			Posts: &posts,
			Pagination: &presenter.Pagination{
				Pager:       &pagination,
				QueryParams: "",
			},
		})
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

// IndexByKeyword displays a listing of the resource.
func (pc *PostController) IndexByKeyword() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		code := http.StatusOK
		page, limit, err := pc.Client.GetPageAndLimit(r)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		keyword := r.URL.Query().Get("keyword")
		resp, err := pc.Client.GetPostsByKeyword(r.Context(), keyword, page, limit)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		var posts model.Posts

		if err := json.Unmarshal(body, &posts); err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		buf, err = pc.Presenter.ExecutePostIndexByKeyword(buf, r, &presenter.PostIndexBySearch{
			Keyword: keyword,
			Posts:   &posts,
			Pagination: &presenter.Pagination{
				Pager:       &pagination,
				QueryParams: template.URL(fmt.Sprintf("keyword=%s", keyword)),
			},
		})
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

// IndexByCategory displays a listing of the resource.
func (pc *PostController) IndexByCategory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		code := http.StatusOK
		page, limit, err := pc.Client.GetPageAndLimit(r)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		name := goblin.GetParam(r.Context(), "name")
		resp, err := pc.Client.GetPostsByCategory(r.Context(), name, page, limit)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		var posts model.Posts

		if err := json.Unmarshal(body, &posts); err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		buf, err = pc.Presenter.ExecutePostIndexByCategory(buf, r, &presenter.PostIndexByCategory{
			CategoryName: name,
			Posts:        &posts,
			Pagination: &presenter.Pagination{
				Pager:       &pagination,
				QueryParams: "",
			},
		})
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

// IndexByTag displays a listing of the resource.
func (pc *PostController) IndexByTag() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		code := http.StatusOK
		page, limit, err := pc.Client.GetPageAndLimit(r)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		name := goblin.GetParam(r.Context(), "name")
		resp, err := pc.Client.GetPostsByTag(r.Context(), name, page, limit)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		var posts model.Posts

		if err := json.Unmarshal(body, &posts); err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		buf, err = pc.Presenter.ExecutePostIndexByTag(buf, r, &presenter.PostIndexByTag{
			TagName: name,
			Posts:   &posts,
			Pagination: &presenter.Pagination{
				Pager:       &pagination,
				QueryParams: "",
			},
		})
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

// Show displays the specified resource.
func (pc *PostController) Show() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		code := http.StatusOK
		title := goblin.GetParam(r.Context(), "title")

		resp, err := pc.Client.GetPost(r.Context(), title)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		var post model.Post

		if err := json.Unmarshal(body, &post); err != nil {
			pc.Logger.ErrorContext(r.Context(), err.Error())
			code = http.StatusInternalServerError
			buf, err := pc.Presenter.ExecuteError(buf, code)
			if err != nil {
				pc.Logger.ErrorContext(r.Context(), err.Error())
			}
			bufWriteTo(buf, w, code)
			return
		}

		buf, err = pc.Presenter.ExecutePostShow(buf, r, &presenter.PostShow{
			Post:             &post,
			TwitterShareURL:  "https://twitter.com/share?url=" + os.Getenv("BASE_URL") + "/" + post.Title + "&text=" + post.Title + "&via=bmf_san" + "&related=bmf_san",
			FacebookShareURL: "http://www.facebook.com/share.php?u=" + os.Getenv("BASE_URL") + "/" + post.Title,
			HatenaShareURL:   "http://b.hatena.ne.jp/add?mode=confirm&url=" + os.Getenv("BASE_URL") + "/" + post.Title + "&title=" + post.Title,
			SupportURL:       os.Getenv("BASE_URL") + "/support",
		})
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
