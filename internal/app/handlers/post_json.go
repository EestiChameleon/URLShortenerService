package handlers

import (
	"encoding/json"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
	"net/http"
	"strings"
)

type ReqBody struct {
	URL string `json:"url"`
}

type ResBody struct {
	Result string `json:"result"`
}

// JSONShortURL принимает в теле запроса JSON-объект {"url": "<some_url>"}
// возвращает в ответ объект {"result": "<shorten_url>"}.
func JSONShortURL(w http.ResponseWriter, r *http.Request) {
	// check the content type - we are expecting an incoming JSON
	rContentType := r.Header.Get(resp.HeaderContentType)
	if !strings.Contains(rContentType, resp.MIMEApplicationJSON) {
		log.Println("invalid context-type: ", r.Header.Get(resp.HeaderContentType))
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	// read body
	var reqBody ReqBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Println("unable to parse body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	// check if it's not empty
	longURL := reqBody.URL
	if longURL == "" {
		log.Println("empty incoming url")
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	// get a short url to pair with the orig url
	shortURL, err := storage.Pairs.Put(longURL)
	if err != nil {
		log.Println("storage.Pairs.Put(longURL) error:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	resp.JSON(w, http.StatusCreated, ResBody{shortURL})
}
