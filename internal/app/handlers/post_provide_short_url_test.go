package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostProvideShortURL(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		respMessage string
	}

	tests := []struct {
		name        string
		contentType string
		request     string
		body        string
		want        want
	}{
		{
			name:        "POST test #1: url -> 201",
			contentType: echo.MIMETextPlainCharsetUTF8,
			request:     "http://localhost:8080/",
			body:        "https://jwt.io/",
			want: want{
				contentType: echo.MIMETextPlainCharsetUTF8,
				statusCode:  201,
				respMessage: "",
			},
		},
		{
			name:        "POST test #2: empty url -> 400",
			contentType: echo.MIMETextPlainCharsetUTF8,
			request:     "http://localhost:8080/",
			body:        "",
			want: want{
				contentType: "text/plain; charset=utf-8",
				respMessage: "invalid url",
				statusCode:  400,
			},
		},
		{
			name:        "POST test #3: wrong content type -> 400",
			contentType: echo.MIMEApplicationJSON,
			request:     "http://localhost:8080/",
			body:        "https://jwt.io/",
			want: want{
				contentType: echo.MIMETextPlainCharsetUTF8,
				statusCode:  400,
				respMessage: "invalid url",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			req.Header.Set(echo.HeaderContentType, tt.contentType)
			req.Host = "http://localhost:8080/"
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			// Assertions
			err := PostProvideShortURL(ctx)
			res := rec.Result()
			defer res.Body.Close()
			if err == nil {
				assert.Equal(t, tt.want.statusCode, res.StatusCode)
				assert.Equal(t, tt.want.contentType, res.Header.Get(echo.HeaderContentType))
				url, err := io.ReadAll(res.Body)
				if err != nil {
					t.Log(err)
				}
				assert.NotEqual(t, tt.want.respMessage, url)
			} else {

				respMsg := fmt.Sprintf("code=%d, message=%s", tt.want.statusCode, tt.want.respMessage)
				assert.Equal(t, respMsg, err.Error())
			}

		})
	}
}
