package handlers

import (
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
	"net/http"
)

func GetAllPairs(w http.ResponseWriter, r *http.Request) {
	pairs, ok := storage.User.UserData[storage.User.ID]
	if !ok {
		log.Println("user pairs not found")
		resp.NoContent(w, http.StatusNoContent)
		return
	}

	resp.JSON(w, http.StatusOK, pairs)
}
