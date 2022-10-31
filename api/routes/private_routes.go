package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahadevans87/short-url/controllers"
	"github.com/mahadevans87/short-url/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/shorten", middleware.JWTProtected(), controllers.Shorten)
}
