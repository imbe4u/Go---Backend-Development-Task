package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestLogger(log *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		reqID := uuid.New().String()

		c.Set("X-Request-ID", reqID)

		err := c.Next()

		log.Info("request",
			zap.String("request_id", reqID),
			zap.String("path", c.Path()),
			zap.Duration("duration", time.Since(start)),
		)

		return err
	}
}
