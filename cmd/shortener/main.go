package main

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/EestiChameleon/URLShortenerService/internal/app/server"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
)

func init() {
	// get envs
	cfg.GetEnvs()
}

func main() {
	// get stored pairs
	if err := storage.Pairs.GetFile(); err != nil {
		panic(err)
	}

	// start the server
	server.Start()

}
