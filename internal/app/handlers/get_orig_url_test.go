package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrigURL(t *testing.T) {
	type want struct {
		contentType    string
		statusCode     int
		respMessage    string
		headerLocation string
	}

	tests := []struct {
		name    string
		request string
		want    want
	}{
		{
			name:    "GET test #1: test url -> 307",
			request: "test",
			want: want{
				contentType:    "text/plain; charset=UTF-8",
				statusCode:     307,
				headerLocation: "https://jwt.io/",
			},
		},
		{
			name:    "GET test #2: empty id -> 400",
			request: "",
			want: want{
				statusCode:  400,
				respMessage: "invalid id",
			},
		},
		{
			name:    "GET test #3: wrong id -> 400",
			request: "666xxx",
			want: want{
				statusCode:  400,
				respMessage: "invalid id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Host = "localhost:8080"
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/id")
			c.SetParamNames("id")
			c.SetParamValues(tt.request)

			// Assertions
			err := GetOrigURL(c)
			if err == nil {
				res := rec.Result()
				assert.Equal(t, tt.want.statusCode, res.StatusCode)

				assert.Equal(t, tt.want.contentType, res.Header.Get(echo.HeaderContentType))
			} else {

				respMsg := fmt.Sprintf("code=%d, message=%s", tt.want.statusCode, tt.want.respMessage)
				assert.Equal(t, respMsg, err.Error())
			}

		})
	}
}
