package models

import "time"

type Answer struct {
	ID         uint      `gorm:"primaryKey"`
	QuestionID uint      `gorm:"not null"`
	AnswerText string    `gorm:"type:text; not null"`
	IsCorrect  bool      `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	Question   Question  `gorm:"foreignKey:QuestionID"`
}
