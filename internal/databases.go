package internal

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Thumbnail struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Row struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	ResType   string     `json:"res_type"`
	Thumbnail *Thumbnail `json:"thumbnail"`
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

func (db *Database) GetMediaInfo() ([]Row, error) {
	// Init connection
	conn := db.initConn()
	if conn == nil {
		return nil, errors.New("cannot establish DB connection")
	}
	defer conn.Close()

	// FIXME: HoangLe [Mar-08]: Replace '15' in below query by another
	// Query
	query := `
	select 
		T1.id				as id
		, T1.res_type		as res_type
		, T1.file_name		as file_name
		, T1.thumbnail_id	as thumbnail_id
		, T2.file_name		as thumbnail_name
	from
		%s T1
	LEFT JOIN %s T2
		on 1=1
		and T2.id = T1.thumbnail_id
	limit 15;
	`
	rows, err := conn.Query(fmt.Sprintf(query, db.TABLE, db.TABLE))
	if err != nil {
		logger.Errorf("Error as querying: %s", err)
		return nil, err
	}

	var ret = []Row{}

	var _res_type, _filename string
	var _id int
	var _thumbnail_id sql.NullInt32
	var _thumbnail_name sql.NullString
	for rows.Next() {
		if err := rows.Scan(&_id, &_res_type, &_filename, &_thumbnail_id, &_thumbnail_name); err != nil {
			logger.Errorf("Error as querying: %s", err)
			continue
		}

		var thumbnail *Thumbnail
		if _thumbnail_id.Valid {
			thumbnail = &Thumbnail{ID: int(_thumbnail_id.Int32), Name: _thumbnail_name.String}
		}

		ret = append(ret, Row{
			ID:        _id,
			Name:      _filename,
			ResType:   _res_type,
			Thumbnail: thumbnail,
		})
	}

	return ret, nil
}

// func (db *Database) GetEntry(entryQuery Row) (Row, error) {
// 	retEntry := Row{}

// 	// Init connection
// 	conn := db.initConn()
// 	if conn == nil {
// 		return retEntry, errors.New("cannot establish DB connection")
// 	}
// 	defer conn.Close()

// 	// Query
// 	var id, path, _lastUpdate string
// 	var row *sql.Row = nil
// 	if entryQuery.ID != "" {
// 		logger.Info("Query with: ID")

// 		query := fmt.Sprintf("SELECT id, path, last_update FROM %s WHERE id = $1", db.TABLE)
// 		row = conn.QueryRow(query, entryQuery.ID)
// 	} else {
// 		logger.Error("Cannot proceed DB query")
// 		return retEntry, errors.New("invalid DB query")
// 	}
// 	if err := row.Scan(&id, &path, &_lastUpdate); err != nil {
// 		logger.Errorf("Error as querying: %s", err)
// 		return Row{}, err
// 	}
// 	lastUpdate, err := time.Parse(time.RFC3339, _lastUpdate)
// 	if err != nil {
// 		logger.Errorf("Err as parsing to time: %s", err)
// 		return Row{}, err
// 	}

// 	return Row{
// 		ID:         _id,
// 		Path:       path,
// 		LastUpdate: lastUpdate,
// 	}, nil

// }

// func (db *Database) InsertEntry(entry Row) error {
// 	// Init connection
// 	conn := db.initConn()
// 	if conn == nil {
// 		return errors.New("cannot establish DB connection")
// 	}
// 	defer conn.Close()

// 	_, err := conn.Exec(
// 		fmt.Sprintf("INSERT into %s(path, last_update) VALUES ($1, $2)", db.TABLE),
// 		entry.Path, entry.LastUpdate.Local().Format(time.RFC3339),
// 	)
// 	return err

// }

// func (db *Database) InitConn() *sql.DB {
// 	// Init connection to database
// 	connStr := fmt.Sprintf(
// 		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
// 		db.USER, db.PASS, db.HOST, db.PORT, db.DB_NAME)
// 	conn, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			logger.Errorln("No matched entry found")
// 		} else {
// 			logger.Error(err)
// 		}
// 		return nil
// 	}
// 	return conn
// }
