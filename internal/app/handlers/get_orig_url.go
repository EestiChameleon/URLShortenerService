package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetOrigURL(ctx echo.Context) (err error) {
	// get and check the passed ID
	id := ctx.Param("id")
	if id == "" {
		log.Println("empty shortURL id", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	// check for the short url in map
	shortedURL := fmt.Sprintf("http://%s/%s", ctx.Request().Host, id)
	longURL, ok := store.DB[shortedURL]
	if !ok {
		log.Println("shortURL pair not found", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)
	ctx.Response().Header().Set(echo.HeaderLocation, longURL)
	ctx.Response().WriteHeader(http.StatusTemporaryRedirect)
	return
}
