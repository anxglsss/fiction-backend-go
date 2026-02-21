package models

import "time"

type TournamentStatus string

const (
	TournamentStatusInProgress TournamentStatus = "in_progress"
	TournamentStatusCompleted  TournamentStatus = "completed"
)

type Tournament struct {
	ID        int64            `json:"id" gorm:"primaryKey"`
	UserID    int64            `json:"user_id" gorm:"not null;index"`
	Status    TournamentStatus `json:"status" gorm:"not null;default:'in_progress'"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`

	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
