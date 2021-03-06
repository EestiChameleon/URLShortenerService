package handlers

import (
	"errors"
	"io"
	"log"
	"net/http"

	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/service/process"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
)

// PostProvideShortURL handler receive a text "original URL" and creates a short URL pair.
// In the response it returns text "<short url>".
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

	shortURL, err := process.ShortURLforOrigURL(origURL)
	if err != nil {
		if errors.Is(err, storage.ErrDBOrigURLExists) {
			resp.WriteString(w, http.StatusConflict, shortURL)
			return
		}
		resp.WriteString(w, http.StatusBadRequest, "invalid url")
		return
	}

	log.Println("PostProvideShortURL: end")
	resp.WriteString(w, http.StatusCreated, shortURL)
}
