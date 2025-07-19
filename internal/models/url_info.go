package models

import "time"

type UrlInfo struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Url         string    `json:"url" binding:"required"`
	ShortCode   string    `json:"shortCode"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessCount int       `json:"accessCount"`
}

const Schema = `
CREATE TABLE url_info( 
    id VARCHAR(255),
    url VARCHAR(2000),
    short_code VARCHAR(255),
    access_count VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY(id)
);
`
