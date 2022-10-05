package internal

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DBEntry struct {
	resID      string
	path       string
	lastUpdate time.Time
}

type Database struct {
	PORT    int    `yaml:"DB_PORT"`
	USER    string `yaml:"USER"`
	PASS    string `yaml:"PASS"`
	DB_NAME string `yaml:"DB_NAME"`
	TABLE   string `yaml:"TABLE"`
	HOST    string `yaml:"HOST"`
}

func (db *Database) initConn() *sql.DB {
	// Init connection to database
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		db.USER, db.PASS, db.HOST, db.PORT, db.DB_NAME)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Fatalln("No matched entry found")
		} else {
			logger.Fatal(err)
		}
		return nil
	}
	return conn
}

func (db *Database) GetEntry(resid string) *DBEntry {
	// Init connection
	conn := db.initConn()
	if conn == nil {
		return nil
	}
	defer conn.Close()

	// Query
	var resID, path, _lastUpdate string
	row := conn.QueryRow(fmt.Sprintf("SELECT resid, path, lastupdate FROM %s WHERE resid = $1", db.DB_NAME), resid)
	if err := row.Scan(&resID, &path, &_lastUpdate); err != nil {
		logger.Fatal(err)
		return nil
	}
	lastUpdate, err := time.Parse(time.RFC3339, _lastUpdate)
	if err != nil {
		logger.Fatalf("Err as parsing to time: %s", err)
		return nil
	}

	entry := &DBEntry{
		resID:      resID,
		path:       path,
		lastUpdate: lastUpdate,
	}

	return entry
}

func (db *Database) InsertEntry(entry DBEntry) {
	// Init connection
	conn := db.initConn()
	if conn == nil {
		return
	}
	defer conn.Close()

	_, err := conn.Exec(
		fmt.Sprintf("INSERT into %s(resid, path, lastupdate) VALUES ($1, $2, $3)", db.DB_NAME),
		entry.resID, entry.path, entry.lastUpdate.Local().Format(time.RFC3339),
	)
	if err != nil {
		logger.Fatalf("An error occured while executing query: %v", err)
	}

}
