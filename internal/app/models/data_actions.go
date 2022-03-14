package models

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/robbert229/jwt"
	"log"
)

func ShortURL() (shortURL string, err error) {
	log.Println("data actions: ShortURL start")
	// 7 bytes is enough to provide more than 78kkk diff combinations
	b := make([]byte, 7)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	shortURL = fmt.Sprintf("%s/%x", cfg.Envs.BaseURL, b[0:])
	log.Printf("data actions: ShortURL end. shortURL: %v\n", shortURL)
	return
}

func CreateUserID() (string, error) {
	log.Println("data actions: CreateUserID start")
	b := make([]byte, 7)
	_, err := rand.Read(b)
	if err != nil {
		log.Println(err)
		return ``, err
	}
	uuid := fmt.Sprintf("%x", b[0:])
	log.Println("data actions: CreateUserID end. UserID: ", uuid)
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
	log.Printf("data actions: JWTEncode start. Key: %v, Value: %v\n", key, value)
	algorithm := jwt.HmacSha256(cfg.Envs.CryptoKey)

	claims := jwt.NewClaim()
	claims.Set(key, value)

	token, err := algorithm.Encode(claims)
	if err != nil {
		log.Println("JWTEncode Encode err: ", err)
		return ``, err
	}

	if err = algorithm.Validate(token); err != nil {
		log.Println("JWTEncode validate err: ", err)
		return ``, err
	} else {
		log.Printf("data actions: JWTEncode end. Token: %v\n", token)
		return token, nil
	}
}

func JWTDecode(token, key string) (string, error) {
	log.Printf("data actions: JWTDecode start. Token: %v, Key: %v\n", token, key)
	algorithm := jwt.HmacSha256(cfg.Envs.CryptoKey)

	if err := algorithm.Validate(token); err != nil {
		log.Println("JWTDecode validate err: ", err)
		return ``, err
	}

	claims, err := algorithm.Decode(token)
	if err != nil {
		log.Println("JWTDecode Decode err: ", err)
		return ``, err
	}

	data, err := claims.Get(key)
	if err != nil {
		log.Println("JWTDecode Get err: ", err)
		return ``, err
	}

	value, ok := data.(string)
	if !ok {
		log.Printf("JWTEncode - no data for given key: %v\n", key)
		return ``, errors.New("failed to decode the provided Token")
	}
	log.Printf("data actions: JWTDecode end. Value: %v\n", value)
	return value, nil
}
