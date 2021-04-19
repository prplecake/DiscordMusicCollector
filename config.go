package main

import (
//"os"
)

// A Configuration is the main config object.
type Configuration struct {
	Discord discordConfig `yaml:"discord"`
	Spotify spotifyConfig `yaml:"spotify"`
}

type spotifyConfig struct {
	ClientID     string
	ClientSecret string
}

type discordConfig struct {
	Token string
}
