package models

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Tournaments []Tournament `json:"tournaments,omitempty" gorm:"foreignKey:UserID"`
}
