package app

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net/http"
)

var store = &Store{db: map[string]string{}}

func URLHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		// GET /{id}
		urlID := r.URL.Path[len("/"):]

		shortedURL := fmt.Sprintf("localhost:8080/%v", urlID)

		if !checkURLID(shortedURL) {
			http.Error(w, "provided url id is not valid", 400)
			return
		}

		w.Header().Set("Location", store.db[shortedURL])
		w.WriteHeader(http.StatusTemporaryRedirect)

	case "POST": // POST / in body = url to short
		responseData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "empty body", 400)
			return
		}

		reqURL := string(responseData)
		if reqURL == "" {
			http.Error(w, "empty body", 400)
			return
		}

		shortedURLID, err := shortUrlID()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		shortedURL := fmt.Sprintf("localhost:8080/%v", shortedURLID)

		store.Put(shortedURL, reqURL)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)

		_, err = w.Write([]byte(shortedURL))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

	default:
		// return err if the method isn't GET or POST
		w.WriteHeader(http.StatusBadRequest)
	}
}

func shortUrlID() (shortedURL string, err error) {
	b := make([]byte, 7)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	shortedURL = fmt.Sprintf("%x", b[0:])
	return
}

func checkURLID(id string) bool {
	for k := range store.db {
		if id == k {
			return true
		}
	}
	return false
}
