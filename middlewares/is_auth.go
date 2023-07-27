package middlewares

import (
	"rest-fiber/data"
	"rest-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func IsAuth(c *fiber.Ctx) error {
	token := c.Cookies("token", "")
	/*
		tambah cookies pada postman ketika mau request
		pada postman add cookies pada menu cookies dibawah button send.
	*/

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token not provided",
		})
	}
	user := models.User{}
	for _, u := range data.Users {
		if u.Token == token {
			user = u
		}
	}
	if user.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	c.Locals("userid", user.ID)
	/*
		Locals
		A method that stores variables scoped to the request and, therefore, are available only to the routes that
		match the request.
		Jadi jika kita set ke method Locals ini variable, nanti variable nya bisa kita akses di route tujuan
		car aksesnya tinggal panggil c.Locals("propertyname"). PropertyName ini adl property
		yg kita set saat meng-set
	*/

	return c.Next()
}
