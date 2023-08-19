// TODO: Implement csrf
package middleware

import (
	"net/http"
	"runtime"

	"github.com/bmf-san/bmf-tech-client/app/logger"
	"github.com/bmf-san/bmf-tech-client/app/presenter"
)

// Middelware represents the singular of middleware.
type Middleware struct {
	logger    *logger.Logger
	presenter *presenter.Presenter
}

// NewMiddleware creates a middleware.
func NewMiddleware(l *logger.Logger, p *presenter.Presenter) *Middleware {
	return &Middleware{
		logger:    l,
		presenter: p,
	}
}

// Recovery is a middleware for recovering from panic.
func (mw *Middleware) Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case string:
					mw.logger.Error("[panic]" + e)
				case runtime.Error:
					mw.logger.Error("[panic] " + e.Error())
				case error:
					mw.logger.Error("[panic] " + e.Error())
				default:
					mw.logger.Error("[panic] " + e.(string))
				}
				mw.presenter.Error(w, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// TODO: Implement later
// const (
// 	ctxCSRFToken = "csrf-token"
// )

// // CSRF is a middelware for CSRF.
// func (a *Asset) CSRF(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodGet:
// 			token, err := a.redisHandler.SetCSRFToken()
// 			if err != nil {
// 				a.logger.Error(err.Error())
// 				a.Presenter.Error(w, http.StatusInternalServerError)
// 				return
// 			}
// 			ctx := context.WithValue(r.Context(), ctxCSRFToken, token)
// 			next.ServeHTTP(w, r.WithContext(ctx))
// 			return

// 		case http.MethodPost:
// 			r.ParseForm()
// 			token := r.Form.Get("csrf_token")
// 			if token == "" {
// 				a.logger.Error("CSRF token is invalid")
// 				a.Presenter.Error(w, http.StatusInternalServerError)
// 				return
// 			}

// 			if err := a.redisHandler.HasCSRFToken(token); err != nil {
// 				a.logger.Error(err.Error())
// 				a.Presenter.Error(w, http.StatusInternalServerError)
// 				return
// 			}

// 			next.ServeHTTP(w, r)
// 			return
// 		default:
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 	}
// }
