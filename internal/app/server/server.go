package server

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	cmw "github.com/EestiChameleon/URLShortenerService/internal/app/custommw"
	"github.com/EestiChameleon/URLShortenerService/internal/app/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
)

func Start() error {
	// Chi instance
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	//router.Use(middleware.Timeout(60 * time.Second))

	// Routes
	router.With(cmw.RequestGZIP, cmw.ResponseGZIP).Get("/{id}", handlers.GetOrigURL)

	router.With(cmw.RequestGZIP, cmw.ResponseGZIP).Post("/", handlers.PostProvideShortURL)
	router.With(cmw.RequestGZIP, cmw.ResponseGZIP).Post("/api/shorten", handlers.JSONShortURL)

	// Start server
	s := http.Server{
		Addr:    cfg.Envs.SrvAddr,
		Handler: router,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}

	return s.ListenAndServe()
}
