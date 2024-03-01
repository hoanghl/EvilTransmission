package internal

import (
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DBEntry struct {
	ResID      string
	Path       string
	Hashval    []byte
	LastUpdate time.Time
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
			logger.Errorln("No matched entry found")
		} else {
			logger.Error(err)
		}
		return nil
	}
	return conn
}

var Entry DBEntry

func (db *Database) GetAllEntry() ([]string, error) {
	// Init connection
	conn := db.initConn()
	if conn == nil {
		return nil, errors.New("cannot establish DB connection")
	}
	defer conn.Close()

	// Query

	rows, err := conn.Query(fmt.Sprintf("SELECT resid FROM %s", db.TABLE))
	if err != nil {
		logger.Errorf("Error as querying: %s", err)
		return nil, err
	}

	ret := []string{}
	var s string
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			logger.Errorf("Error as querying: %s", err)
			break
		}

		ret = append(ret, s)
	}

	return ret, err
}

func (db *Database) GetEntry(entryQuery DBEntry) (DBEntry, error) {
	retEntry := DBEntry{}

	// Init connection
	conn := db.initConn()
	if conn == nil {
		return retEntry, errors.New("cannot establish DB connection")
	}
	defer conn.Close()

	// Query
	var resID, path, _lastUpdate string
	var hashVal []byte
	var row *sql.Row = nil
	if entryQuery.ResID != "" {
		logger.Info("Query with: ResID")

		query := fmt.Sprintf("SELECT resid, path, lastupdate, hashval FROM %s WHERE resid = $1", db.TABLE)
		row = conn.QueryRow(query, entryQuery.ResID)
	} else if len(entryQuery.Hashval) != 0 {
		logger.Info("Query with: Hashval")

		query := fmt.Sprintf("SELECT resid, path, lastupdate, hashval FROM %s WHERE hashval = $1", db.TABLE)
		row = conn.QueryRow(query, hex.EncodeToString(entryQuery.Hashval))
	} else {
		logger.Error("Cannot proceed DB query")
		return retEntry, errors.New("invalid DB query")
	}
	if err := row.Scan(&resID, &path, &_lastUpdate, &hashVal); err != nil {
		logger.Errorf("Error as querying: %s", err)
		return DBEntry{}, err
	}
	lastUpdate, err := time.Parse(time.RFC3339, _lastUpdate)
	if err != nil {
		logger.Errorf("Err as parsing to time: %s", err)
		return DBEntry{}, err
	}

	return DBEntry{
		ResID:      resID,
		Path:       path,
		LastUpdate: lastUpdate,
		Hashval:    hashVal,
	}, nil

}

func (db *Database) InsertEntry(entry DBEntry) error {
	// Init connection
	conn := db.initConn()
	if conn == nil {
		return errors.New("cannot establish DB connection")
	}
	defer conn.Close()

	_, err := conn.Exec(
		fmt.Sprintf("INSERT into %s(resid, path, lastupdate, hashval) VALUES ($1, $2, $3, $4)", db.TABLE),
		entry.ResID, entry.Path, entry.LastUpdate.Local().Format(time.RFC3339), hex.EncodeToString(entry.Hashval),
	)
	return err

}

func (db *Database) InitConn() *sql.DB {
	// Init connection to database
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		db.USER, db.PASS, db.HOST, db.PORT, db.DB_NAME)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Errorln("No matched entry found")
		} else {
			logger.Error(err)
		}
		return nil
	}
	return conn
}
