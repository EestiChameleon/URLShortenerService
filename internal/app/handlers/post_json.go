package handlers

import (
	"encoding/json"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"io"
	"log"
	"net/http"
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
	// read body
	var reqBody ReqBody
	log.Println("JSONShortURL start: read r.Body")
	byteBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("JSONShortURL: unable to read body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	log.Println("JSONShortURL: json.Unmarshal(byteBody, &reqBody)")
	if err = json.Unmarshal(byteBody, &reqBody); err != nil {
		log.Println("JSONShortURL: unable to unmarshal body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	// check if it's not empty
	longURL := reqBody.URL
	if longURL == "" {
		log.Println("JSONShortURL: empty r.Body")
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	// get a short url to pair with the orig url
	shortURL, err := storage.User.Put(longURL)
	if err != nil {
		log.Println("JSONShortURL: storage.Pairs.Put(longURL) error:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	log.Println("JSONShortURL end: ")
	resp.JSON(w, http.StatusCreated, ResBody{shortURL})
}
