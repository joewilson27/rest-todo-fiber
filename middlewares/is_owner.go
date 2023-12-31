package middlewares

import (
	"rest-fiber/data"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func IsOwner(c *fiber.Ctx) error {
	userid := c.Locals("userid").(uint)
	todoid, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "data not found",
		})
	}
	for _, todo := range data.Todos {
		// make sure id todo match id that has been sent as request & user match the data
		if todo.ID == uint(todoid) && todo.UserID == userid {
			return c.Next() // pass the middleware because it meets the conditions
		} else if todo.ID == uint(todoid) && todo.UserID != userid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "todo not found", //"you are not the owner",
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "todo not found",
	})
}
