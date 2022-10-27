package routes

import (
	"github.com/adonism2k/golang-hactiv8/api/handlers"
	"github.com/adonism2k/golang-hactiv8/api/middleware"
	_ "github.com/adonism2k/golang-hactiv8/docs"
	"github.com/adonism2k/golang-hactiv8/internal/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func Api(h handlers.Config, Env initializers.Config) *fiber.App {
	app := fiber.New()
	app.Post("/users/login", h.Login)
	app.Post("/users/register", h.Register)
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(logger.New(), cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
	}))

	users := app.Group("/users", middleware.Auth)
	{
		users.Put("/:id", middleware.UserOwner, h.UpdateUser)
		users.Delete("/:id", middleware.UserOwner, h.DeleteUser)
	}

	photos := app.Group("/photos", middleware.Auth)
	{
		photos.Get("/", h.GetPhotos)
		photos.Post("/", h.CreatePhoto)
		photos.Put("/:id", h.UpdatePhoto)
		photos.Delete("/:id", h.DeletePhoto)
	}

	comments := app.Group("/comments", middleware.Auth)
	{
		comments.Get("/", h.GetComments)
		comments.Post("/", h.CreateComment)
		comments.Put("/:id", h.UpdateComment)
		comments.Delete("/:id", h.DeleteComment)
	}

	social := app.Group("/socialmedias", middleware.Auth)
	{
		social.Get("/", h.GetSocialMedias)
		social.Post("/", h.CreateSocialMedia)
		social.Put("/:id", h.UpdateSocialMedia)
		social.Delete("/:id", h.DeleteSocialMedia)
	}

	return app
}
