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
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(logger.New(), cors.New(cors.Config{
		AllowOrigins:     "localhost",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
	}))

	app.Post("/users/login", h.Login)
	app.Post("/users/register", h.Register)

	app.Use(middleware.Auth)
	users := app.Group("/users")
	{
		users.Put("/:id", middleware.UserOwner, h.UpdateUser)
		users.Delete("/:id", middleware.UserOwner, h.DeleteUser)
	}

	photos := app.Group("/photos")
	{
		photos.Get("/", h.GetPhotos)
		photos.Post("/", h.CreatePhoto)
		photos.Put("/:id", middleware.PhotoOwner, h.UpdatePhoto)
		photos.Delete("/:id", middleware.PhotoOwner, h.DeletePhoto)
	}

	comments := app.Group("/comments")
	{
		comments.Get("/", h.GetComments)
		comments.Post("/", h.CreateComment)
		comments.Put("/:id", middleware.CommentOwner, h.EditComment)
		comments.Delete("/:id", middleware.CommentOwner, h.DeleteComment)
	}

	social := app.Group("/socialmedias")
	{
		social.Get("/", h.GetSocialMedias)
		social.Post("/", h.CreateSocialMedia)
		social.Put("/:id", middleware.SocialOwner, h.EditSocialMedia)
		social.Delete("/:id", middleware.SocialOwner, h.DeleteSocialMedia)
	}

	return app
}
