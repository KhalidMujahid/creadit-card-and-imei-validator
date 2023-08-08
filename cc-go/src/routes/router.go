package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khalidMujahid/src/controllers"
)

func SetupRouterRoutes(app *fiber.App) {
	app.Get("/", controllers.RenderWelcomePage)
	app.Get("/verify", controllers.RenderCCPage)
	app.Post("/verify", controllers.Handler)
}
