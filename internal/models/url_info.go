package models

import "time"

type UrlInfo struct {
	Id          int       `json:"id"`
	Url         string    `json:"url"`
	ShortCode   string    `json:"shortCode"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessCount int       `json:"accessCount"`
}
