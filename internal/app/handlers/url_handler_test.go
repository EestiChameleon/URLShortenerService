package handlers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
* 	Сервер должен быть доступен по адресу: http://localhost:8080
*
* 		Сервер должен предоставлять два эндпоинта: POST / и GET /{id}.
*
* 		Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
*
* 		Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
*
* 		Нужно учесть некорректные запросы и возвращать для них ответ с кодом 400.
 */

func TestURLHandlerPOST(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		response    string
	}

	tests := []struct {
		name       string
		requestURL string
		body       *bytes.Buffer
		want       want
	}{
		{
			name:       "POST test #1: url -> 201",
			requestURL: "http://localhost:8080/",
			body:       bytes.NewBuffer([]byte("https://jwt.io/")),
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  201,
				response:    "",
			},
		},
		{
			name:       "POST test #2: empty url -> 400",
			requestURL: "http://localhost:8080/",
			body:       bytes.NewBuffer([]byte("")),
			want: want{
				contentType: "text/plain; charset=utf-8",
				response:    "empty body\n",
				statusCode:  400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.requestURL, tt.body)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(URLHandler)
			h.ServeHTTP(w, request)
			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			if tt.want.response != "" {
				urlResult, err := ioutil.ReadAll(result.Body)
				require.NoError(t, err)
				err = result.Body.Close()
				require.NoError(t, err)

				assert.Equal(t, tt.want.response, string(urlResult))
			}

		})
	}
}

func TestURLHandlerGET(t *testing.T) {
	type want struct {
		contentType    string
		statusCode     int
		response       string
		headerLocation string
	}

	tests := []struct {
		name       string
		requestURL string
		want       want
	}{
		{
			name:       "GET test #1: test url -> 307",
			requestURL: "http://localhost:8080/test",
			want: want{
				contentType:    "",
				statusCode:     307,
				headerLocation: "https://jwt.io/",
				response:       "",
			},
		},
		{
			name:       "GET test #2: wrong url -> 400",
			requestURL: "http://localhost:8080/",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  400,
				response:    "provided url id is not valid\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.requestURL, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(URLHandler)
			h.ServeHTTP(w, request)
			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			if tt.want.response != "" {
				urlResult, err := ioutil.ReadAll(result.Body)
				require.NoError(t, err)
				err = result.Body.Close()
				require.NoError(t, err)

				assert.Equal(t, tt.want.response, string(urlResult))
			}

			if tt.want.headerLocation == "test" {
				assert.Equal(t, tt.want.headerLocation, result.Header.Get("Location"))
			}

		})
	}
}

// отдельный тест для связки пост и гет
//{
//name:    "POST->GET test #1: full logic -> 201 & 307",
//requestMethod: http.MethodPost,
//requestURL: "http://localhost:8080/",
//want: want{
//contentType:    "text/plain",
//postStatusCode: 201,
//getStatusCode:  307,
//postURL:        "https://jwt.io/",
//getURL:         "https://jwt.io/",
//},
//},
