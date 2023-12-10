package router

import (
	"github.com/imzoloft/gonetmaster/api/middleware"
	"github.com/imzoloft/gonetmaster/api/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", keyauth.New(keyauth.Config{
		Validator: middleware.ValidateApiKey,
		KeyLookup: "header:X-API-KEY",
	}), user.GetUsers)
	app.Post("/users", user.CreateUser)
}
