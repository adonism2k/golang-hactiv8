package routes

import (
	"github.com/adonism2k/golang-hactiv8/api/handlers"
	_ "github.com/adonism2k/golang-hactiv8/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Api(h handlers.Config) *fiber.App {
	app := fiber.New()

	users := app.Group("/users")
	{
		users.Post("/login", h.Login)
		users.Post("/register", h.Register)
		users.Put("/:id", h.EditUser)
		users.Delete("/:id", h.DeleteUser)
	}

	photos := app.Group("/photos")
	{
		photos.Get("/", h.GetPhotos)
		photos.Post("/", h.CreatePhoto)
		photos.Put("/:id", h.EditPhoto)
		photos.Delete("/:id", h.DeletePhoto)
	}

	comments := app.Group("/comments")
	{
		comments.Get("/", h.GetComments)
		comments.Post("/", h.CreateComment)
		comments.Put("/:id", h.EditComment)
		comments.Delete("/:id", h.DeleteComment)
	}

	social := app.Group("/socialmedias")
	{
		social.Get("/", h.GetSocialMedias)
		social.Post("/", h.CreateSocialMedia)
		social.Put("/:id", h.EditSocialMedia)
		social.Delete("/:id", h.DeleteSocialMedia)
	}

	app.Get("/swagger/*", swagger.HandlerDefault)

	return app
}
