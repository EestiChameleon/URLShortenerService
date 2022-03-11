package cfg

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	SrvAddr         string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`      //флаг -a, отвечающий за адрес запуска HTTP-сервера (переменная SERVER_ADDRESS);
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`     //флаг -b, отвечающий за базовый адрес результирующего сокращённого URL (переменная BASE_URL);
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"tmp/urlPairsData"` //флаг -f, отвечающий за путь до файла с сокращёнными URL (переменная FILE_STORAGE_PATH).
	CryptoKey       string `env:"CRYPTO_KEY" envDefault:"secret_123456789"`        //secret word to encrypt/decrypt cookies
}

var Envs Config

type ContextKey string

func GetEnvs() error {
	return env.Parse(&Envs)
}
