package handlers

import (
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
	"net/http"
)

func GetAllPairs(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllPairs start: search pairs")
	pairs, err := storage.User.GetUserURLs()
	if err != nil {
		log.Println("GetAllPairs: user pairs not found -> 204")
		resp.NoContent(w, http.StatusNoContent)
		return
	}
	log.Println("GetAllPairs end: user pairs found -> JSON 200")
	resp.JSON(w, http.StatusOK, pairs)
}
