package controller

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/haris97m/go-fiber/helper"
	"github.com/haris97m/go-fiber/model/web"
	"github.com/haris97m/go-fiber/service"
	"gorm.io/gorm"
)

type AuthorControllerImpl struct {
	Service service.AuthorService
}

// CONSTRUCTOR
func NewAuthorControllerImpl(authorService service.AuthorService) *AuthorControllerImpl {
	return &AuthorControllerImpl{
		Service: authorService,
	}
}

// ROUTER
func (controller *AuthorControllerImpl) Route(app *fiber.App) {
	api := app.Group("/api")

	author := api.Group("/author")
	author.Get("/", controller.FindAll)
	author.Get("/:id", controller.FindById)
	author.Post("/", controller.Insert)
	author.Put("/:id", controller.Save)
	author.Delete("/:id", controller.Delete)
}

func (controller *AuthorControllerImpl) FindAll(c *fiber.Ctx) error {
	authors, err := controller.Service.FindAll()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Not Found",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK",
		"data":    authors,
	})
}

func (controller *AuthorControllerImpl) FindById(c *fiber.Ctx) error {
	stringId := c.Params("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err.Error(),
		})
	}

	author, err := controller.Service.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Not Found",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK",
		"data":    author,
	})
}

func (controller *AuthorControllerImpl) Insert(c *fiber.Ctx) error {
	request := web.AuthorCreateRequest{}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err.Error(),
		})
	}

	if err := helper.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err,
		})
	}

	author, err := controller.Service.Insert(request)
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK",
		"data":    author,
	})
}

func (controller *AuthorControllerImpl) Save(c *fiber.Ctx) error {
	request := web.AuthorUpdateRequest{}

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

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err.Error(),
		})
	}

	if err := helper.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err,
		})
	}

	newAuthor, err := controller.Service.Save(request)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Not Found",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK",
		"data":    newAuthor,
	})

}

func (controller *AuthorControllerImpl) Delete(c *fiber.Ctx) error {
	stringId := c.Params("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err.Error(),
		})
	}

	if err := controller.Service.Delete(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Not Found",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK",
		"data":    id,
	})

}
