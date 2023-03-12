package dto

type PodcastStats struct {
	PodcastID     uint
	LatestEpisode string
}

type PodcastStatsDetail struct {
	PodcastStats
	Counters map[PodcastStatsKey]PodcastPlayedStatus
}

type PodcastPlayedStatus struct {
	PodcastID uint
	IsPlayed  bool
	Count     int
}

type PodcastStatsKey struct {
	PodcastID uint
	IsPlayed  bool
}
