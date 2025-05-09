package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig(path string) (Config, error) {
	var cfg Config
	cfgFile, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	log.Println("Successfully read config.")

	return cfg, yaml.Unmarshal(cfgFile, &cfg)
}

func MustReadConfig(path string) Config {
	cfg, err := ReadConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
