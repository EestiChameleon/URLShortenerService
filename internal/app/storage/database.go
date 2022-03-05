package storage

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"os"
	"path/filepath"
)

var (
	Pairs = NewFile()
)

type data struct {
	File     *os.File
	FileName string
	FileDir  string
	FileData map[string]string
}

func NewFile() *data {
	if cfg.Envs.FileStoragePath == "" {
		cfg.GetEnvs() // ?!
	}

	dir, name := filepath.Split(cfg.Envs.FileStoragePath)

	return &data{
		FileName: name,
		FileDir:  dir,
		FileData: map[string]string{},
	}
}

//TestNewFile provides test data
func TestNewFile() *data {
	cfg.Envs.FileStoragePath = "test/fileTest"
	dir, name := filepath.Split(cfg.Envs.FileStoragePath)

	return &data{
		FileName: name,
		FileDir:  dir,
		FileData: map[string]string{"http://localhost:8080/test": "https://jwt.io/"},
	}
}

func (d *data) Get(key string) string {
	return d.FileData[key]
}

func (d *data) Put(value string) (key string, err error) {
	key, err = ShortURL()
	if err != nil {
		return "", err
	}
	_, ok := d.Check(key)
	if !ok {
		err = d.SaveData(key, value)
		if err != nil {
			return "", err
		}
		return key, nil
	} else {
		return d.Put(value)
	}
}

func (d *data) Check(key string) (value string, ok bool) {
	value, ok = d.FileData[key]
	if ok {
		return value, true
	} else {
		return "", false
	}
}

func ShortURL() (shortedURL string, err error) {
	// 7 bytes is enough to provide more than 78kkk diff combinations
	b := make([]byte, 7)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	shortedURL = fmt.Sprintf("%s/%x", cfg.Envs.BaseURL, b[0:])
	return
}

// Get stored pairs from file ----------------------------------------

func (d *data) GetFile() error {
	//create dir for storage file, If directory already exists, CreateDir does nothing and returns nil
	err := d.CreateDir()
	if err != nil {
		return err
	}

	// create/open file
	file, err := os.OpenFile(cfg.Envs.FileStoragePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	d.File = file

	bytes, err := os.ReadFile(cfg.Envs.FileStoragePath)
	if err != nil {
		return err
	}

	if len(bytes) != 0 {
		err = json.Unmarshal(bytes, &d.FileData)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *data) SaveData(shortURL string, longURL string) error {
	// save to map
	d.FileData[shortURL] = longURL
	// rewrite the file
	jsonByte, err := json.Marshal(d.FileData)
	if err != nil {
		return err
	}

	_, err = d.File.Write(jsonByte)
	if err != nil {
		return err
	}
	return nil
}

func (d *data) CloseFile() error {
	return d.File.Close()
}

func (d *data) CreateDir() error {
	absPath, err := filepath.Abs("")
	if err != nil {
		return err
	}

	fileDir := filepath.Join(absPath, d.FileDir)
	err = os.MkdirAll(fileDir, 0777)
	if err != nil {
		return err
	}

	return nil
}
