package custommw

import (
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/service/data"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
	"net/http"
)

func CheckCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("[INFO] custommw -> CheckCookie: start getcookie")
		cookie, err := r.Cookie("UserID")
		if err != nil {
			log.Println("[ERROR] custommw -> CheckCookie r.Cookie(\"UserID\") err: ", err)
			userID, err := data.CreateUserID()
			if err != nil {
				log.Println("[ERROR] custommw -> CheckCookie CreateUserID err: ", err)
				resp.NoContent(w, http.StatusInternalServerError)
				return
			}
			log.Println("[DEBUG] custommw -> CheckCookie storage.User.ID created & saved: ", userID)
			storage.User.SetUserID(userID)
			token, err := data.JWTEncode("userID", userID)
			if err != nil {
				resp.NoContent(w, http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, resp.CreateCookie("UserID", token))
			log.Print("[DEBUG] custommw -> UserID cookie added, new storage.User.ID saved")
			next.ServeHTTP(w, r)
			return
		}
		log.Println("[DEBUG] custommw -> CheckCookie: cookie found - ", cookie)
		userID, err := data.JWTDecode(cookie.Value, "userID")
		if err != nil {
			log.Println("[ERROR] custommw -> cookie JWTDecode err: ", err)
			resp.NoContent(w, http.StatusInternalServerError)
			return
		}
		log.Println("[DEBUG] custommw -> CheckCookie storage.User.ID decoded & saved: ", userID)
		storage.User.SetUserID(userID)
		log.Print("[DEBUG] custommw -> UserID cookie found & decoded, storage.User.ID saved")
		next.ServeHTTP(w, r)
	})
}
