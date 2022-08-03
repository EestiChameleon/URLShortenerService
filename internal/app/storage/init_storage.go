package storage

import (
	"log"

	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
)

var (
	User Data
)

// Data interface used by DB and Memory structures for URL interactions.
type Data interface {
	GetOrigURL(shortURL string) (string, error) // find OrigURL by ShortURL
	GetShortURL(origURL string) (string, error) // find ShortURL by OrigURL
	SavePair(pair Pair) error                   // save origURL and shortURL pair to storage
	GetUserURLs() ([]Pair, error)               // provide a list of all shorten links by userID
	BatchDelete(shortURLs []string) error       // db function - set shorten_pairs.deleted = true for the given shortURLs lsit

	SetUserID(userID string)         // save userID to struct
	GetUserID() string               // provide userID
	CreateShortURL() (string, error) // creates a new shortURL (checks for being unique in the storage)

	Shutdown() error // close the storage
}

// Pair structure is used to create pairs "original URL":"short URL".
type Pair struct {
	ShortURL string `json:"short_url"`
	OrigURL  string `json:"original_url"`
}

// InitStorage method provides a Memory/File/DB storage, based on config data.
func InitStorage() (err error) {
	log.Println("[INFO] storage -> InitStorage: start")
	// not the default db for checks
	if cfg.Envs.DatabaseDSN != "" {
		log.Println("[DEBUG] storage -> InitStorage: DB case")
		User, err = InitDBStorage()
		if err != nil {
			log.Println("[ERROR] storage -> InitStorage: InitDBStorage err -", err)
			return err
		}
	} else {
		log.Println("[DEBUG] storage -> InitStorage: Memory case")
		User, err = InitMemoryStorage()
		if err != nil {
			log.Println("[ERROR] storage -> InitStorage: InitMemoryStorage err -", err)
			return err
		}
	}

	log.Println("[DEBUG] storage -> InitStorage: OK")
	return nil
}
