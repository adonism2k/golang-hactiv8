package routes

import (
	"github.com/adonism2k/golang-hactiv8/api/handlers"
	"github.com/adonism2k/golang-hactiv8/api/middleware"
	_ "github.com/adonism2k/golang-hactiv8/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func Api(h handlers.Config) *fiber.App {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(logger.New(), cors.New(cors.Config{
		AllowOrigins:     "localhost",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
	}))

	auth := app.Group("/", middleware.JWTProtected())

	app.Post("/users/login", h.Login)
	app.Post("/users/register", h.Register)
	users := auth.Group("/users")
	{
		users.Put("/:id", h.EditUser)
		users.Delete("/:id", h.DeleteUser)
	}

	photos := auth.Group("/photos")
	{
		photos.Get("/", h.GetPhotos)
		photos.Post("/", h.CreatePhoto)
		photos.Put("/:id", h.EditPhoto)
		photos.Delete("/:id", h.DeletePhoto)
	}

	comments := auth.Group("/comments")
	{
		comments.Get("/", h.GetComments)
		comments.Post("/", h.CreateComment)
		comments.Put("/:id", h.EditComment)
		comments.Delete("/:id", h.DeleteComment)
	}

	social := auth.Group("/socialmedias")
	{
		social.Get("/", h.GetSocialMedias)
		social.Post("/", h.CreateSocialMedia)
		social.Put("/:id", h.EditSocialMedia)
		social.Delete("/:id", h.DeleteSocialMedia)
	}
	return app
}
