package main

import (
	"context"
	"embed"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"syscall"
	"time"

	"net/http/pprof"

	"github.com/bmf-san/bmf-tech-client/app/api"
	"github.com/bmf-san/bmf-tech-client/app/controller"
	"github.com/bmf-san/bmf-tech-client/app/logger"
	"github.com/bmf-san/bmf-tech-client/app/middleware"
	"github.com/bmf-san/bmf-tech-client/app/presenter"
	"github.com/bmf-san/goblin"
)

//go:embed templates/*
var templates embed.FS

const timeout time.Duration = 10 * time.Second

func main() {
	level, _ := strconv.Atoi(os.Getenv("LOG_LEVEl"))
	logger := logger.NewLogger(level)

	defer func() {
		if x := recover(); x != nil {
			logger.Error(string(debug.Stack()))
		}
		os.Exit(1)
	}()

	client := api.NewClient()
	presenter := presenter.NewPresenter(templates)
	mw := middleware.NewMiddleware(logger, presenter)

	hc := controller.NewHomeController(logger, client, presenter)
	pc := controller.NewPostController(logger, client, presenter)
	cc := controller.NewCategoryController(logger, client, presenter)
	tc := controller.NewTagController(logger, client, presenter)
	// TODO: implement later.
	// cmc := controller.NewCommentController(logger, client, presenter)
	pf := controller.NewProfileController(logger, presenter)
	sp := controller.NewSupportController(logger, presenter)
	sc := controller.NewSitemapController(logger, client, presenter)
	fc := controller.NewFeedController(logger, client, presenter)
	pp := controller.NewPrivacyPolicyController(logger, presenter)

	r := goblin.NewRouter()

	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/", http.HandlerFunc(pprof.Index))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/heap", pprof.Handler("heap"))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/mutex", pprof.Handler("mutex"))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/block", pprof.Handler("block"))

	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/favicon.ico`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	}))

	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/ads.txt`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	}))

	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/profile.png`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	}))

	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/css/style.css`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	}))

	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/js/toc.js`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	}))

	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/`, hc.Index())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/posts`, pc.Index())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/posts/search`, pc.IndexByKeyword())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/posts/:title`, pc.Show())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/posts/categories/:name`, pc.IndexByCategory())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/posts/tags/:name`, pc.IndexByTag())
	// TODO: implement later.
	// r.Methods("/posts/:title/comments").Use(mw.Recovery).Handler(`/posts/:title/comments`, cc.Store())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/categories`, cc.Index())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/tags`, tc.Index())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/profile`, pf.Index())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/support`, sp.Index())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/sitemap`, sc.Index())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/feed`, fc.Index())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/privacy_policy`, pp.Index())

	s := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: r,
	}

	logger.Info("Start server")

	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			logger.Error(err.Error())
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-q

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Gracefully shutdown")
}
