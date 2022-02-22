package handlers

import (
	"encoding/json"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"io/ioutil"
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
	if !strings.Contains(r.Header.Get(resp.HeaderContentType), resp.MIMEApplicationJSON) {
		log.Println("invalid context-type: ", r.Header.Get(resp.HeaderContentType))
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	reqBody := ReqBody{}
	//read the url in the body
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil.ReadAll(r.Body) error:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}
	err = json.Unmarshal(bytes, &reqBody)
	if err != nil {
		log.Println("json.Unmarshal(bytes, &reqBody) error:", err)
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
	shortUrl, err := storage.Pairs.Put(longURL)
	if err != nil {
		log.Println("storage.Pairs.Put(longURL) error:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	resp.JSON(w, http.StatusCreated, ResBody{shortUrl})
	return
}
