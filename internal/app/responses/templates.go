package responses

import (
	"encoding/json"
	"net/http"
)

const (
	// HEADERS ------------------------------------
	HeaderContentType   = "Content-Type"
	HeaderLocation      = "Location"
	HeaderXForwardedFor = "X-Forwarded-For"
	HeaderXRealIP       = "X-Real-IP"

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
	data, err := json.Marshal(i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set(HeaderContentType, MIMEApplicationJSONCharsetUTF8)
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func NoContent(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

func WriteString(w http.ResponseWriter, code int, s string) {
	w.Header().Set(HeaderContentType, MIMETextPlainCharsetUTF8)
	w.WriteHeader(code)
	w.Write([]byte(s)) // Проверять ошибку здесь смысла нет: если ты не можешь записать байтики в подключение, то ответ со статусом вернуть тоже не сможешь, тк соединение уже вероятно разорвано) (c)
}

// RedirectString send a redirect header to the indicated link - s
func RedirectString(w http.ResponseWriter, s string) {
	w.Header().Set(HeaderContentType, MIMETextPlainCharsetUTF8)
	w.Header().Set(HeaderLocation, s)
	w.WriteHeader(http.StatusTemporaryRedirect)
	return
}
