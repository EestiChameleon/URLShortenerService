package storage

import (
	"crypto/rand"
	"fmt"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
)

type Store struct {
	db map[string]string
}

func NewStore() *Store {
	store := &Store{db: map[string]string{}}
	return store
}

var Pit = NewStore()

//test pit
func TestStore() *Store {
	store := &Store{db: map[string]string{"http://localhost:8080/test": "https://jwt.io/"}}
	return store
}

func (k Store) Get(key string) string {
	return k.db[key]
}

func (k Store) Put(value string) (key string, err error) {
	key, err = ShortURL()
	if err != nil {
		return "", err
	}
	_, ok := k.Check(key)
	if !ok {
		k.db[key] = value
		return key, nil
	} else {
		return k.Put(value)
	}
}

func (k Store) Check(key string) (value string, ok bool) {
	value, ok = k.db[key]
	if ok {
		return value, true
	} else {
		return "", false
	}
}

func ShortURL() (shortedURL string, err error) {
	// 7 bytes is enough to provide more than 78kkk diff combinations
	b := make([]byte, 7)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	shortedURL = fmt.Sprintf("%s/%x", cfg.Envs.BaseURL, b[0:])
	return
}
