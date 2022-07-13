package process

import (
	"context"
	"log"

	"golang.org/x/sync/errgroup"

	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
)

func BatchDelete(shortURL []string) {
	userURLs := filterByUserID(shortURL)
	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return storage.User.BatchDelete(userURLs)
	})
	if err := g.Wait(); err != nil {
		log.Println("[FATAL] unable to delete Users shortURLs:", err)
	}
}

func filterByUserID(shortURL []string) []string {
	var filtredList []string

	pairs, err := storage.User.GetUserURLs()
	if err != nil {
		log.Println("[FATAL] unable to get Users pairs list:", err)
	}

	for _, sURL := range shortURL {
		for _, pair := range pairs {
			if sURL == pair.ShortURL {
				filtredList = append(filtredList, sURL)
			}
		}
	}

	return filtredList
}
