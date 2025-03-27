package models

import "time"

type SheetStats struct {
	ID               uint          `gorm:"primaryKey"`
	SheetID          uint          `gorm:"not null"`
	UserID           uint          `gorm:"not null"`
	TotalQuestions   int           `gorm:"not null"`
	CorrectAnswers   int           `gorm:"not null"`
	TimeTakenSeconds int           `gorm:"not null"`
	CompletedAt      time.Time     `gorm:"not null"`
	CreatedAt        time.Time     `gorm:"autoCreateTime"`
	UpdatedAt        time.Time     `gorm:"autoUpdateTime"`
	QuestionSheet    QuestionSheet `gorm:"foreignKey:SheetID"`
	User             User          `gorm:"foreignKey:UserID"`
}
