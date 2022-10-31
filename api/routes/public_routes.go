package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahadevans87/short-url/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/resolve/:url", controllers.Resolve)
	route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens
}
