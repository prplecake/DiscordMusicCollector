package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
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
	// If the message is "ping" reply with "Pong!"
	if strings.Contains(m.Content, "youtube.com") {
		fmt.Println("got youtube link")
		links := extractYtLinks(m.Content)
		fmt.Println(links)
	}
	if strings.Contains(m.Content, "spotify.com") {
		fmt.Println("got spotify link")
		links := extractSpotifyLinks(m.Content)
		fmt.Println(links)
	}
	if strings.Contains(m.Content, "music.apple.com") {
		fmt.Println("got apple music link")
		links := extractAppleMusicLinks(m.Content)
		fmt.Println(links)
	}
}

func extractYtLinks(message string) []string {
	re := regexp.MustCompile(`(?:https?:)?(?:\/\/)?(?:www\.)?(?:youtu\.be\/|youtube(?:\-nocookie)?\.(?:[A-Za-z]{2,4}|[A-Za-z]{2,3}\.[A-Za-z]{2})\/)(?:watch|embed\/|vi?\/)*(?:\?[\w=&]*vi?=)?([^#&\?\/]{11}).*`)
	return re.FindAllString(message, -1)
}

func extractSpotifyLinks(message string) []string {
	re := regexp.MustCompile(`(https:\/\/open.spotify.com\/track\/([a-zA-Z0-9]+)|spotify:user:([a-zA-Z0-9]+):playlist:)([a-zA-Z0-9]+)\?si=([a-zA-z0-9]+)`)
	return re.FindAllString(message, -1)
}

func extractAppleMusicLinks(message string) []string {
	re := regexp.MustCompile(`(https:\/\/music.apple.com\/us\/album\/([a-zA-Z0-9\-]+)\/([0-9]+)\?i=+([0-9]+))`)
	return re.FindAllString(message, -1)
}

func addLinkToDB(link string, service string) {}
