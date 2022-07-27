package main

import (
	"fmt"
	"log"
	"os"

	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/EestiChameleon/URLShortenerService/internal/app/server"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
)

func init() {

}

var (
	buildVersion = `N/A`
	buildDate    = `N/A`
	buildCommit  = `N/A`
)

func main() { //nolint:typecheck
	// incr 19
	fmt.Fprintf(os.Stdout, "Build version: %s\nBuild date: %s\nBuild commit: %s\n", buildVersion, buildDate, buildCommit)

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
