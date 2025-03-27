package models

import "time"

type Question struct {
	ID            uint          `gorm:"primaryKey"`
	SheetID       uint          `gorm:"not nulL"`
	UserID        uint          `gorm:"not null"`
	QuestionText  string        `gorm:"type:text;not null"`
	CreatedAt     time.Time     `gorm:"autoCreateTime"`
	UpdatedAt     time.Time     `gorm:"autoUpdateTime"`
	QuestionSheet QuestionSheet `gorm:"foreignKey:SheetID"`
	User          User          `gorm:"foreignKey:UserID"`
	Answers       []Answer
}
