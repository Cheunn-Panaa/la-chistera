package entity

import (
	"gorm.io/gorm"
)

//Construct your model under entity.
type Competition struct {
	gorm.Model
	Title     string `json:"title"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Groups    []Group
	Teams     []Team `gorm:"many2many:standings"`
	Matches   []Matches
}

//NewCompetition create a new competiton
func NewCompetition(title string, startDate string, endDate string) (*Competition, error) {
	c := &Competition{
		Title:     title,
		StartDate: startDate,
		EndDate:   endDate,
	}
	return c, nil
}
