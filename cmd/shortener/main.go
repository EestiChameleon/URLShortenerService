package main

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/EestiChameleon/URLShortenerService/internal/app/server"
)

func main() {

	// get envs
	cfg.Envs = cfg.GetEnvs()
	// start the server
	server.Start()

}
