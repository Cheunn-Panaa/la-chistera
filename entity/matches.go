package entity

import (
	"time"

	"gorm.io/gorm"
)

//Construct your model under entity.
type Matches struct {
	gorm.Model
	Date      time.Time `json:"date"`
	Away      int       `gorm:"primaryKey"`
	Home      int       `gorm:"primaryKey"`
	AwayScore int       `json:"away_score"`
	HomeScore int       `json:"home_score"`

	AwayQuotation int
	HomeQuotation int
	DrawQuotation int

	GroupID       uint
	CompetitionID int

	Bets []*MatchBets `gorm:"many2many:bets;"`
}
