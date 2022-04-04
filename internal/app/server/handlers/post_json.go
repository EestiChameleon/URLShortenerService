package handlers

import (
	"encoding/json"
	"errors"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/service/process"
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
	log.Println("[INFO] handlers -> JSONShortURL: start - read r.Body")
	byteBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR] handlers -> JSONShortURL: unable to read body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	log.Println("[DEBUG] handlers -> JSONShortURL: json.Unmarshal(byteBody, &reqBody)")
	if err = json.Unmarshal(byteBody, &reqBody); err != nil {
		log.Println("[ERROR] handlers -> JSONShortURL: unable to unmarshal body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	// check if it's not empty
	origURL := reqBody.URL
	if origURL == "" {
		log.Println("[DEBUG] handlers -> JSONShortURL: empty r.Body")
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	shortURL, err := process.ShortURLforOrigURL(origURL)
	if err != nil {
		if errors.Is(err, storage.ErrDBOrigURLExists) {
			log.Println("[DEBUG] handlers -> JSONShortURL: shortURL & origURL pair exists")
			resp.JSON(w, http.StatusConflict, ResBody{shortURL})
			return
		}
		log.Println("[ERROR] handlers -> JSONShortURL: ShortURLforOrigURL err:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	log.Println("[INFO] handlers -> JSONShortURL: OK")
	resp.JSON(w, http.StatusCreated, ResBody{shortURL})
}
