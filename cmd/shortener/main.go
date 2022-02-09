package main

import (
	"github.com/EestiChameleon/URLShortenerService/internal/app/server"
	"log"
)

func main() {
	// start the server and catch err if it comes
	log.Fatal(server.Start())
}
