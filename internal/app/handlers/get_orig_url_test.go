package handlers

import (
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
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
			request: "/test",
			want: want{
				contentType:    resp.MIMETextPlainCharsetUTF8,
				statusCode:     307,
				headerLocation: "https://jwt.io/",
			},
		},
		{
			name:    "GET test #2: empty id -> 404",
			request: "/",
			want: want{
				contentType: resp.MIMETextPlainCharsetUTF8,
				statusCode:  404,
				respMessage: "404 page not found\n",
			},
		},
		{
			name:    "GET test #3: wrong id -> 400",
			request: "/666xxx",
			want: want{
				contentType: resp.MIMETextPlainCharsetUTF8,
				statusCode:  400,
				respMessage: "invalid id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.request, nil)
			if tt.request == "/test" {
				storage.Pit = storage.TestStore()
			}
			// создаём новый Recorder
			w := httptest.NewRecorder()
			r := chi.NewRouter()
			// определяем хендлер
			r.Get("/{id}", GetOrigURL)
			// запускаем сервер
			r.ServeHTTP(w, request)
			res := w.Result()

			assert.Equal(t, tt.want.statusCode, res.StatusCode)
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))

			if tt.want.respMessage != "" {
				defer res.Body.Close()
				urlResult, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)

				assert.Equal(t, tt.want.respMessage, string(urlResult))
			}

		})
	}
}
