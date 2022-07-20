package routes

import (
	"fmt"
	databse "url-shortener/api/database"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")
	r := databse.CreateClient(0)
	defer r.Close()

	oriUrl, err := r.Get(databse.Ctx, url).Result()
	fmt.Println(oriUrl)
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short not found in the database"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connet to DB"})
	}
	rInr := databse.CreateClient(1)
	rInr.Incr(databse.Ctx, "counter")
	defer rInr.Close()

	return c.Redirect(oriUrl, 301)
}
