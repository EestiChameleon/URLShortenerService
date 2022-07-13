package handlers

import (
	"log"
	"net/http"

	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
)

// PingDatabase handler verifies the DB connection.
// It executes an empty sql statement against DB pool.
// If the sql returns without error, the PingDatabase is considered successful, otherwise, the error is returned.
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
