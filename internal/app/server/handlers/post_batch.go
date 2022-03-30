package handlers

import (
	"encoding/json"
	"errors"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/service/process"
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
		shortURL string
	)

	// read body
	log.Println("[INFO] handlers -> PostBatch: start - read r.Body")
	byteBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR] handlers -> PostBatch: unable to read body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	log.Println("[DEBUG] handlers -> PostBatch: json.Unmarshal(byteBody, &reqBody)")
	if err = json.Unmarshal(byteBody, &reqBody); err != nil {
		log.Println("PostBatch: unable to unmarshal body:", err)
		resp.WriteString(w, http.StatusBadRequest, "invalid data")
		return
	}

	for _, v := range reqBody {
		// check if it's not empty
		origURL := v.OrigURL
		if origURL == "" {
			log.Println("[ERROR] handlers -> PostBatch: empty r.Body")
			resp.WriteString(w, http.StatusBadRequest, "invalid data")
			return
		}

		// get a short url to pair with the orig url
		shortURL, err = process.ShortURLforOrigURL(origURL)
		if err != nil && !errors.Is(err, storage.ErrDBOrigURLExists) {
			log.Println("[ERROR] handlers -> JSONShortURL: ShortURLforOrigURL err:", err)
			resp.WriteString(w, http.StatusBadRequest, "invalid data")
			return
		}

		respBody = append(respBody, BatchRespPair{
			CorID:    v.CorID,
			ShortURL: shortURL,
		})
	}

	log.Println("[INFO] handlers -> PostBatch: OK")
	resp.JSON(w, http.StatusCreated, respBody)
}
