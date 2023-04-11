package exception

import "github.com/gofiber/fiber/v2"

func ErrorBadRequest(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Bad Request",
			"data":    err.Error(),
		})
	}

	return nil
}
