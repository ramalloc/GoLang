package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ramalloc/go-url-shortner/database"
	"github.com/ramalloc/go-url-shortner/helpers"
)

// Defining Request struct
type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"custom_short"`
	Expiry      time.Duration `json:"expiry"`
}

// Defining Response struct
type response struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"custom_short"`
	Expiry      time.Duration `json:"expiry"`

	// Below are for declaring limits for front-end request and Reseting the limits
	XRateRemaning   int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	// STEP 1
	body := new(request)

	// parser which is provided by fibre to understand json by go. To keep josn into struct
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	// If there is not error then it means request data perse in body

	// STEP 2
	// Implement Rate Limiting
	// everytime a user queries, check if the IP is already in database,
	// if yes, decrement the calls remaining by one, else add the IP to database with expiry of `30mins`.
	//  So in this case the user will be able to send 10
	// requests every 30 minutes

	// -> We will check that the ip address already in the database
	// If not - then we will create a quota for 10 rate limit api requestes
	// If yes - then we will decreament the rate limit / api request

	// Creating R2 as Second Database to store user's IP Address and there API_QUOTA
	r2 := database.CreateClient(1)
	defer r2.Close()

	// Passing IP as key
	value, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		// Getting API Quoota as value
		valInt, _ := strconv.Atoi(value)

		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":            "Rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
		}
	}

	// STEP 3
	// Checking that the input URL is valid/actual or not
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// STEP 4
	// Check for Domain Error
	// We are using the RemoveDomainError function to prevent users from shortening or interacting with links that belong
	// to the same domain as the application itself. This ensures that the application does not create self-referencing
	// links, which could lead to infinite redirections, abuse, or unintended behavior.
	// attack example if main domain is shortable -> short.com/malicious → short.com/login?redirect=http://attacker.com
	// Throwing error if url and domain is same
	if !helpers.IsUrlAndDomainSame(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "You can't access this domain URL"})
	}

	// Enforce HTTPS, SSL
	// Adding http:// in starting of url if not present
	body.URL = helpers.EnforceHTTP(body.URL)

	// The function decrements (reduces) the value by 1 stored in Redis for a given key. Decrement the Rate Remaining.
	r2.Decr(database.Ctx, c.IP())

	// Here we are providing funcnality to user to mkae their own customized short url
	// In order to do that we to accept the custom url from user then we will search in database
	// if already exists then send message to user already in use
	// If not then create random shorthen url using uuid

	// Creating Unique ID
	var id string
	if body.CustomShort == "" {
		// Create new id using uuid
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	// Creating Database 0 to store short url and Actual URL
	r := database.CreateClient(0)
	defer r.Close()

	// Check already in use or not, Getting value (Actual URL) with respect to uer's provided custom_short url from DB
	val, _ := r.Get(database.Ctx, id).Result()
	// If got the value (actual url) with respect to provided key (custom_url) in DB then we will not save the user's
	// custom_short in database and will send an error message
	if val != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "provided custom url already in use",
		})
	}

	if body.Expiry == 0 {
		body.Expiry = 30
	}

	// Setting Whole data in database
	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to connect to server",
		})
	}

	// Response that will be send to user
	resp := response{
		URL:             body.URL,
		CustomShort:     "",
		Expiry:          body.Expiry,
		XRateRemaning:   100,
		XRateLimitReset: 30,
	}

	// Getting Rate remaining with respect to an IP from Database
	rem, _ := r2.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaning, _ = strconv.Atoi(rem)

	// Calculating Rate limit Reset
	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	// Creating Short URL
	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id

	return c.Status(fiber.StatusOK).JSON(resp)

}
