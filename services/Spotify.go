package services

import (
	"context"
	"log"
	"regexp"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify"

	"github.com/prplecake/DiscordMusicCollector/app"
)

const (
	spotifyURLRegex = `https:\/\/open.spotify.com\/track\/(?P<trackId>[a-zA-Z0-9]+)\?si=(?P<si>[a-zA-z0-9\-]+)`
)

var (
	client spotify.Client
)

// AuthenticateSpotify provides a wrapper for the zmb3/spotify library
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

	client = spotify.Authenticator{}.NewClient(token)
}

func extractSpotifyTrackID(url string) spotify.ID {
	re := regexp.MustCompile(spotifyURLRegex)
	result := app.GetParams(re, url)
	return spotify.ID(result["trackId"])
}

// HandleSpotifyURL uses the Spotify API to gather information on a
// track.
func HandleSpotifyURL(url string) (app.Track, error) {
	log.Print("Handling Spotify URL...")

	sr, err := client.GetTrack(extractSpotifyTrackID(url))
	if err != nil {
		return app.Track{}, err
	}
	artists := []string{}
	for _, artist := range sr.SimpleTrack.Artists {
		artists = append(artists, artist.Name)
	}
	track := app.Track{
		Title:   sr.SimpleTrack.Name,
		Artists: artists,
		Album:   sr.Album.Name,
		Service: "spotify",
	}
	log.Printf("Got track:\n\tTitle:\t\t%s\n\tArtist(s):\t%s\n\tAlbum:\t\t%s",
		track.Title, track.Artists, track.Album)
	return track, nil
}
