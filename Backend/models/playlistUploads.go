package models

import "time"

type PlaylistVideos struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	Items         []struct {
		Kind           string `json:"kind"`
		Etag           string `json:"etag"`
		ID             string `json:"id"`
		ContentDetails struct {
			VideoID          string    `json:"videoId"`
			VideoPublishedAt time.Time `json:"videoPublishedAt"`
		} `json:"contentDetails"`
	} `json:"items"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}
