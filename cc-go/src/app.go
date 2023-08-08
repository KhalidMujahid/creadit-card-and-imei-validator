package src

import (
	"log"
	"os"

	"github.com/khalidMujahid/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func Run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "public")

	// routes
	routes.SetupRouterRoutes(app)

	// page not found
	// app.Use(middlewares.PageNotFound)

	// error handler
	// app.Use(middlewares.ErrorHandler)

	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))

	return nil
}
