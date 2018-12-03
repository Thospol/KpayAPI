package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

//Config to use for Setup Server and Database
type Config struct {
	Server   string
	Database string
}

//Read is Readfile in config.toml It's have to set server and database
func (c *Config) Read() {
	_, err := toml.DecodeFile("config/config.toml", &c)
	if err != nil {
		log.Fatal(err)
	}
}
