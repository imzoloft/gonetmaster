package middleware

import (
	"crypto/sha256"
	"fmt"

	"github.com/imzoloft/gonetmaster/api/account"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func ValidateApiKey(ctx *fiber.Ctx, key string) (bool, error) {
	hashedKey := sha256.Sum256([]byte(key))
	isApiKeyValid, err := account.RetrieveApiKey(fmt.Sprintf("%x", hashedKey))

	if err != nil {
		return false, err
	}

	if !isApiKeyValid {
		return false, keyauth.ErrMissingOrMalformedAPIKey
	}
	return true, nil
}
