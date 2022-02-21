package cfg

import (
	"github.com/caarlos0/env/v6"
	"log"
)

/*
Добавьте возможность конфигурировать сервис с помощью переменных окружения:

1. адрес запуска HTTP-сервера с помощью переменной - SERVER_ADDRESS.
2. базовый адрес результирующего сокращённого URL с помощью переменной - BASE_URL="http//localhost:41927".
*/

type Config struct {
	SrvAddr string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL string `env:"BASE_URL" envDefault:"http://localhost:8080"`
}

func GetEnvs() Config {
	var data Config
	err := env.Parse(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

var Envs Config
