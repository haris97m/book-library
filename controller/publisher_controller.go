package controller

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/haris97m/go-fiber/exception"
	"github.com/haris97m/go-fiber/helper"
	"github.com/haris97m/go-fiber/model/web"
	"github.com/haris97m/go-fiber/service"
	"gorm.io/gorm"
)

type PublisherControllerImpl struct {
	PublisherService service.PublisherService
}

// CONSTRUCTOR
func NewPublisherControllerImpl(publisherService service.PublisherService) *PublisherControllerImpl {
	return &PublisherControllerImpl{
		PublisherService: publisherService,
	}
}

// ROUTER
func (controller *PublisherControllerImpl) Route(app *fiber.App) {
	api := app.Group("/api")

	publisher := api.Group("/publisher")
	publisher.Get("/", controller.FindAll)
	publisher.Get("/:id", controller.FindById)
	publisher.Post("/", controller.Insert)
	publisher.Put("/:id", controller.Save)
	publisher.Delete("/:id", controller.Delete)
}

// GET ALL
func (controller *PublisherControllerImpl) FindAll(c *fiber.Ctx) error {
	// execute query
	publishers, err := controller.PublisherService.FindAll()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Not Found",
			"data":    err.Error(),
		})
	}

	// give response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK",
		"data":    publishers,
	})
}

// GET BY ID /controller
func (controller *PublisherControllerImpl) FindById(c *fiber.Ctx) error {
	// get id from parameters
	stringId := c.Params("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err.Error(),
		})
	}

	// execute query
	publisher, err := controller.PublisherService.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Not Found",
			"data":    err.Error(),
		})
	}

	// give response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK",
		"data":    publisher,
	})
}

// CREATE
func (controller *PublisherControllerImpl) Insert(c *fiber.Ctx) error {
	request := web.CreatePublisherRequest{}

	// parsing request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err.Error(),
		})
	}

	// validate request
	if err := helper.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err,
		})
	}

	// send request data
	publisher, err := controller.PublisherService.Insert(request)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "Bad Request",
				"data":    err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal Server Error",
			"data":    err.Error(),
		})
	}

	// give response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK",
		"data":    publisher,
	})
}

// UPDATE
func (controller *PublisherControllerImpl) Save(c *fiber.Ctx) error {
	request := web.UpdatePublisherRequest{}

	// get id from parameters
	stringId := c.Params("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err.Error(),
		})
	}
	request.ID = id

	// parsing request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err.Error(),
		})
	}

	// validate request
	if err := helper.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err,
		})
	}

	// send request data
	newPublisher, err := controller.PublisherService.Save(request)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Not Found",
			"data":    err.Error(),
		})
	}

	// give response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK",
		"data":    newPublisher,
	})
}

// DELETE
func (controller *PublisherControllerImpl) Delete(c *fiber.Ctx) error {
	//get id from parameters
	stringId := c.Params("id")
	id, err := strconv.Atoi(stringId)
	exception.ErrorBadRequest(nil, err)

	// execute query
	if err := controller.PublisherService.Delete(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Not Found",
			"data":    err.Error(),
		})
	}

	// give response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK",
		"data":    id,
	})
}
