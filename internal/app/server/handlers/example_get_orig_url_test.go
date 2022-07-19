package handlers

func ExampleGetOrigURL() {
	//request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/test", nil)
	//// envs
	//cfg.Envs.BaseURL = "http://localhost:8080"
	////cfg.Envs.FileStoragePath = "tmp/testFile"
	////cfg.Envs.DatabaseDSN = "postgresql://localhost:5432/yandex_practicum_db"
	//if err := storage.InitStorage(); err != nil {
	//	log.Fatal(err)
	//}
	//storage.User.SetUserID("testUser")
	//storage.User.SavePair(storage.Pair{
	//	ShortURL: "http://localhost:8080/test",
	//	OrigURL:  "https://jwt.io/",
	//})
	//defer os.Remove(cfg.Envs.FileStoragePath)
	//
	//// создаём новый Recorder
	//w := httptest.NewRecorder()
	//r := chi.NewRouter()
	//// определяем хендлер
	//r.Get("/{id}", GetOrigURL)
	//// запускаем сервер
	//r.ServeHTTP(w, request)
	//res := w.Result()
	//defer res.Body.Close()

	// Output:
	// # Request
	// GET /test HTTP/1.1
	//
	// # Response
	// HTTP/1.1 307 Temporary Redirect
	// Content-Type: text/plain; charset=utf-8
	// Location: https://jwt.io/
}
