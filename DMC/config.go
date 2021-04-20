package dmc

import (
//"os"
)

// A Configuration is the main config object.
type Configuration struct {
	Discord discordConfig  `yaml:"discord"`
	Spotify spotifyConfig  `yaml:"spotify"`
	DB      DatabaseConfig `yaml:"database"`
}

type spotifyConfig struct {
	ClientID     string
	ClientSecret string
}

type discordConfig struct {
	Token string
}

// A DatabaseConfig holds the configuration to connect to a database
type DatabaseConfig struct {
	Username, Password, Hostname, Name, Type string
	Port                                     int
}
