package handlers

import (
	"encoding/json"
	"fmt"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/httpserver/responses"
	"io"
	"log"
	"net/http"

	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/EestiChameleon/URLShortenerService/internal/app/service/process"
)

// DeleteBatch handler allows to delete multiple shorted URLs at single API call.
func DeleteBatch(w http.ResponseWriter, r *http.Request) {
	var reqBody []string

	// read body
	log.Println("[INFO] handlers -> DeleteBatch: start - read r.Body")
	byteBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR] handlers -> DeleteBatch: unable to read body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	log.Println("[DEBUG] handlers -> DeleteBatch: json.Unmarshal(byteBody, &reqBody)")
	if err = json.Unmarshal(byteBody, &reqBody); err != nil {
		log.Println("[ERROR] handlers -> DeleteBatch: unable to unmarshal body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	//check for empty body
	if len(reqBody) < 1 {
		log.Println("[DEBUG] handlers -> DeleteBatch: empty body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	var shortURLs []string
	for _, id := range reqBody {
		shortURL := fmt.Sprintf("%s/%s", cfg.Envs.BaseURL, id)
		shortURLs = append(shortURLs, shortURL)
	}

	// BatchDelete call from the process package. Minimum DB interaction in the handler.
	log.Println("[INFO] handlers -> DeleteBatch: shortURLs list sent to process.BatchDelete")
	process.BatchDelete(shortURLs)

	log.Println("[INFO] handlers -> DeleteBatch: OK")
	resp.NoContent(w, http.StatusAccepted)
}
