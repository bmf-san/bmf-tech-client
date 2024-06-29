package middleware

import (
	"bytes"
	"net/http"
	"runtime"

	"log/slog"

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

// Log is a middleware for logging. It logs the access log. It also adds a trace id to the context.
func (mw *Middleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mw.logger.InfoContext(r.Context(), "access log", slog.String("http_method", r.Method), slog.String("path", r.URL.Path), slog.String("remote_addr", r.RemoteAddr), slog.String("user_agent", r.UserAgent()),slog.String("trace_id", r.Header.Get("X-Trace-ID"))
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

// Recovery is a middleware for recovering from panic.
func (mw *Middleware) Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case string:
					mw.logger.ErrorContext(r.Context(), "[panic] "+e)
				case runtime.Error:
					mw.logger.ErrorContext(r.Context(), "[panic] "+e.Error())
				case error:
					mw.logger.ErrorContext(r.Context(), "[panic] "+e.Error())
				default:
					mw.logger.ErrorContext(r.Context(), "[panic] "+e.(string))
				}
				buf := new(bytes.Buffer)
				code := http.StatusInternalServerError
				buf, eerr := mw.presenter.ExecuteError(buf, code)
				if eerr != nil {
					mw.logger.ErrorContext(r.Context(), eerr.Error())
				}
				w.WriteHeader(code)
				if buf == nil {
					w.Write([]byte("Response Error"))
				}
				if _, err := buf.WriteTo(w); err != nil {
					w.Write([]byte(err.Error()))
				}
				return
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
// 				a.Logger.ErrorContext(r.Context(),err.Error())
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
// 				a.Logger.ErrorContext(r.Context(),"CSRF token is invalid")
// 				a.Presenter.Error(w, http.StatusInternalServerError)
// 				return
// 			}

// 			if err := a.redisHandler.HasCSRFToken(token); err != nil {
// 				a.Logger.ErrorContext(r.Context(),err.Error())
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
