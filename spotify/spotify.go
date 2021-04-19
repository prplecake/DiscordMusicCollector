package spotify

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify"
)

// Authenticate provides a wrapper for the zmb3/spotify library
func Authenticate(clientID, secretKey string) spotify.Client {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: secretKey,
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)
	return client
}
