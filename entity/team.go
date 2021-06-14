package entity

import (
	"gorm.io/gorm"
)

//Construct your model under entity.
type Team struct {
	gorm.Model

	Competitions []Competition `gorm:"many2many:standings"`
	Matches      []Team        `gorm:"foreignkey:id;associationforeignkey:away,home;many2many:matches"`

	Country string `json:"country"`
	Iso     string `json:"iso"`
	//TODO remove ? where ?
	Quotation int `json:"quotation"`
}

// Only one struct per file should exists unless another struct is closely related with the one defined in this file.
type DeleteRequest struct {
	ID string `json:"id" binding:"required,gte=1"`
}
