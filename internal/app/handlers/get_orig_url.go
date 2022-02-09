package handlers

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

func GetOrigURL(ctx *fasthttp.RequestCtx) {
	// GET /{id}
	//var empty interface{}
	id := ctx.UserValue("id").(string)
	//
	//uri := strings.Split(ctx.Request.URI().String(), "/")
	//id := uri[3]

	if id == "" {
		ctx.Error("invalid id", fasthttp.StatusBadRequest)
		return
	}

	shortedURL := fmt.Sprintf("http://%s/%s", string(ctx.Request.Host()), id) //.(string))

	longURL, ok := store.DB[shortedURL]
	if !ok {
		ctx.Error(fasthttp.ErrNoArgValue.Error(), fasthttp.StatusBadRequest)
		return
	}

	ctx.Response.Header.Set("Location", longURL)
	ctx.Response.SetStatusCode(http.StatusTemporaryRedirect)
}
