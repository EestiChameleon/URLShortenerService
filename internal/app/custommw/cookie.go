package custommw

import (
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
	"net/http"
	"net/url"
)

func CheckCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("UserID")
		if err != nil {
			log.Println(err)
			userID, err := storage.CreateUserID()
			if err != nil {
				resp.NoContent(w, http.StatusInternalServerError)
				return
			}
			encID, err := storage.JWTEncode("userID", userID)
			if err != nil {
				resp.NoContent(w, http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, resp.CreateCookie("UserID", url.QueryEscape(encID)))
			storage.User.ID = userID
			log.Print("UserID cookie was missing - added, new storage.User.ID saved")
			next.ServeHTTP(w, r)
			return
		}
		log.Println(cookie)
		userID, err := storage.JWTDecode(cookie.Value, "userID")
		if err != nil {
			log.Println(err)
			resp.NoContent(w, http.StatusInternalServerError)
			return
		}

		storage.User.ID = userID
		log.Print("UserID cookie found, storage.User.ID saved")
		next.ServeHTTP(w, r)
	})
}
