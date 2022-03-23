package handlers

import (
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
	"net/http"
)

func PingDatabase(w http.ResponseWriter, r *http.Request) {
	log.Println("PingDatabase: start")
	if err := storage.PingDB(); err != nil {
		log.Println("PingDatabase err: ", err)
		resp.NoContent(w, http.StatusInternalServerError)
		return
	}

	log.Println("PingDatabase: -> 200")
	resp.NoContent(w, http.StatusOK)
}
