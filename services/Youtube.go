package services

import (
	"log"
	"regexp"
)

const (
	youtubeURLRegex = `(?:https?:)?(?:\/\/)?(?:www\.)?(?:youtu\.be\/|youtube(?:\-nocookie)?\.(?:[A-Za-z]{2,4}|[A-Za-z]{2,3}\.[A-Za-z]{2})\/)(?:watch|embed\/|vi?\/)*(?:\?[\w=&]*vi?=)?([^#&\?\/]{11}).*`
)

func extractYoutubeVideoID(message string) []string {
	re := regexp.MustCompile(youtubeURLRegex)
	return re.FindAllString(message, -1)
}

// HandleYoutubeResult uses the Youtube API to gather information about
// a video.
func HandleYoutubeURL(url string) error {
	log.Print("Handling Youtube URL...")
	return nil
}
