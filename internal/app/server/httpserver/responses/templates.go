// Package responses simplify the API interaction.
package responses

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Header and Content-Type const to simplify the choice. No more typo etc.
const (
	// HEADERS ------------------------------------
	HeaderContentType     = "Content-Type"
	HeaderLocation        = "Location"
	HeaderXForwardedFor   = "X-Forwarded-For"
	HeaderXRealIP         = "X-Real-IP"
	HeaderAcceptEncoding  = "Accept-Encoding"
	HeaderContentEncoding = "Content-Encoding"

	// CONTENT TYPE -------------------------------------------------------------
	charsetUTF8 = "charset=utf-8"

	MIMEApplicationJSON            = "application/json"
	MIMEApplicationJSONCharsetUTF8 = MIMEApplicationJSON + "; " + charsetUTF8
	MIMEApplicationXML             = "application/xml"
	MIMEApplicationXMLCharsetUTF8  = MIMEApplicationXML + "; " + charsetUTF8
	MIMETextXML                    = "text/xml"
	MIMETextXMLCharsetUTF8         = MIMETextXML + "; " + charsetUTF8
	MIMEApplicationForm            = "application/x-www-form-urlencoded"
	MIMETextPlain                  = "text/plain"
	MIMETextPlainCharsetUTF8       = MIMETextPlain + "; " + charsetUTF8
	MIMEMultipartForm              = "multipart/form-data"
)

// JSON function provides a simple way for handler to response with a JSON.
// Just provide the http.StatusCode and interface to json.Marshal.
func JSON(w http.ResponseWriter, code int, i interface{}) {
	log.Printf("templates JSON start. Code: %v, Interface: %v\n", code, i)
	data, err := json.Marshal(i)
	if err != nil {
		log.Println("template JSON json.Marshal(i) err: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set(HeaderContentType, MIMEApplicationJSONCharsetUTF8)
	w.WriteHeader(code)
	w.Write(data) //nolint:errcheck
}

// NoContent sends nil body. Only http.StatusCode.
func NoContent(w http.ResponseWriter, code int) {
	log.Println("templates NoContent start. Code: ", code)
	w.WriteHeader(code)
}

// WriteString sends a response with a text.
// Just provide the http.StatusCode and string to send.
func WriteString(w http.ResponseWriter, code int, s string) {
	log.Printf("templates WriteString start. Code: %v, String: %s\n", code, s)
	w.Header().Set(HeaderContentType, MIMETextPlainCharsetUTF8)
	w.WriteHeader(code)
	w.Write([]byte(s)) //nolint:errcheck    // Проверять ошибку здесь смысла нет: если ты не можешь записать байтики в подключение, то ответ со статусом вернуть тоже не сможешь, тк соединение уже вероятно разорвано) (c)
}

// RedirectString send a response with a redirect header to the indicated link.
func RedirectString(w http.ResponseWriter, s string) {
	log.Printf("templates RedirectString start. String: %s\n", s)
	w.Header().Set(HeaderContentType, MIMETextPlainCharsetUTF8)
	w.Header().Set(HeaderLocation, s)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// CreateCookie function creates a cookie "key=value" based on given params. Path = "/". Expiration time is set to 1h.
func CreateCookie(key string, value string) *http.Cookie {
	log.Printf("templates CreateCookie start. Key: %s, Value: %s\n", key, value)
	return &http.Cookie{
		Name:    key,
		Value:   value,
		Path:    "/",
		Expires: time.Now().Add(time.Second * 60 * 60),
	}
}
