package middleware

import (
	"strings"

	"github.com/Gylmynnn/chat-cuy-be/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.SendError(c, fiber.StatusUnauthorized, "Authorization header is required")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return utils.SendError(c, fiber.StatusUnauthorized, "Invalid authorization format")
		}

		claims, err := utils.ValidateJWT(jwtSecret, parts[1])
		if err != nil {
			return utils.SendError(c, fiber.StatusUnauthorized, "Invalid or expired token")
		}

		c.Locals("userID", claims.UserID)
		c.Locals("username", claims.Username)
		c.Locals("email", claims.Email)

		return c.Next()
	}
}
