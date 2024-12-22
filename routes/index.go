package routes

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/rizqullorayhan/go-fiber-gorm/config"
	"github.com/rizqullorayhan/go-fiber-gorm/handlers/authhandler"
	"github.com/rizqullorayhan/go-fiber-gorm/routes/apiroutes"
	"github.com/rizqullorayhan/go-fiber-gorm/routes/devroutes"
	"github.com/rizqullorayhan/go-fiber-gorm/utils"
)

func RouteInit(app *fiber.App) {
	app.Static("/public", config.ProjectRootPath+"/public/asset")

	app.Post("/login", authhandler.Login)
	app.Post("/register", authhandler.Register)
	devroutes.RouteDev(app)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: utils.SecretKey},
	}))
	apiroutes.RouteApi(app)
}