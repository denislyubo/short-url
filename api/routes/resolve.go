package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/mahadevans87/short-url/database"
)

type JSONResult struct {
	Error string `json:"error"`
}

// Resolve godoc
// @Summary Returns original URL if exists.
// @Description returns not shortened url.
// @Tags root
// @Accept */*
// @Param url path string true "shortened url"
// @Success      301
// @Failure      400 {object} map[string]interface{} "desc"
// @Failure      404 {object} map[string]interface{} "desc"
// @Failure      500 {object} map[string]interface{} "desc"
// @Router /resolve/{url} [get]
func Resolve(ctx *fiber.Ctx) error {
	url := ctx.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short-url not found in db"})
	} else if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal error"})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return ctx.Redirect(value, 301)
}
