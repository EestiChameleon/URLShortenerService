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

	flag.StringVar(&cfg.Envs.SrvAddr, "a", "localhost:8080", "SERVER_ADDRESS to listen on")
	flag.StringVar(&cfg.Envs.BaseURL, "b", "http://localhost:8080", "BASE_URL of the shorten result URL")
	flag.StringVar(&cfg.Envs.FileStoragePath, "f", "tmp/pairs", "FILE_STORAGE_PATH. Directory of the origin&shorten url pairs file")
	// get envs
	if err := cfg.GetEnvs(); err != nil {
		log.Fatal(err)
	}
	flag.Parse()

	// get stored pairs
	if err := storage.Pairs.GetFile(); err != nil {
		log.Fatal(err)
	}

	// start the server
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
