package custommw

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/models"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
	"net/http"
)

func CheckCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("CheckCookie: start getcookie")
		cookie, err := r.Cookie("UserID")
		if err != nil {
			log.Println("CheckCookie r.Cookie(\"UserID\") err: ", err)
			userID, err := models.CreateUserID()
			if err != nil {
				resp.NoContent(w, http.StatusInternalServerError)
				return
			}
			log.Println("CheckCookie storage.User.ID created & saved: ", userID)
			storage.User.ID = userID
			token, err := models.JWTEncode("userID", userID)
			if err != nil {
				resp.NoContent(w, http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, resp.CreateCookie("UserID", token))
			log.Print("UserID cookie added, new storage.User.ID saved")
			next.ServeHTTP(w, r)
			return
		}
		log.Println("CheckCookie: cookie found - ", cookie)
		userID, err := models.JWTDecode(cookie.Value, "userID")
		if err != nil {
			log.Println("cookie JWTDecode err: ", err)
			resp.NoContent(w, http.StatusInternalServerError)
			return
		}
		log.Println("CheckCookie storage.User.ID decoded & saved: ", userID)
		storage.User.ID = userID
		log.Print("UserID cookie found & decoded, storage.User.ID saved")
		next.ServeHTTP(w, r)
	})
}
