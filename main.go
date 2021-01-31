package main

import (
	"graphql/gorest/controller"
	"graphql/gorest/repository"
	"graphql/gorest/service"

	"github.com/gofiber/fiber/v2"
)

const port string = ":8080"

func main() {
	postRepository := repository.NewFireStoreRepository()
	app := fiber.New()

	postService := service.NewPostService(postRepository)
	controller.NewPostController(app, postService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(port)
}
