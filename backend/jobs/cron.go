package jobs

import (
	"example/hello/config"
	"example/hello/db"
	"example/hello/db/models"
	"example/hello/utils"
	"log"

	"github.com/mmcdole/gofeed"
	"github.com/robfig/cron/v3"
)

func InitCron() {
	config := config.GetConfig()
	cron := cron.New()
	cron.AddFunc("@every "+config.GetString("cron.interval"), updatePodcastItemsList)
	cron.Start()
}

func updatePodcastItemsList() {
	var podcastsList []models.Podcast
	db.GetDb().Find(&podcastsList)
	for _, podcast := range podcastsList {
		var guidList []string
		db.GetDb().Table("podcast_items").
			Where("podcast_id = ?", podcast.ID).
			Select("guid").
			Find(&guidList)
		guidMap := make(map[string]int)
		for _, guid := range guidList {
			guidMap[guid] = 1
		}

		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(podcast.FeedURL)
		var podcastItemsArray []models.PodcastItem
		for _, item := range feed.Items {
			if _, guidExist := guidMap[item.GUID]; guidExist {
				continue
			}
			podcastItemsArray = append(podcastItemsArray, models.PodcastItem{
				Title:           item.Title,
				Summary:         utils.ConvertHtmlToMarkdown(item.Description),
				PublicationDate: *item.PublishedParsed,
				FileURL:         utils.GetSafeFileURL(item),
				GUID:            item.GUID,
				Image:           utils.GetSafeImageURL(item),
				PodcastID:       int(podcast.ID),
				UserID:          int(podcast.UserId),
			})
		}
		if len(podcastItemsArray) > 0 {
			result := db.GetDb().Create(&podcastItemsArray)
			log.Println("Podcast_ID: ", podcast.ID, "\tNew records: ", result.RowsAffected)
		}
	}
}
