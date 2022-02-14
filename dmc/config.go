package dmc

//"os"

// A Configuration is the main config object.
type Configuration struct {
	Discord discordConfig  `yaml:"discord"`
	Spotify spotifyConfig  `yaml:"spotify"`
	Youtube youtubeConfig  `yaml:"youtube"`
	DB      DatabaseConfig `yaml:"database"`
}

type spotifyConfig struct {
	ClientID     string
	ClientSecret string
}

type discordConfig struct {
	Token string
}

type youtubeConfig struct {
}

// A DatabaseConfig holds the configuration to connect to a database
type DatabaseConfig struct {
	Username, Password, Host, Name, Type, SSLMode string
	Port                                          int
}
