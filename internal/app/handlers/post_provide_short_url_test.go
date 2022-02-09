package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func createPOSTRequest() *fasthttp.RequestCtx {
	req := &fasthttp.RequestCtx{
		Request:  fasthttp.Request{},
		Response: fasthttp.Response{},
	}
	return req
}

func TestPostProvideShortURL(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		response    string
	}

	tests := []struct {
		name       string
		requestURL string
		body       string
		want       want
	}{
		{
			name:       "POST test #1: url -> 201",
			requestURL: "http://localhost:8080/",
			body:       "https://jwt.io/",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  201,
				response:    "",
			},
		},
		{
			name:       "POST test #2: empty url -> 400",
			requestURL: "http://localhost:8080/",
			body:       "",
			want: want{
				contentType: "text/plain; charset=utf-8",
				response:    "empty body",
				statusCode:  400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := createPOSTRequest()
			ctx.Request.SetBodyString(tt.body)
			ctx.Request.SetHost(tt.requestURL)
			PostProvideShortURL(ctx)

			sc := ctx.Response.StatusCode()
			assert.Equal(t, tt.want.statusCode, sc)

			assert.Equal(t, tt.want.contentType, string(ctx.Response.Header.ContentType()))

			if tt.want.response != "" {
				assert.Equal(t, tt.want.response, string(ctx.Response.Body()))
			}
		})
	}
}
