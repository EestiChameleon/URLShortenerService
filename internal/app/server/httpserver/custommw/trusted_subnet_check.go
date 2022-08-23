package custommw

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/httpserver/responses"
	"log"
	"net"
	"net/http"
)

// TrustSubnetCheck - middleware that verify if the endpoint is called by a trusted subnet form Env.TrustedSubnet list
func TrustSubnetCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cfg.Envs.TrustedSubnet == `` {
			log.Println("TrustSubnetCheck: no trusted subnet. Access denied.")
			resp.NoContent(w, http.StatusForbidden)
			return
		}

		// смотрим заголовок запроса X-Real-IP
		ipStr := r.Header.Get("X-Real-IP")
		// парсим ip
		ip := net.ParseIP(ipStr)
		if ip == nil {
			log.Println("TrustSubnetCheck: no IP in the `X-Real-IP` Header. Access denied.")
			resp.NoContent(w, http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
