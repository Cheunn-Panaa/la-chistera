package entity

import (
	"gorm.io/gorm"
)

//Construct your model under entity.
type Standings struct {
	gorm.Model

	TeamID        int `gorm:"primaryKey"`
	CompetitionID int `gorm:"primaryKey"`
	// TODO : Group name or External Table ????

	Won      int `json:"won"`
	Drawn    int `json:"drawn"`
	Lost     int `json:"lost"`
	Played   int `json:"played"`
	Position int `json:"position"`
	Diff     int `json:"diff"`
	Against  int `json:"against"`
	Scored   int `json:"scored"`
}

//func (Standings) BeforeCreate(db *gorm.DB) error {
// ...
//}

// Change model Person's field Addresses' join table to PersonAddress
// PersonAddress must defined all required foreign keys or it will raise error
//  db.SetupJoinTable(&Person{}, "Addresses", &PersonAddress{})
