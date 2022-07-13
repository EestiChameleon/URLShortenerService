package handlers

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	resp "github.com/EestiChameleon/URLShortenerService/internal/app/server/responses"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

const testJSON = `[{"correlation_id":"uSKqzeg","original_url":"http://uSKqzeg.com"},{"correlation_id":"vOPtWfu","original_url":"http://vOPtWfu.com"},{"correlation_id":"qfitywe","original_url":"http://qfitywe.com"},{"correlation_id":"CWXwYBO","original_url":"http://CWXwYBO.com"},{"correlation_id":"lilLBon","original_url":"http://lilLBon.com"},{"correlation_id":"YSbkZRY","original_url":"http://YSbkZRY.com"},{"correlation_id":"RfjhFdE","original_url":"http://RfjhFdE.com"},{"correlation_id":"LuJydSD","original_url":"http://LuJydSD.com"},{"correlation_id":"XZYdATu","original_url":"http://XZYdATu.com"},{"correlation_id":"kJAXkID","original_url":"http://kJAXkID.com"},{"correlation_id":"BbeSZIN","original_url":"http://BbeSZIN.com"},{"correlation_id":"FUzUXOS","original_url":"http://FUzUXOS.com"},{"correlation_id":"iFxvoNJ","original_url":"http://iFxvoNJ.com"},{"correlation_id":"cJTpoFh","original_url":"http://cJTpoFh.com"},{"correlation_id":"gGkXxlK","original_url":"http://gGkXxlK.com"},{"correlation_id":"rMfeIJk","original_url":"http://rMfeIJk.com"},{"correlation_id":"iZPYqqc","original_url":"http://iZPYqqc.com"},{"correlation_id":"gySJuDH","original_url":"http://gySJuDH.com"},{"correlation_id":"WGNCzwI","original_url":"http://WGNCzwI.com"},{"correlation_id":"ItQwUND","original_url":"http://ItQwUND.com"},{"correlation_id":"QFmKNlF","original_url":"http://QFmKNlF.com"},{"correlation_id":"IbpIYQo","original_url":"http://IbpIYQo.com"},{"correlation_id":"UbkOdMF","original_url":"http://UbkOdMF.com"},{"correlation_id":"DMwjKdO","original_url":"http://DMwjKdO.com"},{"correlation_id":"OEoHxSm","original_url":"http://OEoHxSm.com"},{"correlation_id":"XdtmQbc","original_url":"http://XdtmQbc.com"},{"correlation_id":"rqPPKUO","original_url":"http://rqPPKUO.com"},{"correlation_id":"lsqVFQx","original_url":"http://lsqVFQx.com"},{"correlation_id":"NWrXKsb","original_url":"http://NWrXKsb.com"},{"correlation_id":"RQzECwR","original_url":"http://RQzECwR.com"},{"correlation_id":"kJfhBoV","original_url":"http://kJfhBoV.com"},{"correlation_id":"VIpUsZl","original_url":"http://VIpUsZl.com"},{"correlation_id":"WHCzkyR","original_url":"http://WHCzkyR.com"},{"correlation_id":"Juhdycv","original_url":"http://Juhdycv.com"},{"correlation_id":"tddpXZi","original_url":"http://tddpXZi.com"},{"correlation_id":"AMfRCuO","original_url":"http://AMfRCuO.com"},{"correlation_id":"btGWxUG","original_url":"http://btGWxUG.com"},{"correlation_id":"SpixSMH","original_url":"http://SpixSMH.com"},{"correlation_id":"ICkjlTD","original_url":"http://ICkjlTD.com"},{"correlation_id":"ydsAYZD","original_url":"http://ydsAYZD.com"},{"correlation_id":"JdsXxQr","original_url":"http://JdsXxQr.com"},{"correlation_id":"qlTWLyf","original_url":"http://qlTWLyf.com"},{"correlation_id":"nXBvlMQ","original_url":"http://nXBvlMQ.com"},{"correlation_id":"QNRaHPY","original_url":"http://QNRaHPY.com"},{"correlation_id":"oRkukUT","original_url":"http://oRkukUT.com"},{"correlation_id":"rMsZBhJ","original_url":"http://rMsZBhJ.com"},{"correlation_id":"gKSJIIn","original_url":"http://gKSJIIn.com"},{"correlation_id":"UfSElzu","original_url":"http://UfSElzu.com"},{"correlation_id":"nZqNaOk","original_url":"http://nZqNaOk.com"},{"correlation_id":"YsKZcMt","original_url":"http://YsKZcMt.com"},{"correlation_id":"GOuGbhJ","original_url":"http://GOuGbhJ.com"},{"correlation_id":"gXCJsYG","original_url":"http://gXCJsYG.com"},{"correlation_id":"hLGXTCA","original_url":"http://hLGXTCA.com"},{"correlation_id":"bNhiPKy","original_url":"http://bNhiPKy.com"},{"correlation_id":"aHxtetX","original_url":"http://aHxtetX.com"},{"correlation_id":"OAAkRFm","original_url":"http://OAAkRFm.com"},{"correlation_id":"iMeefnU","original_url":"http://iMeefnU.com"},{"correlation_id":"HVplHEq","original_url":"http://HVplHEq.com"},{"correlation_id":"FETiALv","original_url":"http://FETiALv.com"},{"correlation_id":"reQRLYR","original_url":"http://reQRLYR.com"},{"correlation_id":"rtbRnsv","original_url":"http://rtbRnsv.com"},{"correlation_id":"JXWldqa","original_url":"http://JXWldqa.com"},{"correlation_id":"vyGMUvP","original_url":"http://vyGMUvP.com"},{"correlation_id":"LwUEvKI","original_url":"http://LwUEvKI.com"},{"correlation_id":"GhaxyUg","original_url":"http://GhaxyUg.com"},{"correlation_id":"geLyLgR","original_url":"http://geLyLgR.com"},{"correlation_id":"KvVxuKP","original_url":"http://KvVxuKP.com"},{"correlation_id":"tGSWmfK","original_url":"http://tGSWmfK.com"},{"correlation_id":"ysGazXt","original_url":"http://ysGazXt.com"},{"correlation_id":"wlTvowM","original_url":"http://wlTvowM.com"},{"correlation_id":"wTJRcnV","original_url":"http://wTJRcnV.com"},{"correlation_id":"JmGYgwr","original_url":"http://JmGYgwr.com"},{"correlation_id":"cpGIhZT","original_url":"http://cpGIhZT.com"},{"correlation_id":"XpfXbjl","original_url":"http://XpfXbjl.com"},{"correlation_id":"cJBrEUc","original_url":"http://cJBrEUc.com"},{"correlation_id":"SBEGIMR","original_url":"http://SBEGIMR.com"},{"correlation_id":"HoJgHFv","original_url":"http://HoJgHFv.com"},{"correlation_id":"kdhDzCl","original_url":"http://kdhDzCl.com"},{"correlation_id":"myokRhl","original_url":"http://myokRhl.com"},{"correlation_id":"pQEbXCS","original_url":"http://pQEbXCS.com"},{"correlation_id":"YqwNdwW","original_url":"http://YqwNdwW.com"},{"correlation_id":"lLDBqjB","original_url":"http://lLDBqjB.com"},{"correlation_id":"qASoOYd","original_url":"http://qASoOYd.com"},{"correlation_id":"EGCPmvm","original_url":"http://EGCPmvm.com"},{"correlation_id":"FDuvUtP","original_url":"http://FDuvUtP.com"},{"correlation_id":"AIDWVup","original_url":"http://AIDWVup.com"},{"correlation_id":"twmJzTO","original_url":"http://twmJzTO.com"},{"correlation_id":"zqQuJMY","original_url":"http://zqQuJMY.com"},{"correlation_id":"NLHGZcb","original_url":"http://NLHGZcb.com"},{"correlation_id":"qhvcIKF","original_url":"http://qhvcIKF.com"},{"correlation_id":"dVcNhkr","original_url":"http://dVcNhkr.com"},{"correlation_id":"VxhOhdV","original_url":"http://VxhOhdV.com"},{"correlation_id":"EYICOBE","original_url":"http://EYICOBE.com"},{"correlation_id":"LqEwler","original_url":"http://LqEwler.com"},{"correlation_id":"xBgEWDO","original_url":"http://xBgEWDO.com"},{"correlation_id":"IPEXzZf","original_url":"http://IPEXzZf.com"},{"correlation_id":"hvATlgs","original_url":"http://hvATlgs.com"},{"correlation_id":"JYStPTR","original_url":"http://JYStPTR.com"},{"correlation_id":"tYNifIx","original_url":"http://tYNifIx.com"}]`

