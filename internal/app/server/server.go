package server

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/handlers"
	"net/http"
)

func Start() (err error) {
	http.HandleFunc("/", handlers.URLHandler)

	err = http.ListenAndServe("localhost:8080", nil)

	return
}
