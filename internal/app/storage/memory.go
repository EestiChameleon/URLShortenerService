package storage

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/EestiChameleon/URLShortenerService/internal/app/service/data"
)

var (
	ErrMemoryNotFound = errors.New("not found")
)

// MemoryStorage structure imitates the DB.
type MemoryStorage struct {
	ID       string
	Pairs    map[string]string
	UserData map[string][]Pair
}

// InitMemoryStorage function initializes the storage file where we save the URLs pairs data.
func InitMemoryStorage() (*MemoryStorage, error) {
	d := &MemoryStorage{
		ID:       "",
		Pairs:    map[string]string{},
		UserData: map[string][]Pair{},
	}

	log.Println("memory_storage InitMemoryStorage: start")
	if cfg.Envs.FileStoragePath != "" {
		// create/open file
		log.Println("memory_storage InitMemoryStorage: openfile")
		f, err := os.OpenFile(cfg.Envs.FileStoragePath, os.O_RDONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer f.Close()

		log.Println("memory_storage InitMemoryStorage: read file")
		bytes, err := os.ReadFile(cfg.Envs.FileStoragePath)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		log.Println("memory_storage InitMemoryStorage: check len(bytes)")
		if len(bytes) != 0 {
			if err = json.Unmarshal(bytes, &d.Pairs); err != nil {
				log.Println(err)
				return nil, err
			}
		}
	}

	log.Println("memory_storage InitMemoryStorage: OK")
	return d, nil
}

// GetOrigURL method gets from the storage file the original URL for the passed short URL.
func (m *MemoryStorage) GetOrigURL(shortURL string) (string, error) {
	log.Println("memory_storage - GetOrigURL: start")
	origURL, ok := m.Pairs[shortURL]
	if !ok || origURL == "" {
		return ``, ErrMemoryNotFound
	}
	return origURL, nil
}

// GetShortURL method gets from the storage file the short URL for the passed original URL.
func (m *MemoryStorage) GetShortURL(origURL string) (string, error) {
	log.Println("memory_storage - GetShortURL: start")
	for s, o := range m.Pairs {
		if o == origURL {
			return s, nil
		}
	}
	return ``, ErrMemoryNotFound
}

// SavePair method saves in the storage file new pair "original URL":"short URL" with user ID.
func (m *MemoryStorage) SavePair(pair Pair) (err error) {
	log.Println("memory_storage - SavePair: start")
	// save data
	log.Printf("memory_storage SavePair: save to Pairs. ShortURL: %s, OrigURL: %s\n", pair.ShortURL, pair.OrigURL)
	m.Pairs[pair.ShortURL] = pair.OrigURL
	log.Printf("memory_storage SavePair: append to UserData. ID: %s, ShortURL: %s, OrigURL: %s\n",
		m.ID, pair.ShortURL, pair.OrigURL)
	m.UserData[m.ID] = append(m.UserData[m.ID], pair)

	if cfg.Envs.FileStoragePath != "" {
		// update file
		log.Println("memory_storage SavePair: UpdateFile")
		if err = m.UpdateFile(); err != nil {
			log.Println(err)
			return
		}
	}

	log.Println("memory_storage SavePair: OK")
	return nil
}

// GetUserURLs method provides the user pairs list from the storage file.
func (m *MemoryStorage) GetUserURLs() ([]Pair, error) {
	log.Println("memory_storage GetUserURLs: start")
	pairs, ok := m.UserData[m.ID]
	if !ok {
		return nil, ErrMemoryNotFound
	}
	return pairs, nil
}

// Shutdown method closes the storage file with saving the latest data.
func (m *MemoryStorage) Shutdown() error {
	log.Println("memory_storage CloseMemoryStorage: start")
	if cfg.Envs.FileStoragePath != "" {
		if err := m.UpdateFile(); err != nil {
			log.Println(err)
			return err
		}
	}

	log.Println("memory_storage CloseMemoryStorage: OK")
	return nil
}

// SetUserID stores the user ID in the MemoryStorage structure.
func (m *MemoryStorage) SetUserID(userID string) {
	log.Println("memory_storage SetUserID: start")
	m.ID = userID
}

// GetUserID shows the stored user ID in the MemoryStorage structure.
func (m *MemoryStorage) GetUserID() string {
	return m.ID
}

// CreateShortURL creates a unique new short URL.
func (m *MemoryStorage) CreateShortURL() (shortURL string, err error) {
	log.Println("memory_storage GetShortURL: start")
	shortURL, err = data.ShortURL()
	if err != nil {
		log.Println(err)
		return ``, err
	}

	log.Println("memory_storage GetShortURL: check for already existing shortURL")
	_, ok := m.Pairs[shortURL]
	if ok {
		log.Println("memory_storage GetShortURL: shortURL already exists -> try again")
		return m.CreateShortURL()
	}

	log.Println("memory_storage GetShortURL: OK")
	return shortURL, nil
}

// UpdateFile method rewrite the storage file with the latest data.
func (m *MemoryStorage) UpdateFile() error {
	// open & rewrite file
	log.Println("memory_storage UpdateFile: start")
	f, err := os.OpenFile(cfg.Envs.FileStoragePath, os.O_WRONLY, 0777)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	// prepare data
	log.Println("memory_storage UpdateFile: json.Marshal(m.Pairs)")
	jsonByte, err := json.Marshal(m.UserData)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("memory_storage UpdateFile: f.Write(jsonByte)")
	_, err = f.Write(jsonByte)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("memory_storage UpdateFile: OK")
	return nil
}

// BatchDelete - mock for interface.
func (m *MemoryStorage) BatchDelete(shortURLs []string) error {
	return nil
}

func (m *MemoryStorage) GetStats() (int, int, error) {
	return len(m.Pairs), len(m.UserData), nil
}
