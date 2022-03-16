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

	var shortURL string
	// check for already existing short URL
	shortURL, err = storage.User.GetShortURL(origURL)
	if err != nil && err != storage.ErrMemoryNotFound {
		log.Println("PostProvideShortURL: GetShortURL err:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}
	// return the already existing shortURL
	if shortURL != "" {
		log.Println("PostProvideShortURL: ShortURL already exists - ", shortURL)
		resp.WriteString(w, http.StatusConflict, shortURL)
		return
	}

	// get a NEW short url to pair with the orig url
	shortURL, err = storage.User.CreateShortURL()
	if err != nil {
		log.Println("PostProvideShortURL: CreateShortURL err:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	if err = storage.User.SavePair(storage.Pair{ShortURL: shortURL, OrigURL: origURL}); err != nil {
		log.Println("PostProvideShortURL: storage.User.SavePair err:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}
	log.Println("PostProvideShortURL: end")
	resp.WriteString(w, http.StatusCreated, shortURL)
}
