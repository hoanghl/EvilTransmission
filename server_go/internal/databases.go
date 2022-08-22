package internal

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

var db map[string]interface{}

func InitDB() {
	pathRoot, _ := os.LookupEnv("ROOTPATH")

	pathDb := filepath.Join(pathRoot, Conf.PATH_RES)
	if _, err := os.Stat(pathDb); os.IsNotExist(err) {
		logger.Errorf("Path to res not existed: %s", pathDb)
	}

	// Read JSON
	fileJSON, err := os.Open(pathDb)
	if err != nil {
		logger.Errorf("Read DB file error: %s", pathDb)
		os.Exit(1)
	}
	defer fileJSON.Close()
	bytesJSON, _ := io.ReadAll(fileJSON)

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(bytesJSON), &db)
}

// TODO: HoangLe [Aug-21]: Enable later
// func GetEntry(entry string) (string, bool) {
// 	if val, existed := db[entry]; existed {
// 		return "got", existed
// 	} else {
// 		return "", existed
// 	}
// }
