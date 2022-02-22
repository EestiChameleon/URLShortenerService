package handlers

import (
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"io/ioutil"
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

	//read the url in the body
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil.ReadAll(r.Body) error:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid url")
		return
	}

	// check if it's not empty
	longURL := string(bytes)
	if longURL == "" {
		log.Println("empty incoming url")
		resp.WriteString(w, http.StatusBadRequest, "invalid url")
		return
	}

	// get a short url to pair with the orig url
	shortUrl, err := storage.Pairs.Put(longURL)

	resp.WriteString(w, http.StatusCreated, shortUrl)
	return
}
