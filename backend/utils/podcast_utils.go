package utils

import (
	"log"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/mmcdole/gofeed"
)

var Converter = md.NewConverter("", true, nil)

type PodcastStats struct {
	PodcastID uint
	IsPlayed  bool
	Count     int
}

type PodcastStatsKey struct {
	PodcastID uint
	IsPlayed  bool
}

func GetSafeFileURL(item *gofeed.Item) string {
	if item.Enclosures == nil || len(item.Enclosures) == 0 {
		return ""
	}
	return item.Enclosures[0].URL
}

func GetSafeImageURL(item *gofeed.Item) string {
	if item.Image == nil {
		return ""
	}
	return item.Image.URL
}

func ConvertHtmlToMarkdown(toConvert string) string {
	markdown, err := Converter.ConvertString(toConvert)
	if err != nil {
		log.Println(err)
		return ""
	}
	return markdown
}
