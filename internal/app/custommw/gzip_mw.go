package custommw

import (
	"compress/gzip"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"io"
	"io/ioutil"
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
		// проверяем, что клиент поддерживает gzip-сжатие
		if !strings.Contains(r.Header.Get(resp.HeaderAcceptEncoding), "gzip") {
			log.Println("request accept-encoding != gzip")
			// если gzip не поддерживается, передаём управление
			// дальше без изменений
			next.ServeHTTP(w, r)
			return
		}
		log.Println("request accept-encoding = gzip")
		// создаём gzip.Writer поверх текущего w
		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			resp.NoContent(w, http.StatusBadRequest)
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		log.Println("w.header Content-encoding = gzip")
		// передаём обработчику страницы переменную типа gzipWriter для вывода данных
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
		log.Println("passed gzip w -> next")
	})
}

// RequestGZIP - middleware that decompress request.body when header Content-Encoding = gzip
func RequestGZIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hce := r.Header.Get(resp.HeaderContentEncoding)
		log.Printf("request header content encoding: %s", hce)
		if r.Header.Get(resp.HeaderContentEncoding) == "gzip" {
			log.Println("start decompressing")

			if rb, err := ioutil.ReadAll(r.Body); err != nil {
				log.Println("failed to read the body")
				resp.WriteString(w, http.StatusInternalServerError, err.Error())
				return
			} else {
				log.Println(string(rb))
			}

			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				log.Println("failed to read the body 2 ")
				resp.WriteString(w, http.StatusInternalServerError, err.Error())
				return
			}
			log.Println("request body changed with decompressed")
			r.Body = gz
			defer gz.Close()
		}

		next.ServeHTTP(w, r)
	})
}
