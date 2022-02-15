package handlers

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
)

func PostProvideShortURL(ctx echo.Context) (err error) {
	//// check the content type - we are expecting an incoming text url
	//if ctx.Request().Header.Get("Content-Type") != echo.MIMETextPlainCharsetUTF8 {
	//	log.Println("invalid context-type")
	//	return echo.NewHTTPError(http.StatusBadRequest, "invalid content-type")
	//}
	// read the url in the body
	bytes, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Println("ioutil.ReadAll(ctx.Request.Body", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid url")
	}
	// check if it's not empty
	longURL := string(bytes)
	if longURL == "" {
		log.Println("empty incoming url")
		return echo.NewHTTPError(http.StatusBadRequest, "invalid url")
	}

	// get a short url to pair with the orig url
	shortUrl, err := storage.Pit.Put(longURL)

	return ctx.String(http.StatusCreated, shortUrl)
}
