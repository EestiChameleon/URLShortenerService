package handlers

import (
	"errors"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
	"net/http"
)

func GetAllPairs(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] handlers -> GetAllPairs start: search pairs")
	pairs, err := storage.User.GetUserURLs()

	if err != nil {
		if errors.Is(err, storage.ErrMemoryNotFound) {
			log.Println("[INFO] handlers -> GetAllPairs: user pairs not found")
			resp.NoContent(w, http.StatusNoContent)
			return
		}
		log.Println("[ERROR] handlers -> GetAllPairs err: ", err)
		resp.NoContent(w, http.StatusBadRequest)
		return
	}

	if len(pairs) == 0 {
		log.Println("[INFO] handlers -> GetAllPairs: user pairs not found")
		resp.NoContent(w, http.StatusNoContent)
		return
	}

	log.Println("[INFO] handlers -> GetAllPairs end: user pairs found -> ", pairs)
	resp.JSON(w, http.StatusOK, pairs)
}
