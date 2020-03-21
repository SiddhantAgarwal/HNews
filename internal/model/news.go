package model

import "time"

// News : Represents a news object
type News struct {
	Headline string    `json:"headline"`
	Content  string    `json:"content"`
	URL      string    `json:"url"`
	Time     time.Time `json:"time"`
}
