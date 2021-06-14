package entity

import "gorm.io/gorm"

//Construct your model under entity.
type Group struct {
	gorm.Model
	Title string `json:"title"`

	CompetitionID int

	TeamID  int       `gorm:"primaryKey"`
	Matches []Matches `gorm:"foreignkey:GroupID"`
}
