package cfg

import (
	"github.com/caarlos0/env/v6"
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
	return env.Parse(&Envs)
}
