package services

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify"

	"git.sr.ht/~mjorgensen/DiscordMusicCollector/DCM"

)

var (
	Client spotify.Client
)

// Authenticate provides a wrapper for the zmb3/spotify library
func AuthenticateSpotify(clientID, secretKey string) {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: secretKey,
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	Client = spotify.Authenticator{}.NewClient(token)
}

func HandleSpotifyResult(result spotify.ID) error {
	log.Print("Handling Spotify result...")

	sr, err := Client.GetTrack(result)
	if err != nil {
		log.Print("error getting spotify track data: ", err)
	}
	artists := []string{}
	for _, artist := range sr.SimpleTrack.Artists {
		artists = append(artists, artist.Name)
	}
	track := DCM.Track{
		Name: sr.SimpleTrack.Name,
		Artists: artists,
		Album: sr.Album.Name,
	}
	log.Print(track)
	return nil
}