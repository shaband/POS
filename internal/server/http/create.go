package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

//	@title			POS API
//	@version		1.0
//	@description	This is a POS api docs
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	Shabandinho
//	@contact.email	mahmoudshaband@gmail.com
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:3500
//	@BasePath		/
func Create() *fiber.App {
	app := fiber.New(fiber.Config{})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: time.Second * 60,
	}))

	return app
}
