package handlers

import (
	"rest-fiber/data"
	"rest-fiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteTodoHandler(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	for i, todo := range data.Todos {
		if todo.ID == uint(id) {
			// [1, 2, 3, 4] --> data awal, anggap aja id
			// [1, 3, 4 ] --> data 2 di take out
			//  0  1  2  3 --> urutan index pada data
			/*
					misal pada contoh disini kita mau hapus data id 2 (kalo contoh ini, berarti si id 2
					berada pada index 1) lalu index 1 akan kosong, maka kita akan membuat data di belakangnya
				 	maju 1 index)
			*/
			// parameter kedua di copy ke parameter pertama
			copy(data.Todos[i:], data.Todos[i+1:]) // [i:] artinya data dari i sampai selesai (seluruh data)
			// read: https://stackoverflow.com/questions/47722542/what-does-the-symbol-mean-in-go

			data.Todos[len(data.Todos)-1] = models.Todo{} // ata.Todos[len(data.Todos)-1] --> cari data paling terakhir, kemudian value di isi data kosong models.Todo{}
			data.Todos = data.Todos[:len(data.Todos)-1]   // [:len(data.Todos)-1] --> dari awal sampai akhir data | perbaharui data todos dengan data baru dan indexing
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": "data deleted",
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "todo not found",
	})
}
