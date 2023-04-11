package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/haris97m/go-fiber/config"
	"github.com/haris97m/go-fiber/controller"
	"github.com/haris97m/go-fiber/repository"
	"github.com/haris97m/go-fiber/service"
)

func main() {
	env := config.New()
	db := config.NewDatabase(env)

	publisherRepository := repository.NewPublisherRepositoryImpl(db)
	publisherService := service.NewPublisherServiceImpl(publisherRepository)
	publisherController := controller.NewPublisherControllerImpl(publisherService)

	authorRepository := repository.NewAuthorRepositoryImpl(db)
	authorService := service.NewAuthorServiceImpl(authorRepository)
	authorController := controller.NewAuthorControllerImpl(authorService)

	app := fiber.New()

	publisherController.Route(app)
	authorController.Route(app)

	app.Listen(":3000")
}
