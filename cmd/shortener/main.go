package main

import (
	"flag"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/EestiChameleon/URLShortenerService/internal/app/server"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
)

func init() {

}

func main() {
	// flags declaration
	log.Println("main: flag declaration start")
	flag.StringVar(&cfg.Envs.SrvAddr, "a", "localhost:8080", "SERVER_ADDRESS to listen on")
	flag.StringVar(&cfg.Envs.BaseURL, "b", "http://localhost:8080", "BASE_URL of the shorten result URL")
	flag.StringVar(&cfg.Envs.FileStoragePath, "f", "tmp/urlPairsData", "FILE_STORAGE_PATH. Directory of the origin&shorten url pairs file")
	flag.StringVar(&cfg.Envs.DatabaseDSN, "d", "postgresql://localhost:5432/yandex_practicum_db", "DATABASE_DSN. Address for connection to DB")

	// get envs
	log.Println("main: GetEnvs start")
	if err := cfg.GetEnvs(); err != nil {
		log.Fatal(err)
	}

	log.Println("main: flag parse start")
	flag.Parse()

	// database initiation
	log.Println("main: start InitStorage")
	if err := storage.User.InitStorage(); err != nil {
		log.Fatal(err)
	}
	defer storage.User.CloseStorage() // save data before exit

	// start the server
	log.Println("main: start Server")
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
