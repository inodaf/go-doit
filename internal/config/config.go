package config

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func GetDatabasePath() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appConfigDir := filepath.Join(userConfigDir, "godoit")
	if _, err := os.Stat(appConfigDir); err != nil {
		if err := os.Mkdir(appConfigDir, os.ModePerm); err != nil {
			log.Fatalln("GetDB: ", err)
		}
	}

	return filepath.Join(appConfigDir, "data.db"), nil
}

const TempFileName = "tmp.md"
const DisplayTimeFormat = time.RFC822
