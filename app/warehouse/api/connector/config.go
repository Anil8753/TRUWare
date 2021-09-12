package connector

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	WalletPath            string
	ConnectionProfilePath string
	User                  string
}

func NewConfig() *Config {

	configPath := "./config.json"
	data, err := os.ReadFile(configPath)
	check(err)

	cfg := Config{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	log.Printf("Connection config file \n%s\n", string(data))
	return &cfg
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
