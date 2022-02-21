package main

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/EestiChameleon/URLShortenerService/internal/app/server"
)

func main() {

	// get envs
	cfg.Envs = cfg.GetEnvs()
	//cfg.Envs = cfg.Config{BaseURL: "localhost:8081", SrvAddr: "localhost:8081"}
	// start the server
	server.Start()

}
