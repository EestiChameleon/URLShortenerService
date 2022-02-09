package main

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/server"
	"log"
)

func main() {

	log.Fatal(server.Start())
}
