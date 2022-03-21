package handlers

import (
	"encoding/json"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"io"
	"log"
	"net/http"
)

type BatchReqPair struct {
	CorID   string `json:"correlation_id"`
	OrigURL string `json:"original_url"`
}

type BatchRespPair struct {
	CorID    string `json:"correlation_id"`
	ShortURL string `json:"short_url"`
}

type BatchReq []BatchReqPair

type BatchResp []BatchRespPair

// PostBatch принимает в теле запроса JSON-объект -
//	[
//		{
//			"correlation_id": "<строковый идентификатор>",
//			"original_url": "<URL для сокращения>"
//		},
//	]

// В качестве ответа PostBatch должен возвращать данные в формате:
//	[
//		{
//		 "correlation_id": "<строковый идентификатор из объекта запроса>",
//		"short_url": "<результирующий сокращённый URL>"
//		},
//	   ...
//	]

func PostBatch(w http.ResponseWriter, r *http.Request) {
	var (
		reqBody  BatchReq
		respBody BatchResp
	)

	// read body
	log.Println("PostBatch: start - read r.Body")
	byteBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("PostBatch: unable to read body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	log.Println("PostBatch: json.Unmarshal(byteBody, &reqBody)")
	if err = json.Unmarshal(byteBody, &reqBody); err != nil {
		log.Println("PostBatch: unable to unmarshal body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	for _, v := range reqBody {
		// check if it's not empty
		origURL := v.OrigURL
		if origURL == "" {
			log.Println("PostBatch: empty r.Body")
			resp.WriteString(w, http.StatusBadRequest, "invalid data")
			return
		}

		// get a short url to pair with the orig url
		shortURL, err := storage.User.CreateShortURL()
		if err != nil {
			log.Println("PostBatch: GetShortURL err:", err)
			resp.WriteString(w, http.StatusBadRequest, "invalid data")
			return
		}

		if err = storage.User.SavePair(storage.Pair{ShortURL: shortURL, OrigURL: origURL}); err != nil {
			log.Println("PostBatch: storage.User.SavePair err:", err)
			resp.WriteString(w, http.StatusBadRequest, "invalid data")
			return
		}

		respBody = append(respBody, BatchRespPair{
			CorID:    v.CorID,
			ShortURL: shortURL,
		})
	}

	log.Println("PostBatch: OK")
	resp.JSON(w, http.StatusCreated, respBody)
}
