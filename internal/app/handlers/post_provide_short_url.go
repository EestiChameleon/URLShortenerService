package handlers

import (
	"crypto/rand"
	"fmt"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"github.com/valyala/fasthttp"
)

var store = &storage.Store{DB: map[string]string{"http://localhost:8080/test": "https://jwt.io/"}}

func PostProvideShortURL(ctx *fasthttp.RequestCtx) {

	reqURL := string(ctx.Request.Body())
	if reqURL == "" {
		ctx.Error("empty body", fasthttp.StatusBadRequest)
		return
	}

	shortedURL, err := shortUrl()
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	store.Put(shortedURL, reqURL)
	ctx.Response.Header.Set("Content-Type", "text/plain; charset=utf-8")
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
	ctx.Response.SetBodyString(shortedURL)

}

func shortUrl() (shortedURL string, err error) {
	b := make([]byte, 7)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	shortedURL = fmt.Sprintf("http://localhost:8080/%x", b[0:])
	return
}
