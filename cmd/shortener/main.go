package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	// channel to alert about shutdown
	cracefullShutdownChan := make(chan struct{})
	// channel to redirect the interrupt
	// we are looking after 3 syscall
	sigint := make(chan os.Signal, 3) // or size could be 1?
	// redirect registration
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// launch goroutine for received interrupt
	go func() {
		// we need only 1 signal to start the procedure
		<-sigint

		if err := server.Shutdown(); err != nil {
			// ошибки закрытия Listener
			log.Printf("HTTP server Shutdown err: %v", err)
		}
		// сообщаем основному потоку,
		// что все сетевые соединения обработаны и закрыты
		close(cracefullShutdownChan)
	}()

	// database initiation
	log.Println("[INFO] main -> storage.InitStorage()")
	if err := storage.InitStorage(); err != nil {
		log.Fatal(err)
	}

	// start the server
	log.Println("[INFO] main -> server.Start()")
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	// waiting the end of graceful shutdown procedure
	<-cracefullShutdownChan
	// получили оповещение о завершении
	// здесь можно освобождать ресурсы перед выходом,
	// например закрыть соединение с базой данных,
	// закрыть открытые файлы

	if err := storage.User.Shutdown(); err != nil {
		log.Printf("Database Shutdown err: %v", err)
	}

	log.Println("Server Shutdown gracefully")
}
