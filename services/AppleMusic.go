package services

import (
	"log"
	"regexp"
)

const (
	appleMusicURLRegex = `(https:\/\/music.apple.com\/us\/album\/([a-zA-Z0-9\-]+)\/([0-9]+)\?i=+([0-9]+))`
)

func extractAppleMusicTrackID(message string) []string {
	re := regexp.MustCompile(appleMusicURLRegex)
	return re.FindAllString(message, -1)
}

// HandleAppleMusicResult uses the Apple Music API to gather information
// about a track.
func HandleAppleMusicURL(url string) error {
	log.Print("Handling Apple Music URL...")
	return nil
}
