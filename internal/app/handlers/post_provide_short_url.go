package handlers

import (
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
	"net/http"
	"strings"
)

func PostProvideShortURL(w http.ResponseWriter, r *http.Request) {
	// check the content type - we are expecting an incoming text url
	if !strings.Contains(r.Header.Get("Content-Type"), resp.MIMETextPlain) {
		log.Println("invalid context-type: ", r.Header.Get("Content-Type"))
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	// read body
	byteBody, ok := r.Context().Value("bodyURL").([]byte)
	if !ok {
		log.Println("unable to decode body: bodyURL missing in the context")
		resp.WriteString(w, http.StatusBadRequest, "invalid url")
		return
	}

	// check if it's not empty
	longURL := string(byteBody)
	if longURL == "" {
		log.Println("empty incoming url")
		resp.WriteString(w, http.StatusBadRequest, "invalid url")
		return
	}

	// get a short url to pair with the orig url
	shortURL, err := storage.User.Put(longURL)
	if err != nil {
		log.Println("storage.Pairs.Put(longURL) error:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid url")
		return
	}

	resp.WriteString(w, http.StatusCreated, shortURL)
}
