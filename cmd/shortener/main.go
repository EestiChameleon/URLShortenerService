package main

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", app.URLHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
