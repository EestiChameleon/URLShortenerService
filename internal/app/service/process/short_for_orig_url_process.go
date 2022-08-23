package process

import (
	"errors"
	"log"

	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
)

// ShortURLforOrigURL function creates new pair "original url":"short url" and save it.
// It returns the new created short URL.
func ShortURLforOrigURL(origURL string) (string, error) {
	// get a NEW short url to pair with the orig url
	shortURL, err := storage.STRG.CreateShortURL()
	if err != nil {
		log.Println("[ERROR] service -> ShortURLforOrigURL: CreateShortURL err:", err)
		return ``, err
	}

	if err = storage.STRG.SavePair(storage.Pair{ShortURL: shortURL, OrigURL: origURL}); err != nil {
		if errors.Is(err, storage.ErrDBOrigURLExists) {
			shortURL, err = storage.STRG.GetShortURL(origURL)
			if err != nil && err != storage.ErrNotFound { // errors.Is не работает
				log.Println("[ERROR] service -> ShortURLforOrigURL: GetShortURL err:", err)
				return ``, err
			}
			log.Println("[INFO] service -> ShortURLforOrigURL: ShortURL already exists - ", shortURL)
			return shortURL, storage.ErrDBOrigURLExists
		} else {
			log.Println("[ERROR] service -> ShortURLforOrigURL: storage.User.SavePair err:", err)
			return ``, err
		}
	}

	return shortURL, nil
}
