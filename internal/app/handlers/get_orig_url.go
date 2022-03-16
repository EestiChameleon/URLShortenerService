package handlers

import (
	"fmt"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func GetOrigURL(w http.ResponseWriter, r *http.Request) {
	// get and check the passed ID
	log.Println("GetOrigURL start: search id")
	id := chi.URLParam(r, "id")
	// if id is empty - chi router will provide 404 error as "unknown path GET /"

	// check for the short url in map
	shortURL := fmt.Sprintf("%s/%s", cfg.Envs.BaseURL, id)
	log.Println("GetOrigURL: search for shortURL pair - ", shortURL)
	origURL, err := storage.User.GetOrigURL(shortURL)
	if err != nil || origURL == "" {
		log.Println("GetOrigURL: orig URL not found -> 400")
		resp.WriteString(w, http.StatusBadRequest, "invalid id")
		return
	}

	log.Println("GetOrigURL end: orig URL found -> 307")
	resp.RedirectString(w, origURL)
}
