package config

import (
	"encoding/json"
	"log"
	"os"
)

var Config struct {
	// The address to listen on
	Listen string `json:"listen"`
	// The admin token
	Token string `json:"token"`
	// The database token
	Database string `json:"database"`
}

const Version = "1.0.0"

// LoadConfig loads the config file from a location
func LoadConfig(location string) {
	bytes, err := os.ReadFile(location)
	if err != nil {
		log.Fatalf("Cannot read config file: %s\n", err)
	}
	err = json.Unmarshal(bytes, &Config)
	if err != nil {
		log.Fatalf("Cannot parse config file: %s\n", err)
	}
}
