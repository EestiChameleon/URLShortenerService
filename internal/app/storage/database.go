package storage

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const ShortLinkHost = "http://localhost:8080"

var (
	Pairs = NewFileData()
)

type data struct {
	File     *os.File
	FileName string
	FileDir  string
	FileData map[string]string
}

func NewFileData() *data {
	if cfg.Envs.FileStoragePath == "" {
		cfg.GetEnvs()
	}

	dir, name := filepath.Split(cfg.Envs.FileStoragePath)

	return &data{
		FileName: name,
		FileDir:  dir,
		FileData: map[string]string{},
	}
}

//TestStore provides test data
func TestNewFileData() *data {
	return &data{
		FileName: "testFile.txt",
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
		err = d.AddDataAndSaveToFile(key, value)
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
	//change dir to fileDir
	err := d.ChangeDir()
	if err != nil {
		return err
	}
	// create/open file
	file, err := os.OpenFile(d.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	sl := ""

	d.File = file

	// make a read buffer
	r := bufio.NewReader(d.File)
	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		sl += string(buf)
	}
	pairs := strings.Split(sl, "\n")
	for _, el := range pairs {
		if !strings.Contains(el, " : ") {
			break
		}
		v := strings.Split(el, " : ")
		d.FileData[v[0]] = v[1]
	}

	return nil
}

func (d *data) WriteFile(s string) error {

	_, err := d.File.WriteString(s + "\n")
	return err
}

func (d *data) AddDataAndSaveToFile(shortURL string, longURL string) error {

	d.FileData[shortURL] = longURL
	err := d.WriteFile(shortURL + " : " + longURL)
	return err
}

func (d *data) CloseFile() error {
	return d.File.Close()
}

func (d *data) ChangeDir() error {
	absPath, err := filepath.Abs("")
	if err != nil {
		return err
	}

	fileDir := filepath.Join(absPath, d.FileDir)
	err = os.MkdirAll(fileDir, 0777)
	if err != nil {
		return err
	}

	err = os.Chdir(fileDir)
	if err != nil {
		return err
	}

	return nil
}
