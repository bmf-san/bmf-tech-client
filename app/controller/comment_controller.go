package controller

import (
	"net/http"

	"log/slog"

	"github.com/bmf-san/bmf-tech-client/app/api"
	"github.com/bmf-san/bmf-tech-client/app/presenter"
)

// A CommentControlller is a controller for a comment.
type CommentControlller struct {
	Logger    *slog.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewCommentController creates a NewCommentController.
func NewCommentController(logger *slog.Logger, client *api.Client, presenter *presenter.Presenter) *CommentControlller {
	return &CommentControlller{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Store displays a listing of the resource.
func (cc *CommentControlller) Store(w http.ResponseWriter, r *http.Request) {
	// TODO: implement later.
}
