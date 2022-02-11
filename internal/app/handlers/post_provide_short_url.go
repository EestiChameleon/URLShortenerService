package handlers

import (
	"crypto/rand"
	"fmt"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
)

var store = &storage.Store{DB: map[string]string{"http://localhost:8080/test": "https://jwt.io/"}}

func PostProvideShortURL(ctx echo.Context) (err error) {
	// check the content type - we are expecting an incoming text url
	if ctx.Request().Header.Get("Content-Type") != echo.MIMETextPlainCharsetUTF8 {
		log.Println("invalid context-type")
		return echo.NewHTTPError(http.StatusBadRequest, "invalid content-type")
	}
	// read the url in the body
	defer ctx.Request().Body.Close()
	bytes, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Println("ioutil.ReadAll(ctx.Request.Body", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid url")
	}
	// check if it's not empty
	longURL := string(bytes)
	if longURL == "" {
		log.Println("empty incoming url")
		return echo.NewHTTPError(http.StatusBadRequest, "empty url")
	}
	// get a short url to pair with the orig url
	shortURL, err := shortUrl()
	if err != nil {
		log.Println("save shortURL: longURL pair failed", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	// save the pair
	store.Put(shortURL, longURL)

	return ctx.String(http.StatusCreated, shortURL)
}

func shortUrl() (shortedURL string, err error) {
	// 7 bytes is enough to provide more than 78kkk diff combinations
	b := make([]byte, 7)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	shortedURL = fmt.Sprintf("http://localhost:8080/%x", b[0:])
	return
}
