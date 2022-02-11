package server

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Start() {
	// Echo instance
	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// Routes
	router.GET("/:id", handlers.GetOrigURL)

	router.POST("/", handlers.PostProvideShortURL)

	// Start server
	s := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		router.Logger.Fatal(err)
	}
}
