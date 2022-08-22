package internal

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type config struct {
	PORT     int    `yaml:"PORT"`
	PATH_RES string `yaml:"PATH_RES"`
}

var Conf = config{}

func (conf *config) loadConfig() *config {
	pathRoot, found := os.LookupEnv("ROOTPATH")
	if !found {
		logger.Error("Env not found: ROOTPATH")
		os.Exit(1)
	}

	configPath := filepath.Join(pathRoot, "configs/configs.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logger.Errorf("Path to config not existed: %s", configPath)
	}

	buff, err := os.ReadFile(configPath)
	if err != nil {
		return nil
	}

	err = yaml.Unmarshal(buff, conf)
	if err != nil {
		logger.Errorf("Parse config file error: ", err)
		return nil
	}

	return conf
}

func Initialize() {
	// Get configs
	Conf = *Conf.loadConfig()

	//Initialize components
	InitServer()
	InitDB()

}
