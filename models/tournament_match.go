package models

import "time"

type TournamentMatch struct {
	ID              int64     `json:"id" gorm:"primaryKey"`
	TournamentID    int64     `json:"tournament_id" gorm:"not null;index"`
	Round           int       `json:"round" gorm:"not null"`
	SlotInRound     int       `json:"slot_in_round" gorm:"not null"`
	Contestant1Slug string    `json:"contestant1_slug"`
	Contestant2Slug string    `json:"contestant2_slug"`
	WinnerSlug      string    `json:"winner_slug,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
