package cfg

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	SrvAddr         string `json:"server_address" env:"SERVER_ADDRESS"`       // адрес запуска HTTP-сервера
	BaseURL         string `json:"base_url" env:"BASE_URL"`                   // базовый адрес результирующего сокращённого URL
	FileStoragePath string `json:"file_storage_path" env:"FILE_STORAGE_PATH"` // путь до файла с сокращёнными URL
	CryptoKey       string `env:"CRYPTO_KEY" envDefault:"secret_123456789"`   // secret word to encrypt/decrypt JWT for cookies
	DatabaseDSN     string `json:"database_dsn" env:"DATABASE_DSN"`           // Строка с адресом подключения к БД postgresql://localhost:5432/yandex_practicum_db
	EnableHTTPS     bool   `json:"enable_https" env:"ENABLE_HTTPS"`           // Параметр включения HTTPS у сервера.

	JSONConfig    string `env:"CONFIG"`
	TrustedSubnet string `json:"trusted_subnet" env:"TRUSTED_SUBNET"`
	ServerType    string `json:"server_type" env:"SERVER_TYPE"`
}

const (
	serverAddressDefault   = "localhost:8080"
	baseURLDefault         = "http://localhost:8080"
	fileStoragePathDefault = "tmp/urlPairsData"
	databaseDSNDefault     = "postgresql://localhost:5432/yandex_practicum_db"
	enableHTTPSDefault     = false
	serverTypeDefault      = "http"
)

var (
	Envs    Config
	cfgJSON Config
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
	flag.StringVar(&Envs.TrustedSubnet, "t", "", "строковое представление бесклассовой адресации (CIDR).")

	log.Println("[INFO] cfg -> GetEnvs: parse envs start")
	if err := env.Parse(&Envs); err != nil {
		return err
	}
	configJSON = os.Getenv("CONFIG") // if it's empty or null - we will receive empty string

	log.Println("[INFO] cfg -> GetFlag: flag parse start")
	flag.Parse()

	//json config flag parse
	if configJSON != "" {
		err := json.Unmarshal([]byte(configJSON), &cfgJSON)
		if err != nil {
			return err
		}
	}

	FillEmptyEnvs(&cfgJSON)

	return nil
}

// FillEmptyEnvs verifies that the params are not empty. In case the envs or flags were null - we will use local params.
// Like this, if some values were passed via envs or flags, it will not be overwritten. Only empty fields.
func FillEmptyEnvs(config *Config) {
	if Envs.SrvAddr == "" {
		Envs.SrvAddr = serverAddressDefault
	}
	if Envs.BaseURL == "" {
		Envs.BaseURL = baseURLDefault
	}
	if Envs.FileStoragePath == "" {
		Envs.FileStoragePath = fileStoragePathDefault
	}
	if Envs.DatabaseDSN == "" {
		Envs.DatabaseDSN = databaseDSNDefault
	}
	if config.EnableHTTPS {
		Envs.EnableHTTPS = enableHTTPSDefault
	}
	if config.ServerType == "" {
		Envs.ServerType = serverTypeDefault
	}
}
