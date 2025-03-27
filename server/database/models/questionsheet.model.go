package models

import "time"

type QuestionSheet struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	Title       string    `gorm:"size:100;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Questions   []Question
	User        User `gorm:"foreignKey:UserID"`
}
