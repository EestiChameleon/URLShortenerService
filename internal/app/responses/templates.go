package responses

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

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

func NoContent(w http.ResponseWriter, code int) {
	log.Println("templates NoContent start. Code: ", code)
	w.WriteHeader(code)
}

func WriteString(w http.ResponseWriter, code int, s string) {
	log.Printf("templates WriteString start. Code: %v, String: %s\n", code, s)
	w.Header().Set(HeaderContentType, MIMETextPlainCharsetUTF8)
	w.WriteHeader(code)
	w.Write([]byte(s)) //nolint:errcheck    // Проверять ошибку здесь смысла нет: если ты не можешь записать байтики в подключение, то ответ со статусом вернуть тоже не сможешь, тк соединение уже вероятно разорвано) (c)
}

// RedirectString send a redirect header to the indicated link - s
func RedirectString(w http.ResponseWriter, s string) {
	log.Printf("templates RedirectString start. String: %s\n", s)
	w.Header().Set(HeaderContentType, MIMETextPlainCharsetUTF8)
	w.Header().Set(HeaderLocation, s)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// CreateCookie func provides a cookie "key=value" based on given params
func CreateCookie(key string, value string) *http.Cookie {
	log.Printf("templates CreateCookie start. Key: %s, Value: %s\n", key, value)
	return &http.Cookie{
		Name:    key,
		Value:   value,
		Path:    "/",
		Expires: time.Now().Add(time.Second * 60 * 60),
	}
}
