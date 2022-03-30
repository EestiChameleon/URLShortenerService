package cfg

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	SrvAddr         string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`                                //адрес запуска HTTP-сервера
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`                               //базовый адрес результирующего сокращённого URL
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"tmp/urlPairsData"`                           //путь до файла с сокращёнными URL
	CryptoKey       string `env:"CRYPTO_KEY" envDefault:"secret_123456789"`                                  //secret word to encrypt/decrypt JWT for cookies
	DatabaseDSN     string `env:"DATABASE_DSN" envDefault:"postgresql://localhost:5432/yandex_practicum_db"` //Строка с адресом подключения к БД
}

var Envs Config

type ContextKey string

func GetEnvs() error {
	log.Println("[INFO] cfg -> GetFlag: flag declaration start")
	flag.StringVar(&Envs.SrvAddr, "a", "localhost:8080", "SERVER_ADDRESS to listen on")
	flag.StringVar(&Envs.BaseURL, "b", "http://localhost:8080", "BASE_URL of the shorten result URL")
	flag.StringVar(&Envs.FileStoragePath, "f", "tmp/urlPairsData", "FILE_STORAGE_PATH. Directory of the origin&shorten url pairs file")
	flag.StringVar(&Envs.DatabaseDSN, "d", "postgresql://localhost:5432/yandex_practicum_db", "DATABASE_DSN. Address for connection to DB")

	log.Println("[INFO] cfg -> GetEnvs: parse envs start")
	if err := env.Parse(&Envs); err != nil {
		return err
	}

	log.Println("[INFO] cfg -> GetFlag: flag parse start")
	flag.Parse()

	return nil
}
