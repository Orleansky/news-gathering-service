package models

import "time"

type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	Link        string    `json:"link"`
}
