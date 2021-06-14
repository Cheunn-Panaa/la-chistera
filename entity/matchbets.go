package entity

import (
	"gorm.io/gorm"
)

//Construct your model under entity.
type MatchBets struct {
	gorm.Model

	MatchID int `gorm:"primaryKey"`
	UserID  int `gorm:"primaryKey"`

	AwayScore int `json:"away_score"`
	HomeScore int `json:"home_score"`
}
