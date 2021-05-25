package routes

import (
	"github.com/cheunn-panaa/la-chistera/usecase/user"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.UseCase) {
	app.Get("/users", getAllUsers(service))
	app.Post("/users", createUser(service))
}

func createUser(service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		/* var requestBody entity.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result, dberr := service.CreateUser(&requestBody)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		}) */
		return nil
	}
}

func getAllUsers(service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
		/*	err := service.ListUsers()
			if err != nil {
				_ = c.JSON(&fiber.Map{
					"status": false,
					"error":  err,
				})
			}
			return c.JSON(&fiber.Map{
				"status": true,
				"users":  "fetched",
			})*/
	}
}
