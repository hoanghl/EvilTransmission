package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/io/server_go/internal"

	"gopkg.in/yaml.v3"
)

type config struct {
	PORT int `yaml:"PORT"`
}

var logger = internal.GetLog()

func (conf *config) loadConfig() *config {
	rootPath := os.Getenv("ROOTPATH")

	configPath := filepath.Join(rootPath, "configs/configs.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logger.Errorf("Path to config not existed: %s", configPath)
	}

	buff, err := os.ReadFile(configPath)
	if err != nil {
		return nil
	}

	err = yaml.Unmarshal(buff, conf)
	if err != nil {
		logger.Er
		return nil
	}
	return conf
}

func main() {

	// Get configs
	conf := config{}
	conf = *conf.loadConfig()

	// Set up GIN
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Routing
	r.GET("/res/:rid", internal.GetRes)

	r.Run(fmt.Sprintf(":%d", conf.PORT))
}
