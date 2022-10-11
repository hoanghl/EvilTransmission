package internal

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Configs struct {
	PORT int      `yaml:"PORT"`
	DB   Database `yaml:"DB"`
}

var Conf Configs

func Initialize() {
	// Load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("No .env found or got error")
	}

	// Load config
	pathConf := "configs.yaml"
	file, err := os.ReadFile(pathConf)
	if err != nil {
		logger.Fatalf("Err as loading 'configs.yaml': %s", err)
		os.Exit(1)
	}
	yaml.Unmarshal(file, &Conf)

	//Initialize components
	InitServer()

}

func ExtractThumbnail(pathVideo string) {
	// Create path for storing thumbnail
	pathWithoutExt := strings.TrimSuffix(pathVideo, filepath.Ext(pathVideo))
	path_thumbnail := fmt.Sprintf("%s_thumb.png", pathWithoutExt)

	// Extract thumbnail from video
	command_args := fmt.Sprintf("-ss 00:00:02 -i %s -frames:v 1 %s", pathVideo, path_thumbnail)
	exec.Command("ffmpeg", strings.Split(command_args, " ")...).Output()
}

func GetFileHash(data []byte) ([]byte, error) {
	hasher := sha256.New()
	hasher.Write(data)

	return hasher.Sum(nil), nil
}
