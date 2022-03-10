package custommw

import (
	"compress/gzip"
	"context"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"io"
	"log"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за gzip-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

// ResponseGZIP - middleware that provides http.ResponseWriter with gzip Writer
// when header Accept-Encoding contains gzip and set w.header Content-Encoding = gzip
func ResponseGZIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get(resp.HeaderAcceptEncoding), "gzip") {
			log.Println("request accept-encoding != gzip")
			next.ServeHTTP(w, r)
			return
		}
		log.Println("request accept-encoding contains gzip")
		// создаём gzip.Writer поверх текущего w
		gz, err := gzip.NewWriterLevel(w, gzip.DefaultCompression)
		if err != nil {
			resp.NoContent(w, http.StatusBadRequest)
			return
		}
		defer gz.Close()

		w.Header().Set(resp.HeaderContentEncoding, "gzip")
		log.Println("response header set -> Content-encoding = gzip")
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

// RequestGZIP - middleware that decompress request.body when header Content-Encoding = gzip
func RequestGZIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reader io.Reader
		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			reader = gz
			defer gz.Close()
		} else {
			reader = r.Body
			defer r.Body.Close()
		}

		bodyRaw, err := io.ReadAll(reader)
		if err != nil || len(bodyRaw) == 0 {
			http.Error(w, "url missing in the body", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "bodyURL", bodyRaw)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
