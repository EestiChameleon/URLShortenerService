package server

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/handlers"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

func Start() (err error) {
	r := router.New()
	r.GET("/{id}", handlers.GetOrigURL)

	r.POST("/", handlers.PostProvideShortURL)

	log.Fatal(fasthttp.ListenAndServe("localhost:8080", r.Handler))

	return
}
