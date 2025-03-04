package routes

import (
	"github.com/ramalloc/go-url-shortner/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// Our Reosolve function job is when someone use the short url then it should redirect to the actual long url
// THis will work like this that when user entered url it will check in db and get the long url with respect to the short url.
func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	// Using Database 1 which has id (short url) as key and Actual/Original URL as Value.
	r := database.CreateClient(0)
	defer r.Close()

	// Check in database that url exist or not
	fetchedUrl, err := r.Get(database.Ctx, url).Result()
	// The key (short url / id) does not exist → err == redis.Nil
	if err == redis.Nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error" : "Given Short URL doesn't exist or invalid"})
	} else if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : "cannot connect to database"})
	}

	// Using Database 2 which has user's IP as key and API Quota as Value.
	// Incrementing the counter/api_quota
	rIncrement := database.CreateClient(1)
	defer rIncrement.Close()

	rIncrement.Incr(database.Ctx, "counter")
	// redirect url, status code 301
	return c.Redirect(fetchedUrl, 301)

}
