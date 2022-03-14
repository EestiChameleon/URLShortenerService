package storage

import (
	"context"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/EestiChameleon/URLShortenerService/internal/app/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var pool *pgxpool.Pool

type DBStorage struct {
	ID string
	DB *pgxpool.Pool
}

func InitDBStorage() (*DBStorage, error) {
	log.Println("db_storage InitDBStorage: connect to DB")
	conn, err := ConnectToDB()
	if err != nil {
		log.Println("db_storage InitDBStorage: err - ", err)
		return nil, err
	}

	// create table if it doesn't exist
	log.Println("db_storage InitDBStorage: check for table existence - create if it's missing")
	_, err = conn.Exec(context.Background(),
		"CREATE TABLE IF NOT EXISTS shorten_pairs (short_url varchar(255) not null, orig_url varchar(255) not null, user_id   varchar(42)); create index IF NOT EXISTS shorten_pairs_short_url_index on shorten_pairs (short_url);")
	if err != nil {
		log.Println("db_storage InitDBStorage: err - ", err)
		return nil, err
	}

	log.Println("db_storage InitDBStorage: OK")
	pool = conn
	return &DBStorage{DB: conn}, nil
}

//-------------------- DATABASE QUERIES --------------------

func (db *DBStorage) GetURL(shortURL string) (origURL string, err error) {
	if err = db.DB.QueryRow(context.Background(),
		"SELECT orig_url FROM shorten_pairs WHERE short_url=$1;", shortURL).Scan(&origURL); err != nil {
		return ``, err
	}

	return origURL, nil
}

func (db *DBStorage) SavePair(pair Pair) (err error) {
	_, err = db.DB.Exec(context.Background(),
		"INSERT INTO shorten_pairs(short_url, orig_url, user_id) VALUES($1, $2, $3);",
		pair.ShortURL, pair.OrigURL, db.ID)

	return err
}

func (db *DBStorage) GetUserURLs() (pairs []Pair, err error) {
	if err = pgxscan.Select(context.Background(), db.DB, &pairs,
		"SELECT short_url, orig_url FROM shorten_pairs WHERE user_id=$1;", db.ID); err != nil {
		return nil, err
	}

	return pairs, nil
}

//-------------------- DATABASE FUNCTIONS --------------------

// ConnectToDB method initialize connection to the indicated DB
func ConnectToDB() (*pgxpool.Pool, error) {
	log.Println("db_storage ConnectToDB: start")
	conn, err := pgxpool.Connect(context.Background(), cfg.Envs.DatabaseDSN)
	if err != nil {
		log.Printf("database ConnectToDB: Unable to connect to database: %v\n", err)
		return nil, err
	}

	log.Println("db_storage ConnectToDB: connected. end")
	return conn, nil
}

// ShutDown closes all connections in the DB pool
func (db *DBStorage) ShutDown() error {
	log.Println("db_storage ShutDown: start")
	db.DB.Close()
	return nil
}

func (db *DBStorage) SetUserID(userID string) {
	db.ID = userID
}

func (db *DBStorage) CreateShortURL() (shortURL string, err error) {
	log.Println("db_storage GetShortURL: start")
	shortURL, err = models.ShortURL()
	if err != nil {
		log.Println(err)
		return ``, err
	}

	log.Println("db_storage GetShortURL: check for already existing shortURL")
	origURL, err := db.GetURL(shortURL)
	if err != nil {
		log.Println(err)
	}
	if origURL != "" {
		log.Println("db_storage GetShortURL: shortURL already exists -> try again")
		return db.CreateShortURL()
	}

	log.Println("db_storage GetShortURL: OK")
	return shortURL, nil
}

// PingDB executes an empty sql statement against DB pool.
// If the sql returns without error, the database Ping is considered successful, otherwise, the error is returned.
func PingDB() error {
	log.Println("db_storage PingDB: start")
	return pool.Ping(context.Background())
}
