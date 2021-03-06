package handlers

import (
	"errors"
	"log"
	"net/http"

	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
)

// GetAllPairs handler provides as response all created pairs "Original URL":"Shorten URL" created for the user.
// User ID is obtained from the cookie.
func GetAllPairs(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] handlers -> GetAllPairs start: search pairs")
	pairs, err := storage.User.GetUserURLs()
	if err != nil {
		if errors.Is(err, storage.ErrMemoryNotFound) {
			log.Println("[ERROR] handlers -> GetAllPairs: user pairs not found")
			resp.NoContent(w, http.StatusNoContent)
			return
		}
		log.Println("[ERROR] handlers -> GetAllPairs err: ", err)
		resp.NoContent(w, http.StatusBadRequest)
		return
	}

	if len(pairs) == 0 || pairs == nil {
		log.Println("[DEBUG] handlers -> GetAllPairs: user pairs not found")
		resp.NoContent(w, http.StatusNoContent)
		return
	}

	log.Println("[INFO] handlers -> GetAllPairs end: user pairs found -> ", pairs)
	resp.JSON(w, http.StatusOK, pairs)
}
