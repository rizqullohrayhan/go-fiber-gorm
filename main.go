package main

import (
	"log"
	// _ "net/http/pprof"
    // "net/http"


	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rizqullorayhan/go-fiber-gorm/database"
	"github.com/rizqullorayhan/go-fiber-gorm/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

    // go func() {
    //     http.ListenAndServe(":6060", nil) // Jalankan server pprof pada port 6060
    // }()

	app := fiber.New()

	database.DatabaseInit()

	routes.RouteInit(app)

	app.Listen(":8000")
}