package routes

import (
	"github.com/adonism2k/golang-hactiv8/api/handlers"
	"github.com/adonism2k/golang-hactiv8/api/middleware"
	_ "github.com/adonism2k/golang-hactiv8/docs"
	"github.com/adonism2k/golang-hactiv8/internal/initializers"
	"github.com/adonism2k/golang-hactiv8/internal/model"
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
	app.Get("/health", func(c *fiber.Ctx) error {
		user := c.Locals("user").(model.User)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "OK",
			"claims":  user,
		})
	})

	users := app.Group("/users")
	{
		users.Put("/:id", h.UpdateUser)
		users.Delete("/:id", h.DeleteUser)
	}

	photos := app.Group("/photos")
	{
		photos.Get("/", h.GetPhotos)
		photos.Post("/", h.CreatePhoto)
		photos.Put("/:id", middleware.IsPhotoOwner, h.EditPhoto)
		photos.Delete("/:id", middleware.IsPhotoOwner, h.DeletePhoto)
	}

	comments := app.Group("/comments")
	{
		comments.Get("/", h.GetComments)
		comments.Post("/", h.CreateComment)
		comments.Put("/:id", middleware.IsCommentOwner, h.EditComment)
		comments.Delete("/:id", middleware.IsCommentOwner, h.DeleteComment)
	}

	social := app.Group("/socialmedias")
	{
		social.Get("/", h.GetSocialMedias)
		social.Post("/", h.CreateSocialMedia)
		social.Put("/:id", middleware.IsSocialOwner, h.EditSocialMedia)
		social.Delete("/:id", middleware.IsSocialOwner, h.DeleteSocialMedia)
	}

	return app
}
