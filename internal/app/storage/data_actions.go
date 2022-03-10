package storage

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/robbert229/jwt"
	"log"
)

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

func CreateUserID() (string, error) {
	b := make([]byte, 7)
	_, err := rand.Read(b)
	if err != nil {
		log.Println(err)
		return ``, err
	}
	uuid := fmt.Sprintf("%x", b[0:])
	return uuid, nil
}

func FindKeyForEl(u string, m map[string]string) (string, bool) {
	for key, el := range m {
		if u == el {
			return key, true
		}
	}
	return ``, false
}

func JWTEncode(key, value string) (string, error) {
	algorithm := jwt.HmacSha256(cfg.Envs.CryptoKey)

	claims := jwt.NewClaim()
	claims.Set(key, value)

	token, err := algorithm.Encode(claims)
	if err != nil {
		log.Println("jwt encode Encode", err)
		return ``, err
	}

	if err = algorithm.Validate(token); err != nil {
		log.Println("jwt encode validate", err)
		return ``, err
	} else {
		return token, nil
	}
}

func JWTDecode(token, key string) (string, error) {
	algorithm := jwt.HmacSha256(cfg.Envs.CryptoKey)
	claims := jwt.NewClaim()

	if err := algorithm.Validate(token); err != nil {
		log.Println("jwt decode validate", err)
		return ``, err
	}

	claims, err := algorithm.Decode(token)
	if err != nil {
		log.Println("jwt decode Decode", err)
		return ``, err
	}

	data, err := claims.Get(key)
	if err != nil {
		log.Println("jwt decode Get", err)
		return ``, err
	}

	value, ok := data.(string)
	if !ok {
		log.Println("no data for given key")
		return ``, errors.New("failed to decode the provided Token")
	}

	return value, nil
}
