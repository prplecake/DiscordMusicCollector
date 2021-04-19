package services

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify"
)

var (
	Client spotify.Client
)

// Authenticate provides a wrapper for the zmb3/spotify library
func AuthenticateSpotify(clientID, secretKey string) spotify.Client {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: secretKey,
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	return spotify.Authenticator{}.NewClient(token)
}