type TestBatchData struct {
	CorID   string `json:"correlation_id"`
	OrigURL string `json:"original_url"`
}

//var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//
//func randSeq(n int) string {
//	b := make([]rune, n)
//	for i := range b {
//		b[i] = letters[rand.Intn(len(letters))]
//	}
//	return string(b)
//}

func BenchmarkPostBatch(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	//var batchData []TestBatchData

	b.Run("Endpoint: POST /api/shorten/batch", func(b *testing.B) {

		//for i := 1; i < 100; i++ {
		//	id := randSeq(7)
		//	batchData = append(batchData, TestBatchData{
		//		CorID:   id,
		//		OrigURL: fmt.Sprintf("http://%s.com", id),
		//	})
		//}
		//jsonBatchData, err := json.Marshal(batchData)
		//if err != nil {
		//	log.Fatal(err)
		//}

		path := "http://localhost:8080/api/shorten/batch"
		contentType := resp.MIMEApplicationJSONCharsetUTF8

		request := httptest.NewRequest(http.MethodPost, path, strings.NewReader(testJSON))
		request.Header.Set(resp.HeaderContentType, contentType)

		// создаём новый Recorder
		w := httptest.NewRecorder()
		// определяем хендлер
		h := http.HandlerFunc(PostBatch)
		// envs
		cfg.Envs.BaseURL = "http://localhost:8080"
		//cfg.Envs.FileStoragePath = "tmp/testFile"
		//cfg.Envs.DatabaseDSN = "postgresql://localhost:5432/yandex_practicum_db"
		// запускаем сервер
		if err := storage.InitStorage(); err != nil {
			log.Fatal(err)
		}
		defer os.Remove(cfg.Envs.FileStoragePath)

		b.ResetTimer()

		h.ServeHTTP(w, request)
	})
}
