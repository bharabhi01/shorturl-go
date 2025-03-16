package repository

import "time"

type URL struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ShortCode string    `gorm:"uniqueIndex;no null" json:"short_code"`
	LongURL   string    `gorm:"not null" json:"long_url"`
	UserID    string    `gorm:"index" json:"user_id,omitempty"`
	Visits    int64     `gorm:"default:0" json:"visits"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}
