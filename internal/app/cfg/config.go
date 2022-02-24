package cfg

import (
	"github.com/caarlos0/env/v6"
	"log"
)

/*
Задание для трека «Сервис сокращения URL»
Поддержите конфигурирование сервиса с помощью флагов командной строки наравне с уже имеющимися переменными окружения:
флаг -a, отвечающий за адрес запуска HTTP-сервера (переменная SERVER_ADDRESS);
флаг -b, отвечающий за базовый адрес результирующего сокращённого URL (переменная BASE_URL);
флаг -f, отвечающий за путь до файла с сокращёнными URL (переменная FILE_STORAGE_PATH).
*/

type Config struct {
	SrvAddr         string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`      //флаг -a, отвечающий за адрес запуска HTTP-сервера (переменная SERVER_ADDRESS);
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`     //флаг -b, отвечающий за базовый адрес результирующего сокращённого URL (переменная BASE_URL);
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"/tmp/defaultFile"` //флаг -f, отвечающий за путь до файла с сокращёнными URL (переменная FILE_STORAGE_PATH).
}

var Envs Config

func GetEnvs() {
	err := env.Parse(&Envs)
	if err != nil {
		log.Fatal(err)
	}
}
