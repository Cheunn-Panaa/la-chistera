package main

import (
	"fmt"

	"github.com/cheunn-panaa/la-chistera/config"
	migrations "github.com/cheunn-panaa/la-chistera/migration"
	"github.com/cheunn-panaa/la-chistera/repository"
	"github.com/cheunn-panaa/la-chistera/routes"
	"github.com/cheunn-panaa/la-chistera/usecase/user"
	"github.com/gofiber/fiber/v2"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	configFile := "config.yaml" //os.Getenv("CONFIG_FILE")
	config, err := config.NewConfig(configFile)

	if err != nil {
		panic("Fail to retrieve configuration")
	}
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	log.Info("Starting the application...")

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture mongo book shop!"))
	})

	api := app.Group("/api")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Db.User,
		config.Db.Password,
		config.Db.Host,
		config.Db.Port,
		config.Db.Name,
	)
	dbCon, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrations.Migrate(dbCon)

	// Initialize each repository
	log.Info("Initialising Repositories...")
	userRepo := repository.NewUserRepo(dbCon)

	// Initialize each pkg
	log.Info("Initialising Services...")
	userService := user.NewService(userRepo)

	// Initialize Routers when services are done
	routes.UserRouter(api, userService)

	_ = app.Listen(":8080")
}
