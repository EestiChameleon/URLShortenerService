package cfg

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	SrvAddr         string `env:"SERVER_ADDRESS"`                           // адрес запуска HTTP-сервера
	BaseURL         string `env:"BASE_URL"`                                 // базовый адрес результирующего сокращённого URL
	FileStoragePath string `env:"FILE_STORAGE_PATH"`                        // путь до файла с сокращёнными URL
	CryptoKey       string `env:"CRYPTO_KEY" envDefault:"secret_123456789"` // secret word to encrypt/decrypt JWT for cookies
	DatabaseDSN     string `env:"DATABASE_DSN"`                             // Строка с адресом подключения к БД postgresql://localhost:5432/yandex_practicum_db
	EnableHTTPS     bool   `env:"ENABLE_HTTPS"`                             // Параметр включения HTTPS у сервера.
}

var (
	Envs        Config
	cfgJSON     *Config
	localConfig string = `{
	"server_address": "localhost:8080",
	"base_url": "http://localhost:8080",
	"file_storage_path": "tmp/urlPairsData",
	"database_dsn": "postgresql://localhost:5432/yandex_practicum_db",
	"enable_https": false
	}`
)

type ContextKey string

func GetEnvs() error {
	log.Println("[INFO] cfg -> GetFlag: flag declaration start")

	var configJSON string

	flag.StringVar(&Envs.SrvAddr, "a", "", "SERVER_ADDRESS to listen on")
	flag.StringVar(&Envs.BaseURL, "b", "", "BASE_URL of the shorten result URL")
	flag.StringVar(&Envs.FileStoragePath, "f", "", "FILE_STORAGE_PATH. Directory of the origin&shorten url pairs file")
	flag.StringVar(&Envs.DatabaseDSN, "d", "", "DATABASE_DSN. Address for connection to DB")
	flag.BoolVar(&Envs.EnableHTTPS, "s", false, "ENABLE_HTTPS parameter. Enable the HTTPS server.")

	flag.StringVar(&configJSON, "c", "", "APP config via JSON file. All params included")

	log.Println("[INFO] cfg -> GetEnvs: parse envs start")
	if err := env.Parse(&Envs); err != nil {
		return err
	}
	configJSON = os.Getenv("CONFIG") // if it's empty or null - we will receive empty string

	log.Println("[INFO] cfg -> GetFlag: flag parse start")
	flag.Parse()

	//json config flag parse
	if configJSON != "" {
		err := json.Unmarshal([]byte(configJSON), cfgJSON)
		if err != nil {
			return err
		}
	} else {
		err := json.Unmarshal([]byte(localConfig), cfgJSON)
		if err != nil {
			return err
		}
	}

	FillEmptyEnvs(cfgJSON)

	return nil
}

// FillEmptyEnvs verifies that the params are not empty. In case the envs or flags were null - we will use local params.
// Like this, if some values were passed via envs or flags, it will not be overwritten. Only empty fields.
func FillEmptyEnvs(config *Config) {
	if Envs.SrvAddr == "" {
		Envs.SrvAddr = config.SrvAddr
	}
	if Envs.BaseURL == "" {
		Envs.BaseURL = config.BaseURL
	}
	if Envs.FileStoragePath == "" {
		Envs.FileStoragePath = config.FileStoragePath
	}
	if Envs.DatabaseDSN == "" {
		Envs.DatabaseDSN = config.DatabaseDSN
	}
	if config.EnableHTTPS {
		Envs.EnableHTTPS = config.EnableHTTPS
	}
}
