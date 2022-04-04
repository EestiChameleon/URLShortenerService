package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
)

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

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return storage.User.BatchDelete(shortURLs)
	})
	if err = g.Wait(); err != nil {
		log.Println("[ERROR] handlers -> DeleteBatch: db query failed:", err)
	}

	log.Println("[INFO] handlers -> DeleteBatch: OK")
	resp.NoContent(w, http.StatusAccepted)
}
