package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/EestiChameleon/URLShortenerService/internal/app/service/data"
)

var (
	pool               *pgxpool.Pool
	ErrDBOrigURLExists = errors.New("pair with given orig_url already exists")
	ErrShortURLDeleted = errors.New("requested shortURL is deleted")
)

// DBStorage structure stores the UserID from the cookie and the DB connection pool structure.
type DBStorage struct {
	ID string
	DB *pgxpool.Pool
}

// InitDBStorage initiates the DB connection and creates the shorten_pairs table.
// If there is an old shorten_pairs table, it's dropped.
func InitDBStorage() (*DBStorage, error) {
	log.Println("[INFO] db -> InitDBStorage: start")
	conn, err := ConnectToDB()
	if err != nil {
		log.Println("db_storage InitDBStorage: err - ", err)
		return nil, err
	}

	// create table if it doesn't exist. Unique column - original_url
	log.Println("db_storage InitDBStorage: check for table existence - create if it's missing")
	_, err = conn.Exec(context.Background(),
		"DROP TABLE IF EXISTS shorten_pairs; "+
			"CREATE TABLE IF NOT EXISTS shorten_pairs (short_url varchar(255) not null, orig_url varchar(255) not null, user_id varchar(42), deleted_at timestamp); "+
			"create index IF NOT EXISTS shorten_pairs_short_url_index on shorten_pairs (short_url); "+
			"create unique index IF NOT EXISTS shorten_pairs_orig_url_uindex on public.shorten_pairs (orig_url);")
	if err != nil {
		log.Println("db_storage InitDBStorage: err - ", err)
		return nil, err
	}

	log.Println("db_storage InitDBStorage: OK")
	pool = conn
	return &DBStorage{DB: conn}, nil
}

//-------------------- DATABASE QUERIES --------------------

// GetOrigURL method gets from the DB the original URL for the passed short URL.
func (db *DBStorage) GetOrigURL(shortURL string) (origURL string, err error) {
	log.Println("[INFO] db -> GetOrigURL: start")
	var deletedAt sql.NullTime
	if err = db.DB.QueryRow(context.Background(),
		"SELECT orig_url, deleted_at FROM shorten_pairs WHERE short_url=$1;", shortURL).Scan(&origURL, &deletedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ``, nil
		}
		return ``, err
	}

	if deletedAt.Valid {
		return ``, ErrShortURLDeleted
	}

	return origURL, nil
}

// GetShortURL method gets from the DB the short URL for the passed original URL.
func (db *DBStorage) GetShortURL(origURL string) (shortURL string, err error) {
	log.Println("[INFO] db -> GetShortURL: start")
	if err = db.DB.QueryRow(context.Background(),
		"SELECT short_url FROM shorten_pairs WHERE orig_url=$1;", origURL).Scan(&shortURL); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ``, nil
		}
		return ``, err
	}

	return shortURL, nil
}

// SavePair method inserts in the DB new pair "original URL":"short URL" with user ID.
func (db *DBStorage) SavePair(pair Pair) error {
	log.Println("[INFO] db -> SavePair: start")
	tag, err := db.DB.Exec(context.Background(),
		"INSERT INTO shorten_pairs (short_url, orig_url, user_id) "+
			"VALUES ($1, $2, $3) "+
			"ON CONFLICT(orig_url) DO NOTHING;",
		pair.ShortURL, pair.OrigURL, db.ID)

	if err != nil {
		return err
	}
	// when no rows were affected, then we have a conflict -> orig_url already exists -> find short_url
	if tag.RowsAffected() == 0 {
		return ErrDBOrigURLExists
	}
	return nil
}

// GetUserURLs method provides the user pairs list from the DB.
func (db *DBStorage) GetUserURLs() (pairs []Pair, err error) {
	log.Println("[INFO] db -> GetUserURLs: start")
	if err = pgxscan.Select(context.Background(), db.DB, &pairs,
		"SELECT short_url, orig_url FROM shorten_pairs WHERE user_id=$1;", db.ID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("[ERROR] db -> GetUserURLs: not found")
			return nil, nil
		}
		log.Println("[ERROR] db -> GetUserURLs:", err)
		return nil, err
	}

	log.Println("[DEBUG] db -> GetUserURLs: OK")
	return pairs, nil
}

//-------------------- DATABASE FUNCTIONS --------------------

