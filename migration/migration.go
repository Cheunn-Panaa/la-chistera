package migrations

import (
	"github.com/cheunn-panaa/la-chistera/entity"

	"gorm.io/gorm"
)

func Migrate(DBCon *gorm.DB) {

	DBCon.AutoMigrate(&entity.User{}, &entity.Competition{}, &entity.Group{}, &entity.Team{}, &entity.Standings{}, &entity.Matches{}, &entity.MatchBets{})

	///add this lines below that will add the foreign keys to the migrations.go
	// database.DBCon.Model(&models.User{}).AddForeignKey("address_id", "address(id)", "CASCADE", "RESTRICT")
	// this will add the address_id foreign key to the user table from the address table
	// database.DBCon.Model(&models.Orders{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
	// this will add the user_id foreign key to the order table which is related to user table
	// database.DBCon.Model(&models.OrderItems{}).AddForeignKey("order_id", "orders(id)", "CASCADE", "RESTRICT")
	// Cascade means whenever we delete the parent record the child  record should be deleted
	// that is for example if we delete a user all his orders shall be deleted
	// database.DBCon.Model(&models.OrderItems{}).AddForeignKey("food_id", "foods(id)", "CASCADE", "RESTRICT")

}
