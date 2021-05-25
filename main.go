package main

import (
	"github.com/cheunn-panaa/la-chistera/repository"
	"github.com/cheunn-panaa/la-chistera/routes"
	"github.com/cheunn-panaa/la-chistera/usecase/user"
	"github.com/gofiber/fiber/v2"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Starting the application...")

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture mongo book shop!"))
	})

	api := app.Group("/api")
	dsn := "root:root@tcp(127.0.0.1:3306)/chistera?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Initialize each repository
	log.Info().Msg("Initialising Repositories...")
	userRepo := repository.NewUserRepo(db)

	// Initialize each pkg
	log.Info().Msg("Initialising Services...")
	userService := user.NewService(userRepo)

	// Initialize Routers when services are done
	routes.UserRouter(api, userService)

	_ = app.Listen(":8080")
}
