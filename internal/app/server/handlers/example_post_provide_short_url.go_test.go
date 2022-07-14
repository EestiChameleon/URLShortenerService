package handlers

import (
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/responses"
	"log"
	"net/http"
	"strings"
)

func ExamplePostProvideShortURL() {
	_, err := http.Post(
		"http://localhost:8080/",
		resp.MIMETextPlainCharsetUTF8,
		strings.NewReader("https://jwt.io/")) // POST "/"
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// # Request
	// POST / HTTP/1.1
	// Content-Type: text/plain
	//
	// # Response
	// HTTP/1.1 201 Created
	// Content-Type: text/plain; charset=utf-8
	// Body: http://localhost:8080/test
}
