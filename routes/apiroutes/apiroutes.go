package apiroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizqullorayhan/go-fiber-gorm/handlers/bookborrowhistoryhandler"
	"github.com/rizqullorayhan/go-fiber-gorm/handlers/bookhandler"
	"github.com/rizqullorayhan/go-fiber-gorm/handlers/categoryhandler"
	"github.com/rizqullorayhan/go-fiber-gorm/handlers/rolehandler"
	"github.com/rizqullorayhan/go-fiber-gorm/handlers/userhandler"
	"github.com/rizqullorayhan/go-fiber-gorm/middleware"
)

func RouteApi(app *fiber.App) {
	api := app.Group("/api")

	api.Route("/book", func(r fiber.Router) {
		r.Get("/", bookhandler.GetAll)
		r.Post("/", bookhandler.Create)
		r.Get("/:id", bookhandler.GetById)
		r.Put("/:id", middleware.RoleMiddleware("Admin"), bookhandler.Update)
		r.Delete("/:id", middleware.RoleMiddleware("Admin"), bookhandler.Delete)
	})

	api.Route("/borrow", func(r fiber.Router) {
		r.Get("/history", bookborrowhistoryhandler.GetAll)
		r.Get("/history/user/:userId", bookborrowhistoryhandler.GetByUser)
		r.Get("/history/book/:bookId", bookborrowhistoryhandler.GetByBook)
		r.Get("/history/user/:userId/book/:bookId", bookborrowhistoryhandler.GetByUserAndBook)
		r.Get("/history/:id", bookborrowhistoryhandler.GetById)
		r.Get("/history/latestBook/:userId", bookborrowhistoryhandler.GetLastBorrowedBook)
		r.Post("/", bookborrowhistoryhandler.Create)
		r.Put("/:id", bookborrowhistoryhandler.Update)
		r.Delete("/:id", bookborrowhistoryhandler.Delete)
	})

	api.Route("/category", func(r fiber.Router) {
		r.Get("/", categoryhandler.GetAll)
		r.Post("/", middleware.RoleMiddleware("Admin"), categoryhandler.Create)
		r.Get("/:id", middleware.RoleMiddleware("Admin"), categoryhandler.GetById)
		r.Put("/:id", middleware.RoleMiddleware("Admin"), categoryhandler.Update)
		r.Delete("/:id", middleware.RoleMiddleware("Admin"), categoryhandler.Delete)
	})

	api.Route("/role", func(r fiber.Router) {
		r.Get("/", middleware.RoleMiddleware("Admin"), rolehandler.GetAll)
		r.Post("/", middleware.RoleMiddleware("Admin"), rolehandler.Create)
		r.Get("/:id", middleware.RoleMiddleware("Admin"), rolehandler.GetById)
		r.Put("/:id", middleware.RoleMiddleware("Admin"), rolehandler.Update)
		r.Delete("/:id", middleware.RoleMiddleware("Admin"), rolehandler.Delete)
	})

	api.Route("/user", func(r fiber.Router) {
		r.Get("/", middleware.RoleMiddleware("Admin"), userhandler.GetAll)
		r.Post("/", middleware.RoleMiddleware("Admin"), userhandler.Create)
		r.Get("/:id", middleware.RoleMiddleware("Admin"), userhandler.GetById)
		r.Put("/:id", middleware.RoleMiddleware("Admin"), userhandler.Update)
		r.Patch("/email/:id", middleware.RoleMiddleware("Admin"), userhandler.UpdateEmail)
		r.Delete("/:id", middleware.RoleMiddleware("Admin"), userhandler.Delete)
	})
}