package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
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
			contentType: resp.MIMETextPlainCharsetUTF8,
			request:     "http://localhost:8080/",
			body:        "https://jwt.io/",
			want: want{
				contentType: resp.MIMETextPlainCharsetUTF8,
				statusCode:  201,
				respMessage: "",
			},
		},
		{
			name:        "POST test #2: empty url -> 400",
			contentType: resp.MIMETextPlainCharsetUTF8,
			request:     "http://localhost:8080/",
			body:        "",
			want: want{
				contentType: resp.MIMETextPlainCharsetUTF8,
				respMessage: "invalid url",
				statusCode:  400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.request, strings.NewReader(tt.body))
			request.Header.Set(resp.HeaderContentType, tt.contentType)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			// определяем хендлер
			h := http.HandlerFunc(PostProvideShortURL)
			// запускаем сервер
			// envs
			cfg.Envs.BaseURL = "http://localhost:8080"
			//cfg.Envs.FileStoragePath = "tmp/testFile"
			//cfg.Envs.DatabaseDSN = "postgresql://localhost:5432/yandex_practicum_db"
			if err := storage.InitStorage(); err != nil {
				log.Fatal(err)
			}
			defer os.Remove(cfg.Envs.FileStoragePath)
			h.ServeHTTP(w, request)
			res := w.Result()

			// проверяем код ответа
			assert.Equal(t, tt.want.statusCode, res.StatusCode)
			// заголовок ответа
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))

			// получаем и проверяем тело запроса
			if tt.want.respMessage != "" {
				defer res.Body.Close()
				urlResult, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)

				assert.Equal(t, tt.want.respMessage, string(urlResult))
			}
		})
	}
}
