package models

import "time"

type UrlInfo struct {
	Id          int       `json:"id" gorm:"primaryKey autoIncrement"`
	Url         string    `json:"url" binding:"required"`
	ShortCode   string    `json:"shortCode"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessCount int       `json:"accessCount"`
}

const Schema = `
CREATE TABLE url_info( 
    id serial not null primary key,
    url varchar(2000) not null,
    short_code varchar(255) not null unique,
    access_count int,
    created_at timestamp default now(),
    updated_at timestamp default now()
);
`
