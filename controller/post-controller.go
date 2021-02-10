package controller

import (
	"graphql/gorest/entity"
	"graphql/gorest/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Controller for
type Controller struct {
	service service.PostService
}

// Response for
type Response struct {
	Message string `json:"message"`
}

// NewPostController for
func NewPostController(router *fiber.App, service service.PostService) {
	handler := &Controller{
		service: service,
	}

	router.Get("/posts", handler.GetPosts)
	router.Post("/posts", handler.AddPost)
}

// GetPosts for
func (cont *Controller) GetPosts(c *fiber.Ctx) error {
	res := Response{}
	posts, err := cont.service.FindAll()
	if err != nil {
		return err
	}
	res.Message = "success"
	return c.JSON(posts)
}

// AddPost for
func (cont *Controller) AddPost(c *fiber.Ctx) error {
	req := entity.Post{}
	err := c.BodyParser(&req)
	if err != nil {
		log.Println(err)
		return err
	}
	err = cont.service.Validate(&req)
	if err != nil {
		return err
	}

	result, err := cont.service.Create(&req)
	if err != nil {
		return err
	}
	return c.JSON(result)
}
