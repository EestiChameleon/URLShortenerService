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
	log.Println(pairs)
	if err != nil || len(pairs) == 0 {
		log.Println("GetAllPairs: user pairs not found - ", err)
		resp.NoContent(w, http.StatusNoContent)
		return
	}

	log.Println("GetAllPairs end: user pairs found -> JSON 200")
	resp.JSON(w, http.StatusOK, pairs)
}
