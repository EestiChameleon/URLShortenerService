package handlers

import (
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"io"
	"log"
	"net/http"
)

func PostProvideShortURL(w http.ResponseWriter, r *http.Request) {
	// read body
	log.Println("PostProvideShortURL: start - read r.Body")
	byteBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("PostProvideShortURL: unable to read body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	// check if it's not empty
	log.Println("PostProvideShortURL: check byteBody")
	origURL := string(byteBody)
	if origURL == "" {
		log.Println("PostProvideShortURL: empty incoming url")
		resp.WriteString(w, http.StatusBadRequest, "invalid url")
		return
	}

	// get a short url to pair with the orig url
	log.Println("PostProvideShortURL: storage.User.Put(origURL) - ", origURL)
	shortURL, err := storage.User.Put(origURL)
	if err != nil {
		log.Println("PostProvideShortURL: storage.Pairs.Put(longURL) error:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid url")
		return
	}
	log.Println("PostProvideShortURL: end")
	resp.WriteString(w, http.StatusCreated, shortURL)
}
