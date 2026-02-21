package models

import "time"

type MediaType string

const (
	MediaTypeAnime  MediaType = "anime"
	MediaTypeMovie  MediaType = "movie"
	MediaTypeSeries MediaType = "series"
)

type Media struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Type        MediaType `json:"type" gorm:"type:varchar(20);not null"`
	ImageURL    string    `json:"image_url"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
