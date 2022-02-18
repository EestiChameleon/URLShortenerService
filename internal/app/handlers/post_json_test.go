package handlers

import (
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJSONShortURL(t *testing.T) {
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
			name:        "POST test #1: correct json -> 201",
			contentType: resp.MIMEApplicationJSONCharsetUTF8,
			request:     "http://localhost:8080/api/shorten",
			body:        `{"url":"https://jwt.io/"}`,
			want: want{
				contentType: resp.MIMEApplicationJSONCharsetUTF8,
				statusCode:  201,
				respMessage: "",
			},
		},
		{
			name:        "POST test #2: wrong incoming json -> 400",
			contentType: resp.MIMEApplicationJSONCharsetUTF8,
			request:     "http://localhost:8080/",
			body:        `{"urly":"https://jwt.io/"}`,
			want: want{
				contentType: resp.MIMETextPlainCharsetUTF8,
				respMessage: "invalid data",
				statusCode:  400,
			},
		},
		{
			name:        "POST test #3: wrong content type -> 400",
			contentType: resp.MIMETextPlainCharsetUTF8,
			request:     "http://localhost:8080/",
			body:        `{"url":"https://jwt.io/"}`,
			want: want{
				contentType: resp.MIMETextPlainCharsetUTF8,
				statusCode:  400,
				respMessage: "invalid data",
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
			h := http.HandlerFunc(JSONShortURL)
			// запускаем сервер
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
