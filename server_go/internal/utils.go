package internal

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func Initialize() {
	// Load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal()
	}

	//Initialize components
	InitServer()
	InitDB()

}

func ExtractThumbnail(pathVideo string) {
	// Create path for storing thumbnail
	pathWithoutExt := strings.TrimSuffix(pathVideo, filepath.Ext(pathVideo))
	path_thumbnail := fmt.Sprintf("%s_thumb.png", pathWithoutExt)

	// Extract thumbnail from video
	command_args := fmt.Sprintf("-ss 00:00:02 -i %s -frames:v 1 %s", pathVideo, path_thumbnail)
	exec.Command("ffmpeg", strings.Split(command_args, " ")...).Output()
}
