package handlers

import (
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/httpserver/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
	"net/http"
)

type ServiceStat struct {
	AllShortedURLs int `json:"urls"`  // all shorted urls in the service
	UsersQuantity  int `json:"users"` // users quantity in the service
}

// GetStat return all shorted urls and users quantity
func GetStat(w http.ResponseWriter, r *http.Request) {
	urls, users, err := storage.STRG.GetStats()
	if err != nil {
		log.Println("get stats err -", err)
		resp.NoContent(w, http.StatusBadRequest)
		return
	}

	resp.JSON(w, http.StatusOK, ServiceStat{
		AllShortedURLs: urls,
		UsersQuantity:  users,
	})
}
