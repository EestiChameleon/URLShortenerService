package storage

import (
	"context"
	"encoding/json"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/jackc/pgx/v4/pgxpool"
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
	DB       *pgxpool.Pool
}

func NewSeesion() *Session {
	log.Println("database NewSessions: start")
	return &Session{
		ID:       "",
		Pairs:    map[string]string{},
		UserData: map[string][]Pair{},
	}
}

// InitStorage method parse data from file and initiate all storage dependencies
func (s *Session) InitStorage() error {
	log.Println("database InitStorage: start")
	if cfg.Envs.FileStoragePath == "" {
		if err := cfg.GetEnvs(); err != nil {
			log.Println(err)
			return err
		}
	}

	// create/open file
	log.Println("database InitStorage: openfile")
	f, err := os.OpenFile(cfg.Envs.FileStoragePath, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	log.Println("database InitStorage: read file")
	bytes, err := os.ReadFile(cfg.Envs.FileStoragePath)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("database InitStorage: check len(bytes)")
	if len(bytes) != 0 {
		if err = json.Unmarshal(bytes, &s.Pairs); err != nil {
			log.Println(err)
			return err
		}
	}

	log.Println("database InitStorage: connect to DB")
	if err = s.ConnectToDB(); err != nil {
		log.Println(err)
		return err
	}

	log.Println("database InitStorage: end")
	return nil
}

func (s *Session) CloseStorage() error {
	log.Println("database CloseStorage: start")
	s.DisconnectFromDB()
	return s.UpdateFile()
}

func (s *Session) Put(origURL string) (shortURL string, err error) {
	log.Println("database Put: start")
	log.Println("database Put: origURL", origURL)
	shortURL, err = ShortURL()
	if err != nil {
		log.Println(err)
		return ``, err
	}

	log.Println("database Put: check for already existing shortURL")
	_, ok := s.Pairs[shortURL]
	if !ok {
		// save data
		log.Printf("database Put: save to Pairs. ShortURL: %s, OrigURL: %s\n", shortURL, origURL)
		s.Pairs[shortURL] = origURL
		log.Printf("database Put: append to UserData. ID: %s, ShortURL: %s, OrigURL: %s\n", s.ID, shortURL, origURL)
		s.UserData[s.ID] = append(s.UserData[s.ID], Pair{
			ShortURL: shortURL,
			OrigURL:  origURL,
		})

		// update file
		log.Println("database Put: UpdateFile")
		if err = s.UpdateFile(); err != nil {
			log.Println(err)
			return ``, err
		}

		log.Println("database Put: end. ShortURL: ", shortURL)
		return shortURL, nil
	} else {
		log.Println("database Put: shortURL already exists in Pairs -> s.Put(origURL)")
		return s.Put(origURL)
	}
}

func (s *Session) UpdateFile() error {
	log.Println("database UpdateFile: start")
	// open & rewrite file
	f, err := os.OpenFile(cfg.Envs.FileStoragePath, os.O_WRONLY, 0777)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	// prepare data
	log.Println("database UpdateFile: json.Marshal(s.Pairs)")
	jsonByte, err := json.Marshal(s.Pairs)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("database UpdateFile: f.Write(jsonByte)")
	_, err = f.Write(jsonByte)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("database UpdateFile: end")
	return nil
}

//-------------------- DATABASE --------------------

// ConnectToDB method initialize connection to the indicated DB
func (s *Session) ConnectToDB() error {
	log.Println("database ConnectToDB: start")
	conn, err := pgxpool.Connect(context.Background(), cfg.Envs.DatabaseDSN)
	if err != nil {
		log.Printf("database ConnectToDB: Unable to connect to database: %v\n", err)
		return err
	}

	s.DB = conn
	log.Println("database ConnectToDB: connected. end")
	return nil
}

// DisconnectFromDB closes all connections in the DB pool
func (s *Session) DisconnectFromDB() {
	log.Println("database DisconnectFromDB: start")
	s.DB.Close()
}

// PingDB executes an empty sql statement against DB pool.
// If the sql returns without error, the database Ping is considered successful, otherwise, the error is returned.
func (s *Session) PingDB() error {
	log.Println("database PingDB: start")
	return s.DB.Ping(context.Background())
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
