package config

import (
	"log"
	"os"
	"path"
)

var (
	FssConfigDir = ""
)

func init() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	FssConfigDir = path.Join(configDir, "fss")

	if _, err := os.Stat(FssConfigDir); os.IsNotExist(err) {
		os.MkdirAll(FssConfigDir, 0777)
	}
}
