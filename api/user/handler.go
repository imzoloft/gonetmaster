package user

import (
	"crypto/rand"

	"github.com/imzoloft/gonetmaster/api/common/model"
	"github.com/imzoloft/gonetmaster/api/common/validation"
	"github.com/imzoloft/gonetmaster/api/config"
	"github.com/imzoloft/gonetmaster/logger"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error {
	users, err := RetrieveUsers()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status:  "error",
			Data:    nil,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Status:  "success",
		Data:    users,
		Message: nil,
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(User)
	parseErr := ctx.BodyParser(user)

	if parseErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status:  "fail",
			Data:    parseErr.Error(),
			Message: nil,
		})
	}
	user.Key, parseErr = generateRandomAESKey()

	validateErr := validation.ValidateStruct(user)

	if validateErr != nil {
		// We do not want to give informations about our internal structure.
		if config.IsDebug() {
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
				Status:  "fail",
				Data:    validateErr,
				Message: nil,
			})
		} else {
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
				Status:  "fail",
				Data:    "error",
				Message: nil,
			})
		}
	}

	insertErr := InsertUser(user.Id, user.Key)

	if insertErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status:  "error",
			Data:    nil,
			Message: insertErr.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Status:  "success",
		Data:    fiber.Map{"key": user.Key},
		Message: "User added successfully",
	})
}

func generateRandomAESKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)

	if err != nil {
		logger.Log.Warn("Error generating random AES key: ", err)
		return nil, err
	}
	return key, nil
}
