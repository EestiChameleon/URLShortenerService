package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
)

func TestGetAllPairs(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		respMessage string
	}

	tests := []struct {
		name    string
		request string
		userID  string
		want    want
	}{
		{
			name:    "GET test #1: no user id - no content -> 204",
			request: "http://localhost:8080/api/user/urls",
			userID:  "",
			want: want{
				statusCode: 204,
			},
		},
		{
			name:    "GET test #2: test user id -> 200",
			request: "http://localhost:8080/api/user/urls",
			userID:  "testUser",
			want: want{
				contentType: resp.MIMEApplicationJSONCharsetUTF8,
				statusCode:  200,
				respMessage: "[{\"short_url\":\"http://localhost:8080/test\",\"original_url\":\"https://jwt.io/\"}]",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.request, nil)
			// envs
			cfg.Envs.BaseURL = "http://localhost:8080"
			//cfg.Envs.FileStoragePath = "tmp/testFile"
			//cfg.Envs.DatabaseDSN = "postgresql://localhost:5432/yandex_practicum_db"
			if err := storage.InitStorage(); err != nil {
				log.Fatal(err)
			}
			if tt.userID != "" {
				storage.User.SetUserID("testUser")
				if err := storage.User.SavePair(storage.Pair{
					ShortURL: "http://localhost:8080/test",
					OrigURL:  "https://jwt.io/",
				}); err != nil {
					log.Fatal(err)
				}
			}
			defer os.Remove(cfg.Envs.FileStoragePath)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			r := chi.NewRouter()
			// определяем хендлер
			storage.User.SetUserID(tt.userID)
			r.Get("/api/user/urls", GetAllPairs)
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