// ConnectToDB method initialize connection to the indicated DB.
func ConnectToDB() (*pgxpool.Pool, error) {
	log.Println("[INFO] db -> Shutdown: start")
	conn, err := pgxpool.Connect(context.Background(), cfg.Envs.DatabaseDSN)
	if err != nil {
		log.Printf("database ConnectToDB: Unable to connect to database: %v\n", err)
		return nil, err
	}

	log.Println("[DEBUG] db -> ConnectToDB: OK")
	return conn, nil
}

// Shutdown closes all connections in the DB pool.
func (db *DBStorage) Shutdown() error {
	log.Println("[INFO] db -> Shutdown: start")
	db.DB.Close()
	return nil
}

// SetUserID stores the user ID in the DBStorage structure.
func (db *DBStorage) SetUserID(userID string) {
	log.Println("[INFO] db -> SetUserID: OK")
	db.ID = userID
}

// GetUserID shows the stored user ID in the DBStorage structure.
func (db *DBStorage) GetUserID() string {
	return db.ID
}

// CreateShortURL creates a unique new short URL.
func (db *DBStorage) CreateShortURL() (shortURL string, err error) {
	log.Println("[INFO] db -> GetShortURL: start")
	shortURL, err = data.ShortURL()
	if err != nil {
		log.Println(err)
		return ``, err
	}

	log.Println("[DEBUG] db GetShortURL: check for already existing shortURL")
	origURL, err := db.GetOrigURL(shortURL)
	if err != nil {
		log.Println("[ERROR] db -> GetOrigURL:", err)
	}
	if origURL != "" {
		log.Println("[DEBUG] db -> GetShortURL: shortURL already exists -> try again")
		return db.CreateShortURL()
	}

	log.Println("[DEBUG] db -> GetShortURL: OK")
	return shortURL, nil
}

// PingDB executes an empty sql statement against DB pool.
// If the sql returns without error, the database Ping is considered successful, otherwise, the error is returned.
func PingDB() error {
	log.Println("db -> PingDB: start")
	return pool.Ping(context.Background())
}

// BatchDelete method calls the sql query for multiple delete case.
// As result - single delete query for all passed short URLs.
func (db *DBStorage) BatchDelete(shortURLs []string) error {
	log.Println("[INFO] db -> BatchDelete: start. ShortURLs to delete - ", shortURLs)
	stmnt := MakeBatchUpdateStatement(shortURLs)
	if _, err := db.DB.Exec(context.Background(), stmnt); err != nil {
		return err
	}

	log.Println("[DEBUG] db -> BatchDelete: OK")
	return nil
}

// MakeBatchUpdateStatement creates a sql statement for batch update for BatchDelete method.
// As "deleted" we consider the pair where "deleted_at" is not null.
func MakeBatchUpdateStatement(shortURLs []string) string {
	log.Println("[INFO] db -> MakeBatchUpdateStatement: start")
	strBegin := "UPDATE shorten_pairs SET deleted_at = current_timestamp FROM ( VALUES "
	strEnd := " ) AS myvalues (shortURL) WHERE shorten_pairs.short_url = myvalues.shortURL;"
	var list []string
	for _, v := range shortURLs {
		value := fmt.Sprintf("('%s')", v)
		list = append(list, value)
	}

	values := strings.Join(list, ", ")
	statement := strBegin + values + strEnd

	log.Println("[DEBUG] db -> MakeBatchUpdateStatement: OK")
	return statement
}

// GetStats provides service info about all shorted urls and all users.
func (db *DBStorage) GetStats() (urlsQ int, usrsQ int, err error) {
	log.Println("[INFO] db -> GetStats: start")
	if err = pgxscan.Get(context.Background(), db.DB, &urlsQ,
		"SELECT count(distinct short_url) FROM shorten_pairs;"); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("[ERROR] db -> GetStats urls: not found")
			return 0, 0, nil
		}
		log.Println("[ERROR] db -> GetStats urls:", err)
		return 0, 0, err
	}

	if err = pgxscan.Get(context.Background(), db.DB, &usrsQ,
		"SELECT count(distinct user_id) FROM shorten_pairs;"); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("[ERROR] db -> GetStats users: not found")
			return 0, 0, nil
		}
		log.Println("[ERROR] db -> GetStats users:", err)
		return 0, 0, err
	}

	log.Println("[DEBUG] db -> GetStats: OK")
	return
}
