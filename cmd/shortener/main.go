package main

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/EestiChameleon/URLShortenerService/internal/app/server"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"log"
)

func init() {

}

func main() { //nolint:typecheck
	log.Println("[INFO] main -> cfg.GetFlag()")
	cfg.GetFlag()

	log.Println("[INFO] main -> cfg.GetEnvs()")
	if err := cfg.GetEnvs(); err != nil {
		log.Fatal(err)
	}

	// database initiation
	log.Println("[INFO] main -> storage.InitStorage()")
	if err := storage.InitStorage(); err != nil {
		log.Fatal(err)
	}
	defer storage.User.ShutDown() // DB - close connection, Memory - save data and exit

	// start the server
	log.Println("[INFO] main -> server.Start()")
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
