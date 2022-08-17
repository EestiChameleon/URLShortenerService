package server

import (
	"context"
	"crypto/tls"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/crypto/acme/autocert"

	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/EestiChameleon/URLShortenerService/internal/app/server/custommw"
	"github.com/EestiChameleon/URLShortenerService/internal/app/server/handlers"

	"net/http"
)

var s *http.Server

// Start starts the server router.
func Start() error {
	// Chi instance
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// custom middleware
	router.Use(custommw.CheckCookie)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	// router.Use(middleware.Timeout(60 * time.Second))

	// Routes
	router.With(custommw.ResponseGZIP).Get("/{id}", handlers.GetOrigURL)
	router.With(custommw.ResponseGZIP).Get("/api/user/urls", handlers.GetAllPairs)
	router.Get("/ping", handlers.PingDatabase)

	router.With(custommw.RequestGZIP, custommw.ResponseGZIP).Post("/", handlers.PostProvideShortURL)
	router.With(custommw.RequestGZIP, custommw.ResponseGZIP).Post("/api/shorten", handlers.JSONShortURL)
	router.With(custommw.RequestGZIP, custommw.ResponseGZIP).Post("/api/shorten/batch", handlers.PostBatch)

	router.With(custommw.RequestGZIP, custommw.ResponseGZIP).Delete("/api/user/urls", handlers.DeleteBatch)

	router.With(custommw.TrustSubnetCheck).Get("/api/internal/stats", handlers.GetStat)

	// Start server

	// HTTPS:
	if cfg.Envs.EnableHTTPS {
		certManager := autocert.Manager{
			Prompt: autocert.AcceptTOS,
			Cache:  autocert.DirCache("certs"),
		}

		s = &http.Server{
			Addr:    cfg.Envs.SrvAddr, //":443", ???
			Handler: router,
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}
		go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
		return s.ListenAndServeTLS("", "")

	} else {
		// HTTP:
		s = &http.Server{
			Addr:    cfg.Envs.SrvAddr,
			Handler: router,
			// ReadTimeout: 30 * time.Second, // customize http.Server timeouts
		}
		return s.ListenAndServe()
	}
}

func Shutdown() error {
	return s.Shutdown(context.Background())
}
