package route

import (
	"github.com/firdausalif/challenge-todolist/platform/database"
	"github.com/gofiber/fiber/v2"
	"os"
)

func GeneralRoute(a *fiber.App) {
	a.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":       fiber.StatusOK,
			"message":      "Welcome to Fiber Go API!",
			"docs":         "/swagger/index.html",
			"health_check": "/h34l7h",
			"environment":  os.Getenv("APP_ENV"),
		})
	})

	a.Get("/h34l7h", func(c *fiber.Ctx) error {
		db, err := database.GetDB().DB.DB()
		err = db.Ping()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"status":    fiber.StatusOK,
			"message":   "Health Check " + database.GetDB().Migrator().CurrentDatabase(),
			"db_online": true,
		})
	})
}

func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  fiber.StatusNotFound,
				"message": "sorry, endpoint is not found",
			})
		},
	)
}
