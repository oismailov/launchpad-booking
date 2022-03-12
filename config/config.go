package config

import (
	"encoding/json"
	"log"
	"os"
)

var Cfg = &Config{}

type Config struct {
	DatabaseSettings     DatabaseSettings
	ServerConfigurations ServerConfigurations
}

type DatabaseSettings struct {
	Name     string
	Username string
	Password string
	Host     string
	Port     string
	SSLMode  string
	Timezone string
}

type ServerConfigurations struct {
	Port int
}

func LoadConfig() {
	loadBaseConfig("config/config.json")
}

func LoadTestConfig() {
	loadBaseConfig("../../config/config_test.json")
}

func loadBaseConfig(path string) {
	configFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(configFile)
	configuration := Config{}
	err = decoder.Decode(&configuration)

	if err != nil {
		log.Fatal("error:", err)
	}

	Cfg = &configuration
}
