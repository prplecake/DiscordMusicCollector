package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify"
	"gopkg.in/yaml.v2"

	"github.com/prplecake/DiscordMusicCollector/dmc"
	"github.com/prplecake/DiscordMusicCollector/dmc/db"
	"github.com/prplecake/DiscordMusicCollector/services"
)

// Global variables
var (
	Token  string
	Config dmc.Configuration
	spc    spotify.Client
	store  *db.Store
)

func init() {
	log.Print("Initializing...")

	// Read configuration
	configFile := "config.yaml"
	cf, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panic(err)
	}
	err = yaml.Unmarshal(cf, &Config)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(Config)

	// Open database, or connection to one
	store, err = db.NewStore(Config.DB)
	if err != nil {
		log.Panic("error opening database: ", err)
	}

	services.AuthenticateSpotify(Config.Spotify.ClientID, Config.Spotify.ClientSecret)
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Config.Discord.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and being listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until C-c or other term signal is received.
	fmt.Println("Bot is now running. Press C-c to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example, but it's good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	for _, field := range strings.Fields(m.Content) {
		if strings.Contains(field, "music.apple.com") {
			err := services.HandleAppleMusicURL(field)
			if err != nil {
				log.Print("services.HandleAppleMusicURL() error: ", err)
			}
		}
		if strings.Contains(field, "spotify.com") {
			track, err := services.HandleSpotifyURL(field)
			if err != nil {
				log.Print("services.HandleSpotifyURL() error: ", err)
			}
			err = db.AddTrackToDB(store, track)
			if err != nil {
				log.Print("db.AddTracktoDB() error: ", err)
			}
		}
		if strings.Contains(field, "youtube.com") {
			err := services.HandleYoutubeURL(field)
			if err != nil {
				log.Print("services.HandleYoutubeURL() error: ", err)
			}
		}
	}
}
