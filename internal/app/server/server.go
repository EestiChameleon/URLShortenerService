package server

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
)

func Start() {
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
	router.Get("/{id}", handlers.GetOrigURL)

	router.Post("/", handlers.PostProvideShortURL)

	// Start server
	s := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
}
