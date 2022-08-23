package httpserver

import (
	"context"
	"crypto/tls"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	custommw "github.com/EestiChameleon/URLShortenerService/internal/app/server/httpserver/custommw"
	handlers "github.com/EestiChameleon/URLShortenerService/internal/app/server/httpserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
)

type HTTPServer struct {
	*http.Server
}

func InitHttpServer() (*HTTPServer, error) {
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

		serv := &http.Server{
			Addr:    cfg.Envs.SrvAddr, //":443", ???
			Handler: router,
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}
		go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
		//return s.ListenAndServeTLS("", "")
		return &HTTPServer{serv}, nil

	} else {
		// HTTP:
		serv := &http.Server{
			Addr:    cfg.Envs.SrvAddr,
			Handler: router,
			// ReadTimeout: 30 * time.Second, // customize http.Server timeouts
		}
		//return s.ListenAndServe()
		return &HTTPServer{serv}, nil
	}
}

func (h *HTTPServer) Start() error {
	if cfg.Envs.EnableHTTPS {
		// HTTPS
		return h.ListenAndServeTLS("", "")
	} else {
		// HTTP
		return h.ListenAndServe()
	}
}

func (h *HTTPServer) ShutDown() error {
	return h.Shutdown(context.Background())
}
