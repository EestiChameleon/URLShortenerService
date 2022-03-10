package storage

import (
	"encoding/json"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"log"
	"os"
)

var (
	User = NewSeesion()
)

type Pair struct {
	ShortURL string `json:"short_url"`
	OrigURL  string `json:"orig_url"`
}

type Session struct {
	ID       string
	Pairs    map[string]string
	UserData map[string][]Pair
}

func NewSeesion() *Session {
	return &Session{
		ID:       "",
		Pairs:    map[string]string{},
		UserData: map[string][]Pair{},
	}
}

// InitStorage method parse data from file and initiate all storage dependencies
func (s *Session) InitStorage() error {
	if cfg.Envs.FileStoragePath == "" {
		if err := cfg.GetEnvs(); err != nil {
			log.Println(err)
			return err
		}
	}

	//create dir for storage file, If directory already exists, CreateDir does nothing and returns nil
	//err := s.CreateDir()
	//if err != nil {
	//	return err
	//}

	// create/open file
	f, err := os.OpenFile(cfg.Envs.FileStoragePath, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	bytes, err := os.ReadFile(cfg.Envs.FileStoragePath)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(bytes) != 0 {
		if err = json.Unmarshal(bytes, &s.Pairs); err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (s *Session) CloseStorage() error {
	return s.UpdateFile()
}

func (s *Session) Put(origURL string) (shortURL string, err error) {
	shortURL, err = ShortURL()
	if err != nil {
		log.Println(err)
		return ``, err
	}
	_, ok := s.Pairs[shortURL]
	if !ok {
		// save data
		s.Pairs[shortURL] = origURL
		s.UserData[s.ID] = append(s.UserData[s.ID], Pair{
			ShortURL: shortURL,
			OrigURL:  origURL,
		})

		// update file
		if err = s.UpdateFile(); err != nil {
			log.Println(err)
			return ``, err
		}

		return shortURL, nil
	} else {
		return s.Put(origURL)
	}
}

func (s *Session) UpdateFile() error {
	// open & rewrite file
	f, err := os.OpenFile(cfg.Envs.FileStoragePath, os.O_WRONLY, 0777)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	// prepare data
	jsonByte, err := json.Marshal(s.Pairs)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = f.Write(jsonByte)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//-------------------- TEST DATA --------------------

//TestUser provides test session
func TestUser() *Session {
	return &Session{
		ID:    "test",
		Pairs: map[string]string{"http://localhost:8080/test": "https://jwt.io/"},
		UserData: map[string][]Pair{"test": {Pair{
			ShortURL: "http://localhost:8080/test",
			OrigURL:  "https://jwt.io/",
		}}},
	}
}

// InitTestStorage method prepares test data
func (s *Session) InitTestStorage() error {
	cfg.Envs.FileStoragePath = "testFile"

	// init storage struct
	User = TestUser()
	//create dir for storage file, If directory already exists, CreateDir does nothing and returns nil
	//err := s.CreateDir()
	//if err != nil {
	//	return err
	//}

	// create/open file
	f, err := os.OpenFile("testFile", os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	bytes, err := os.ReadFile("testFile")
	if err != nil {
		log.Println(err)
		return err
	}

	if len(bytes) != 0 {
		err = json.Unmarshal(bytes, &s.Pairs)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}
