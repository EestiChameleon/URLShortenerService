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
	id := chi.URLParam(r, "id")
	// if id is empty - chi router will provide 404 error as "unknown path GET /"

	// check for the short url in map
	shortedURL := fmt.Sprintf("%s/%s", cfg.Envs.BaseURL, id)

	longURL, ok := storage.Pairs.Check(shortedURL)
	if !ok {
		log.Println("shortURL pair not found")
		resp.WriteString(w, http.StatusBadRequest, "invalid id")
		return
	}

	resp.RedirectString(w, longURL)
	return
}
