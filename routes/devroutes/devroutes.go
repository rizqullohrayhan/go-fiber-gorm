package devroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizqullorayhan/go-fiber-gorm/database"
	"github.com/rizqullorayhan/go-fiber-gorm/database/seeders"
)

func RouteDev(app *fiber.App) {
	dev := app.Group("/dev")
	dev.Get("/seeder", func(c *fiber.Ctx) error {
		if err := seeders.SeederInit(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err,
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "seeder berhasil dijalankan",
		})
	})
	dev.Get("/migrate", func(c *fiber.Ctx) error {
		if err := database.MigrateInit(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err,
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "migration berhasil dijalankan",
		})
	})
}