package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/service/process"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
)

// ReqBody struct is used to unmarshal incoming request body.
type ReqBody struct {
	URL string `json:"url"`
}

// ResBody struct is used to marshal outgoing response body.
type ResBody struct {
	Result string `json:"result"`
}

// JSONShortURL handler receive a JSON {"url": "<original URL>"} and creates a short URL pair.
// In the response it returns JSON {"result": "<short url>"}.
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
