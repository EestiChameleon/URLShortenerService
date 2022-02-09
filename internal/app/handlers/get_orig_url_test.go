package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func createGETRequest() *fasthttp.RequestCtx {
	req := &fasthttp.RequestCtx{
		Request:  fasthttp.Request{},
		Response: fasthttp.Response{},
	}
	return req
}

func TestGetOrigURL(t *testing.T) {
	type want struct {
		contentType    string
		statusCode     int
		response       string
		headerLocation string
	}

	tests := []struct {
		name       string
		requestURL string
		id         string
		want       want
	}{
		{
			name:       "GET test #1: test url -> 307",
			requestURL: "http://localhost:8080/",
			id:         "test",
			want: want{
				contentType:    "text/plain; charset=utf-8",
				statusCode:     307,
				headerLocation: "https://jwt.io/",
				response:       "",
			},
		},
		{
			name:       "GET test #2: empty id -> 400",
			requestURL: "http://localhost:8080/",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  400,
				response:    "invalid id",
			},
		},
		{
			name:       "GET test #3: wrong id -> 400",
			requestURL: "http://localhost:8080/",
			id:         "666xxx",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  400,
				response:    "no Args value for the given key",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := createGETRequest()
			ctx.Request.SetRequestURI(tt.requestURL)
			ctx.SetUserValue("id", tt.id)
			GetOrigURL(ctx)

			sc := ctx.Response.StatusCode()
			assert.Equal(t, tt.want.statusCode, sc)

			assert.Equal(t, tt.want.contentType, string(ctx.Response.Header.ContentType()))

			if tt.want.response != "" {
				assert.Equal(t, tt.want.response, string(ctx.Response.Body()))
			}
		})
	}
}
