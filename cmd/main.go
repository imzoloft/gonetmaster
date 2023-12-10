package main

import (
	"log"

	"github.com/imzoloft/gonetmaster/api/common/model"
	"github.com/imzoloft/gonetmaster/api/database"
	"github.com/imzoloft/gonetmaster/api/router"
	aggregatedLogger "github.com/imzoloft/gonetmaster/logger"
	"github.com/imzoloft/gonetmaster/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	file := util.RetrieveOrCreateFile("./api/logs/logs.txt")

	aggregatedLogger.Log = aggregatedLogger.New(file)

	err := godotenv.Load()

	if err != nil {
		aggregatedLogger.Log.Error("Error loading .env file")
	}

	database.InitializeDatabase()
	defer database.Db.Close()

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${method} | ${path} | ${ip} | ${status} | ${latency}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		Output:     file,
	}))

	app.Use(limiter.New(limiter.Config{
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(model.Response{
				Status:  "error",
				Message: "Too many requests",
				Data:    nil,
			})
		},
	}))

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
