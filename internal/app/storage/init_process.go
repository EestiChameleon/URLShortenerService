package storage

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"log"
)

var (
	User Data
)

type Data interface {
	GetURL(shortURL string) (string, error) // find OrigURL by ShortURL
	SavePair(pair Pair) error               // save origURL and provide shortURL
	GetUserURLs() ([]Pair, error)           // provide a list of all shorten links by userID
	ShutDown() error                        // close the storage
	SetUserID(userID string)
	CreateShortURL() (string, error)
}

type Pair struct {
	ShortURL string `json:"short_url"`
	OrigURL  string `json:"original_url"`
}

// InitStorage method provides a Memory/File/DB storage, based on config data
func InitStorage() (err error) {
	log.Println("init_storage InitStorage: start")
	// not the default db for checks
	if cfg.Envs.DatabaseDSN != "" { //} && cfg.Envs.DatabaseDSN != "postgresql://localhost:5432/yandex_practicum_db" {
		log.Println("init_storage InitStorage: DB case")
		User, err = InitDBStorage()
		if err != nil {
			log.Println("init_process InitStorage: InitDBStorage err - ", err)
			return err
		}
	} else {
		log.Println("init_storage InitStorage: Memory case")
		User, err = InitMemoryStorage()
		if err != nil {
			log.Println("init_process InitStorage: InitMemoryStorage err - ", err)
			return err
		}
	}

	log.Println("init_process InitStorage: end")
	return nil
}
